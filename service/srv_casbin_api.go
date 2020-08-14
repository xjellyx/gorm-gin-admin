package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/setting"
	"github.com/olongfen/user_base/utils"
)

// AddRuleAPI 增加角色api权限
func AddRuleAPI(f *utils.FormRoleAPIPerm) (ret []int64, err error) {
	var (
		e    *casbin.Enforcer
		user = new(models.UserBase)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(f.Uid); err != nil {
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
func RemoveRuleAPI(f *utils.FormRoleAPIPerm) (ret []int64, err error) {
	var (
		e    *casbin.Enforcer
		user = new(models.UserBase)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(f.Uid); err != nil {
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
	d := []*models.RuleAPI{}
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
