package utils

import "time"

// BehaviorQueryForm
type BehaviorQueryForm struct {
	Page      int    `json:"page" form:"page" `
	PageSize  int    `json:"pageSize" form:"pageSize" `
	Username  string `json:"username" form:"username"`
	Method    string `json:"method" form:"method"` // 请求方式
	CreatedAt int64  `json:"createdAt" form:"createdAt"`
}

func (f *BehaviorQueryForm) ToMap() (ret map[string]interface{}) {
	data := make(map[string]interface{})
	if len(f.Method) != 0 {
		data["username"] = f.Username
	}
	if len(f.Method) != 0 {
		data["method"] = f.Method
	}
	if f.CreatedAt > 0 {
		data["created_at"] = time.Unix(f.CreatedAt, f.CreatedAt)
	}
	return data
}

// BehaviorRemoveForm
type BehaviorRemoveForm struct {
	Ids []int64 `json:"ids" form:"ids"`
}
