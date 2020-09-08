package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/setting"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

// AddRoleAPI 增加角色api权限
func AddRoleAPI(uid string,f *utils.FormRoleAPIPerm) (ret []struct{
	Method string `json:"method"`
	Path string `json:"path"`
}, err error) {
	var (
		e    *casbin.Enforcer
		role = new(models.Role)
		dataCasbin = new(models.RoleAPI)
		user = new(models.UserBase)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(uid); err != nil {
		return
	}

	if err = role.GetByRole(f.Role); err != nil {
		return
	}
	// 不能修改等级比自己大的
	if user.Role.GetLevelMust() <= role.GetLevelMust() && user.Role.GetLevelMust()<setting.Setting.MaxRoleLevel {
		err = utils.ErrActionNotAllow
		return
	}
	for _, v := range f.Groups {
		dataGroup := new(models.APIGroup)
		if err = dataGroup.GetBPathAndMethod(v.Path,v.Method); err != nil {
			logServe.Infoln(err,"path: ",dataGroup.Path,"method:",dataGroup.Method)
			continue
		}
		if err  = dataCasbin.GetByPathAndMethodAndRole(dataCasbin.Path,dataGroup.Method,role.Role);err==nil{
			logServe.Infoln( utils.ErrRoleAPIExist,"path: ",dataGroup.Path,"method:",dataGroup.Method)
			continue
		}
		if _, err = e.AddPolicy(f.Role, dataGroup.Path, dataGroup.Method); err != nil {
			ret = append(ret, v)
			continue
		}

	}

	return
}

// AddRoleGroup
func AddRoleGroup(uid string,f *utils.FormRoleAPIPerm)  {

}

// RemoveRoleAPI 删除
func RemoveRoleAPI(uid string,f *utils.FormRoleAPIPerm) ( err error) {
	var (
		e    *casbin.Enforcer
		user = new(models.UserBase)
		role =  new(models.Role)
	)

	if e, err = casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter); err != nil {
		return
	}
	if err = user.GetByUId(uid); err != nil {
		return
	}
	if err = role.GetByRole(f.Role); err != nil {
		return
	}
	if user.Role.GetLevelMust() <= role.GetLevelMust() && user.Role.GetLevelMust()<setting.Setting.MaxRoleLevel  {
		err = utils.ErrActionNotAllow
		return
	}
	for _, v := range f.Groups {
		dataGroup := new(models.APIGroup)
		if err = dataGroup.GetBPathAndMethod(v.Path,v.Method); err != nil {
			continue
		}
		if _, err = e.RemovePolicy(f.Role, dataGroup.Path, dataGroup.Method); err != nil {
			continue
		}

	}
	return
}

type RoleApiResp struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// GetRoleApiList
func GetRoleApiList(role string) (ret []RoleApiResp, err error) {
	var d []*models.RoleAPI
	if d, err = models.GetRoleAPIListByRole(role); err != nil {
		return nil, err
	}
	for _, v := range d {
		ret = append(ret, RoleApiResp{Path: v.Path, Method: v.Method})
	}
	return
}
