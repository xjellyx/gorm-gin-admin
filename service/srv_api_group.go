package service

import (
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
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
		data.Method = v.Method

		datas = append(datas, data)
	}
	if err = models.BatchInsertAPIGroup(datas); err != nil {
		return nil, err
	}
	return models.GetAPIGroupList()
}

func EditAPIGroup(f *utils.FormAPIGroupEdit) (ret *models.APIGroup, err error) {
	var (
		data       = new(models.APIGroup)
		dataCasbin = new(models.RuleAPI)
		m          = map[string]interface{}{}
	)
	if err = data.GetAPIGroup(f.Id); err != nil {
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
		data.Method = f.Method
		m["v2"] = f.Method
	}
	if len(m) > 0 {
		if err = dataCasbin.UpdateRuleAPI(path, method, m); err != nil {
			return nil, err
		}
	}
	if err = data.UpdateAPIGroup(f.Id, data); err != nil {
		return nil, err
	}

	ret = data
	return
}

func DelAPIGroup(id int64) (err error) {
	var (
		data       = new(models.APIGroup)
		dataCasbin = new(models.RuleAPI)
	)
	if err = data.GetAPIGroup(id); err != nil {
		return
	}

	if err = dataCasbin.DeleteRuleAPIData(data.Path, data.Method); err != nil {
		return err
	}

	return data.DeleteAPIGroup(id)
}

// GetAPIGroupList
func GetAPIGroupList() (ret []*models.APIGroup, err error) {
	return models.GetAPIGroupList()
}
