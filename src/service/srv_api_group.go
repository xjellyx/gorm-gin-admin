package service

import (
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"strings"
)

// AddAPIGroup
func AddAPIGroup(f []*utils.FormAPIGroupAdd) (ret []*models.APIGroup, err error) {
	var (
		datas []*models.APIGroup
	)

	for _, v := range f {
		var (
			data = new(models.APIGroup)
		)
		data.ApiGroup = v.ApiGroup
		data.Description = v.Description
		data.Path = v.Path
		data.Method = strings.ToLower(v.Method)

		datas = append(datas, data)
	}
	if err = models.BatchInsertAPIGroup(datas); err != nil {
		return nil, err
	}
	return models.GetAPIGroupList(nil)
}

func EditAPIGroup(f *utils.FormAPIGroupEdit) (ret *models.APIGroup, err error) {
	var (
		data       = new(models.APIGroup)
		dataCasbin = new(models.RoleAPI)
		m          = map[string]interface{}{}
	)
	if err = data.Get(f.Id); err != nil {
		return nil, err
	}
	path := data.Path
	method := data.Method
	if len(f.ApiGroup) != 0 {
		data.ApiGroup = f.ApiGroup
	}
	if len(f.Description) != 0 {
		data.Description = f.Description
	}
	if len(f.Path) != 0 {
		m["v1"] = f.Path
		data.Path = f.Path
	}
	if len(f.Method) != 0 {
		data.Method = strings.ToLower(f.Method)
		m["v2"] = strings.ToLower(f.Method)
	}
	if len(m) > 0 {
		if err = dataCasbin.Update(path, method, m); err != nil {
			logServe.Errorln(err)
			err = nil
		}
	}
	if err = data.Update(f.Id, data); err != nil {
		return nil, err
	}

	ret = data
	return
}

func DelAPIGroup(id int64) (err error) {
	var (
		data       = new(models.APIGroup)
		dataCasbin = new(models.RoleAPI)
	)
	if err = data.Get(id); err != nil {
		return
	}

	if err = dataCasbin.DeleteByPathAndMethod(data.Path, data.Method); err != nil {
		logServe.Errorln(err)
	}

	return data.Delete(id)
}

// GetAPIGroupList
func GetAPIGroupList(f *utils.ApiListForm, isAll bool) (ret []*models.APIGroup, err error) {
	if !isAll {
		q, _err := query.NewQuery(f.Page, f.PageSize).ValidCond(f.ToMap())
		if _err != nil {
			err = _err
			return
		}
		if len(q.Cond) == 0 && q.PageNum == 0 && q.PageSize == 0 {
			q = nil
		}
		return models.GetAPIGroupList(q)
	} else {
		return models.GetAPIGroupList(nil)
	}
}
