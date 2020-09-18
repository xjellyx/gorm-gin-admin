package utils

import (
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"time"
)

// BehaviorQueryForm
type BehaviorQueryForm struct {
	Page      int        `json:"page" form:"page" `
	PageSize  int        `json:"pageSize" form:"pageSize" `
	Username  string     `json:"username" form:"username"`
	Method    string     `json:"method" form:"method"` // 请求方式
	StartTime *time.Time `json:"startTime" form:"startTime"`
	EndTime   *time.Time `json:"endTime" form:"endTime"`
}

func (f *BehaviorQueryForm) ToMap() (ret map[string]interface{}) {
	data := make(map[string]interface{})
	if len(f.Username) != 0 {
		data["username"] = f.Username
	}
	if len(f.Method) != 0 {
		data["method"] = f.Method
	}
	if f.StartTime != nil {
		data["created_at"+query.TagValGte] = f.StartTime.Local()
	}
	if f.EndTime != nil {
		data["created_at"+query.TagValLte] = f.EndTime.Local()
	}
	return data
}

// BehaviorRemoveForm
type BehaviorRemoveForm struct {
	Ids []int64 `json:"ids" form:"ids"`
}
