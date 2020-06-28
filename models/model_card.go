package models

import "gorm.io/gorm"

// UserCard
type UserCard struct {
	gorm.Model
	//
	Name        string `json:"name" `
	IDCard      string `json:"idCard" gorm:"unique_index;type:varchar(18)" ` // 身份证号
	IssueOrg    string `json:"issueOrg" `                                    // 身份证发证机关
	Birthday    string `json:"birthday" gorm:"type:varchar(12)"`             // 出生日期
	ValidPeriod string `json:"validPeriod"  gorm:"type:varchar(12)"`         // 有效时期
	IDCardAddr  string `json:"idCardAddr"  gorm:"type:varchar(64)"`          // 身份证地址
	Sex         int    `json:"sex" `
	Nation      string `json:"nation" `
}
