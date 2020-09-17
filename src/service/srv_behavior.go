package service

import (
	"fmt"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

func GetBehaviorList(f *utils.BehaviorQueryForm) (ret []*models.BehaviorRecord, err error) {
	var q *query.Query
	if q, err = query.NewQuery(f.Page, f.PageSize).ValidCond(f.ToMap()); err != nil {
		return
	}
	return models.GetBehaviorRecordList(q)
}

func RemoveBehavior(uid string, ids []int64) (err error) {
	var (
		r      = new(models.UserBase)
		m      = map[string]interface{}{}
		data   []*models.BehaviorRecord
		delIds []int64
	)
	if err = r.GetByUId(uid); err != nil {
		return
	}
	if data, err = models.GetBehaviorListByIDs(ids); err != nil {
		return
	}
	for _, v := range data {
		if v.Uid == uid {
			err = utils.ErrActionNotAllow.SetMeta("can't delete self")
			return
		}
		if _, ok := m[v.Uid]; ok {
			continue
		}
		u := new(models.UserBase)
		if err = u.GetByUId(v.Uid); err != nil {
			return
		}
		if u.Role.GetLevelMust() >= r.Role.GetLevelMust() {
			err = utils.ErrActionNotAllow.SetMeta(fmt.Sprintf("role level less than %s > %s",
				r.Role.Role, u.Role.Role))
			return
		}
		m[v.Uid] = true
		delIds = append(delIds, int64(v.ID))
	}
	if len(delIds) == 0 {
		return
	}
	return models.DeleteBehaviorList(delIds)
}
