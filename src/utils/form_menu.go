package utils

import "github.com/olongfen/gorm-gin-admin/src/pkg/body"

// FormAddMenu
type FormAddMenu struct {
	Name     string `form:"name" binding:"required"`
	ParentId int   `form:"parentId" binding:"required"`
	Path     string `form:"path" binding:"required"`
	Component string `form:"component" binding:"required"`
	Meta      map[string]interface{} `form:"meta" binding:"required"`
	Sort      int    `form:"sort"`
}

type FormUpdateMenu struct {
	Id int `form:"id" binding:"required"`
	Name     *string `form:"name"`
	ParentId *int   `form:"parentId" `
	Path     *string `form:"path"`
	Component *string `form:"component" `
	Meta       body.Body `form:"meta"`
	Sort      *int    `form:"sort"`
}

func (f *FormUpdateMenu)ToMap() (ret map[string]interface{})  {
	ret = map[string]interface{}{}
	ret["id"] = f.Id
	if f.Path!=nil{
		ret["path"] = *f.Path
	}
	if f.Name!=nil{
		ret["name"] = *f.Name
	}
	if f.Component!=nil{
		ret["component"] = *f.Component
	}
	if f.ParentId!=nil{
		ret["parent_id"] = *f.ParentId
	}
	if f.Sort!=nil{
		ret["sort"] = *f.Sort
	}
	if f.Meta!=nil && len(f.Meta)>0{
		ret["meta"] = f.Meta
	}
	return
}