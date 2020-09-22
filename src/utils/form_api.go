package utils

import "github.com/olongfen/gorm-gin-admin/src/pkg/query"

// ApiListForm
type ApiListForm struct {
	Page        int    `json:"page" form:"page" `
	PageSize    int    `json:"pageSize" form:"pageSize" `
	Path        string `json:"path" form:"path" `               // api路径
	Method      string `json:"method" form:"method"`            // 请求方式
	Description string `json:"description" form:"description" ` // 中文描述
	ApiGroup    string `json:"apiGroup" form:"apiGroup"`        // 组名
}

func (f *ApiListForm) ToMap() map[string]interface{} {
	cond := map[string]interface{}{}
	if len(f.Path) > 0 {
		cond["path"] = f.Path
	}
	if len(f.ApiGroup) > 0 {
		cond["api_group"] = f.ApiGroup
	}
	if len(f.Method) > 0 {
		cond["method"] = f.Method
	}
	if len(f.Description) > 0 {
		cond["description"+query.TagValLike] = "%" + f.Description + "%"
	}
	return cond
}

// FormAPIGroupAdd 添加api数据
type FormAPIGroupAdd struct {
	Path        string `json:"path" form:"path" binding:"required"`               // api路径
	Method      string `json:"method" form:"method" binding:"required"`           // 请求方式
	Description string `json:"description" form:"description" binding:"required"` // 中文描述
	ApiGroup    string `json:"apiGroup" form:"apiGroup" binding:"required"`       // 组名
}

// FormAPIGroupEdit 添加api数据
type FormAPIGroupEdit struct {
	Id          int64  `json:"id" form:"id" binding:"required"`
	Path        string `json:"path" form:"path" binding:"required" `              // api路径
	Method      string `json:"method" form:"method" binding:"required"`           // 请求方式
	Description string `json:"description" form:"description" binding:"required"` // 中文描述
	ApiGroup    string `json:"apiGroup" form:"apiGroup" binding:"required"`       // 组名
}
