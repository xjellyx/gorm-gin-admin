package model

// AddressDetail 地址信息
type AddressDetail struct {
	ID uint64 `gorm:"primary_key"`
	// Aid      string `json:"aid" gorm:"primary_key;size:36;index"`
	Uid      string `json:"uid" gorm:"index;size:36"`
	Country  string `json:"country" gorm:"index;size:128"`
	Province string `json:"province" gorm:"index;size:36"`
	City     string `json:"city" gorm:"index;size:36"`
	District string `json:"district;size:128"` // 普通用户个人资料填写的地址所在区域
	Address  string `json:"address;size:128"`  // 普通用户个人资料填写的详细地址
	TimeData
}

// TableName
func (a *AddressDetail) TableName() string {
	return "address_detail"
}
