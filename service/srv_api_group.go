package service

import (
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/query"
	"github.com/olongfen/user_base/utils"
)

// AddAPIGroup
func AddAPIGroup(f *utils.FormAPIGroupAdd) (ret *models.APIGroup, err error) {
	var (
		data = new(models.APIGroup)
	)
	data.ApiGroup = f.ApiGroup
	data.Description = f.Description
	data.Path = f.Path
	data.Method = f.Method
	// data.CreateUID = uid
	if err = data.InsertAPIGroup(); err != nil {
		return
	}
	ret = data
	return
}

func EditAPIGroup(f *utils.FromAPIGroupEdit) (ret *models.APIGroup, err error) {
	var (
		data = new(models.APIGroup)
	)

	if len(f.ApiGroup) != 0 {
		data.ApiGroup = f.ApiGroup
	}
	if len(f.Description) != 0 {
		data.Description = f.Description
	}
	if len(f.Path) != 0 {
		data.Path = f.Path
	}
	if len(f.Method) != 0 {
		data.Method = f.Method
	}
	if err = data.UpdateAPIGroup(f.Id, data); err != nil {
		return nil, err
	}

	ret = data
	return
}

func DelAPIGroup(id int64) (err error) {
	var (
		data = new(models.APIGroup)
	)
	return data.DeleteAPIGroup(id)
}

// GetAPIGroupList
func GetAPIGroupList(pageNum, pageSize int) (ret []*models.APIGroup, err error) {
	return models.GetAPIGroupList(query.NewQuery(pageNum, pageSize))
}
