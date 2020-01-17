package model

// BankCard 银行卡信息
type BankCard struct {
	Number    string `json:"number" gorm:"not null;primary_key;size:19;index"`
	Name      string `json:"name" gorm:"index;size:64"`       // 用户姓名
	Bank      string `json:"bank" gorm:"index;size:64"`       // 银行卡名称
	BankStart string `json:"bankStart" gorm:"index;size:128"` // 开户银行名称
	Uid       string `json:"uid"gorm:"index;size:36"`
	TimeData
}

// TableName
func (b *BankCard) TableName() string {
	return "bank_card"
}
