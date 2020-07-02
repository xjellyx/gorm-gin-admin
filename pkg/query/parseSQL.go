package query

import (
	"fmt"
	"sort"
	"strings"
)

// Tag key
const (
	// TagSep mean value
	TagSep = '$'
	// TagValNe 不等于
	TagValNe = `$ne$`
	// TagValNo 不等于 alias
	TagValNo = `$no$`
	// TagValLt 小于
	TagValLt = `$lt$`
	// TagValLte 小于等于
	TagValLte = `$lte$`
	// TagValGt 大于
	TagValGt = `$gt$`
	// TagValGte 大于等于
	TagValGte = `$gte$`
	// TagValLike like
	TagValLike = `$like$`
	// TagValText 全文搜索
	TagValText = `$text$`
	// TagQueryKeyOr mean key or
	TagQueryKeyOr = `$or$`
	// TagQueryKeyAnd and
	TagQueryKeyAnd = `$and$`
	// TagQueryKeyIn in
	TagQueryKeyIn = `$in$`
	// TagUpdateInc mean update
	TagUpdateInc = `$inc$` // 数据库级别增
)

const (
	// SQLValEq sql 等于
	SQLValEq = `=`
	// SQLValNe  不等于
	SQLValNe = `!=`
	// SQLValLike like
	SQLValLike = `LIKE`
	// SQLValLt 小于
	SQLValLt = `<`
	// SQLValLte 小于等于
	SQLValLte = `<=`
	// SQLValGt 大于
	SQLValGt = `>`
	// SQLValGte 大于等于
	SQLValGte = `>=`
	// SQLValAnd 与
	SQLValAnd = `AND`
	// SQLValOr 或
	SQLValOr = `OR`
)

var (
	// ParseMapMax 目前允许5层解析
	ParseMapMax = 5
)

// ParseSQL 将songo格式解析为sql
func ParseSQL(m map[string]interface{}, prefix int) (string, []interface{}, error) {
	return parserSQL(m, prefix, "?")
}

const (
	sqlSepAnd = " " + SQLValAnd + " "
	sqlSepOr  = " " + SQLValOr + " "
)

// parserSQLUnit 转为sql条件语句 TODO: 全文搜索
func parserSQLUnit(k string, v interface{}, idx int, sep string) (sql string, val interface{}, err error) {
	// valid
	if len(k) != 0 {
		k = strings.ToLower(k)
	} else {
		err = ErrMapKeyInvalid
		return
	}
	// 防止sql注入, key不能含有空格,括号
	for i := 0; i < len(k); i++ {
		switch k[i] {
		case ' ', '(', ')':
			err = ErrMapKeyInvalid
			return
		default:
			continue
		}
	}

	var tag string // 比较符
	val = v        // 值

	// val:尝试从key中解析比较符
	if k[len(k)-1] == TagSep {
		hIdx := -1
		for i := len(k) - 2; i > -1; i-- {
			if k[i] == TagSep {
				hIdx = i
				break
			}
		}
		if hIdx > -1 {
			if hIdx == 0 {
				err = ErrMapKeyInvalid
				return
			}
			tag = k[hIdx:]
			k = k[0:hIdx]
			if len(k) == 0 {
				err = ErrMapKeyInvalid
				return
			}
		}
	}

	// val:尝试从val中解析比较符
	if len(tag) == 0 {
		switch _val := v.(type) {
		case string:
			if _t, _v := isTagVal(_val); len(_t) > 0 {
				tag = _t
				val = _v
			}
		case map[string]interface{}:
			if len(_val) != 1 {
				// map表示的val只有一对key-val
				err = ErrMapValMapMultiple
				return
			}
			for _t, _v := range _val {
				tag = _t
				val = _v
			}
		default:
			break
		}
	}

	// sql
	if len(tag) > 0 {
		switch tag {
		case TagValLike, TagValText: // TODO: 全文搜索
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValLike)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValLike, idx)
			}
		case TagValLt:
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValLt)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValLt, idx)
			}
		case TagValLte:
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValLte)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValLte, idx)
			}
		case TagValGt:
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValGt)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValGt, idx)
			}
		case TagValGte:
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValGte)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValGte, idx)
			}
		case TagValNo, TagValNe:
			if sep == "?" {
				sql = fmt.Sprintf("%s %s ?", k, SQLValNe)
			} else {
				sql = fmt.Sprintf(`%s %s $%d`, k, SQLValNe, idx)
			}
		default:
			err = ErrMapOperatorInvalid
			return
		}
	} else {
		if sep == "?" {
			sql = fmt.Sprintf("%s %s ?", k, SQLValEq)
		} else {
			sql = fmt.Sprintf(`%s %s $%d`, k, SQLValEq, idx)
		}
	}

	return
}

// parserSQL 解析M为sql
func parserSQL(m map[string]interface{}, prefix int, sep string) (sql string, vals []interface{}, err error) {
	var (
		nameLis []string
	)
	vals = []interface{}{}

	// 相同的map,for的顺序可能不同,会导致bug
	var (
		idx = 0
		lis = make([]string, len(m))
	)
	for k := range m {
		lis[idx] = k
		idx++
	}
	sort.Strings(lis)
	for _, k := range lis {
		// and or 查询, 支持二级嵌套
		v := m[k]
		// 递归解析
		if err = parserSQLOperator(0, &prefix, sep, k, v, &nameLis, &vals); err != nil {
			return
		}
	}
	sql = strings.Join(nameLis, fmt.Sprintf(" %s ", SQLValAnd))
	return
}

// parserSQLOperator 解析操作符
func parserSQLOperator(deep int, prefix *int, sep string, k string, v interface{}, nameLis *[]string, vals *[]interface{}) (err error) {
	if deep+1 >= ParseMapMax {
		err = ErrMapDeepOutOf
		return
	}
	oper := ""
	oper, k = isKeyOperator(k)
	switch oper {
	case TagQueryKeyOr, TagQueryKeyAnd:
		// 解析语句
		var _nameLis []string
		_lis, ok := v.([]interface{})
		if !ok {
			_lis = []interface{}{v} // 默认都是数组
		}
		for _, v2 := range _lis {
			switch _v := v2.(type) {
			case map[string]interface{}:
				var (
					idx = 0
					lis = make([]string, len(_v))
				)
				for k := range _v {
					lis[idx] = k
					idx++
				}
				sort.Strings(lis)

				for _, k := range lis {
					v := _v[k]
					org := k
					oper := ""
					oper, k = isKeyOperator(k)
					switch oper {
					case TagQueryKeyOr, TagQueryKeyAnd:
						// 解析语句
						_nameLis2 := []string{}
						_lis2, ok := v.([]interface{})
						if !ok {
							_lis2 = []interface{}{v} // 默认都是数组
						}

						//
						for _, v3 := range _lis2 {
							if err = parserSQLOperator(deep+1, prefix, sep, org, v3, &_nameLis2, vals); err != nil {
								return
							}
						}

						// OR 查询
						if oper == TagQueryKeyOr && len(_nameLis2) > 0 {
							_nameLis = append(_nameLis, strings.Join(_nameLis2, sqlSepOr))
						} else if len(_nameLis2) > 0 {
							_nameLis = append(_nameLis, strings.Join(_nameLis2, sqlSepAnd))
						}
					default:
						*prefix++
						if _sql, _val, err2 := parserSQLUnit(k, v, *prefix, sep); err2 == nil {
							_nameLis = append(_nameLis, _sql)
							*vals = append(*vals, _val)
						} else {
							err = err2
							return
						}
					}
				}
			default:
				// string int 等
				*prefix++
				if _sql, _val, err2 := parserSQLUnit(k, _v, *prefix, sep); err2 == nil {
					_nameLis = append(_nameLis, _sql)
					*vals = append(*vals, _val)
				} else {
					err = err2
					return
				}
			}
		}

		// OR 查询
		if oper == TagQueryKeyOr && len(_nameLis) > 0 {
			*nameLis = append(*nameLis, strings.Join(_nameLis, sqlSepOr))
		} else if len(_nameLis) > 0 {
			// AND 查询
			*nameLis = append(*nameLis, strings.Join(_nameLis, sqlSepAnd))
		}
	default:
		*prefix++
		if _sql, _val, err2 := parserSQLUnit(k, v, *prefix, sep); err2 == nil {
			*nameLis = append(*nameLis, "("+_sql+")")
			*vals = append(*vals, _val)
		} else {
			err = err2
			return
		}
	}

	return
}

// 判断一个字符串是否含tagVal, 有则返回
func isTagVal(s string) (tag string, val string) {
	if val = s; len(val) <= 2 || val[0] != TagSep {
		return
	}
	if idx := strings.Index(s[1:], string(TagSep)); idx > -1 {
		tag = s[0 : idx+2]
		val = s[idx+2:]
	}
	return
}

// 判断一个key中是否有操作符
func isKeyOperator(s string) (oper string, key string) {
	if key = s; len(key) <= 2 || key[0] != TagSep {
		return
	}
	if idx := strings.Index(s[1:], string(TagSep)); idx > -1 {
		oper = s[0 : idx+2]
		key = s[idx+2:]
	}
	return
}
