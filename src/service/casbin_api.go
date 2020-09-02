package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/setting"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

// AddRuleAPI 增加角色api权限
func AddRuleAPI(uid string,f *utils.FormRoleAPIPerm) (ret []int64, err error) {
	var (
		e    *casbin.Enforcer
		role = new(models.UserBase)
		user = new(models.UserBase)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(f.Uid); err != nil {
		return
	}
	if err = role.GetByUId(uid); err != nil {
		return
	}
	if user.Role >= role.Role {
		err = utils.ErrActionNotAllow
		return
	}
	for _, v := range f.GroupID {
		dataGroup := new(models.APIGroup)
		if err = dataGroup.Get(v); err != nil {
			return
		}
		if _, err = e.AddPolicy(f.Uid, dataGroup.Path, dataGroup.Method); err != nil {
			ret = append(ret, v)
			continue
		}

	}

	return
}

// RemoveRuleAPI 删除
func RemoveRuleAPI(uid string,f *utils.FormRoleAPIPerm) (ret []int64, err error) {
	var (
		e    *casbin.Enforcer
		user = new(models.UserBase)
		role =  new(models.UserBase)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(f.Uid); err != nil {
		return
	}
	if err = role.GetByUId(uid); err != nil {
		return
	}
	if user.Role >= role.Role {
		err = utils.ErrActionNotAllow
		return
	}
	for _, v := range f.GroupID {
		dataGroup := new(models.APIGroup)
		if err = dataGroup.Get(v); err != nil {
			return
		}
		if _, err = e.RemovePolicy(f.Uid, dataGroup.Path, dataGroup.Method); err != nil {
			ret = append(ret, v)
			continue
		}

	}
	return
}

// GetRuleApiList
func GetRuleApiList(uid string) (ret []struct {
	Path   string
	Method string
}, err error) {
	var d []*models.RuleAPI
	if d, err = models.GetRuleAPIListByUID(uid); err != nil {
		return nil, err
	}
	for _, v := range d {
		ret = append(ret, struct {
			Path   string
			Method string
		}{Path: v.Path, Method: v.Method})
	}
	return
}
