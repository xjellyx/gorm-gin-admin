package utils

// FormRegister
type FormRegister struct {
	// 手机号
	Phone string ` json:"phone" form:"phone" binding:"required"`
	// 密码
	Password string ` json:"password" form:"password" binding:"required"`
	Code     string `json:"code" form:"code"` // 手机验证码
}

func (f *FormRegister) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	if len(f.Phone) == 0 {
		err = ErrFormParamInvalid.SetVars("phone")
		return
	}
	if len(f.Password) == 0 {
		err = ErrFormParamInvalid.SetVars("password")
		return
	}

	return
}

// LoginForm 登入参数
type LoginForm struct {
	Username string  `json:"username" form:"username" binding:"required"`
	Password string  `json:"password" form:"password" binding:"required"`
	DeviceId *string `json:"deviceId" form:"deviceId"`
	IP       string
}

func (f *LoginForm) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	return
}

// FormEditUser
type FormEditUser struct {
	Nickname *string `json:"nickname" form:"nickname"`
	Username *string `json:"username" form:"username"`
	Email    *string `json:"email" form:"email"`
	Phone    *string `json:"phone" form:"phone"`
	Sign     *string `json:"sign" form:"sign"`
}

func (f *FormEditUser) Valid() (ret map[string]interface{}, err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	ret = map[string]interface{}{}
	if f.Phone != nil && len(*f.Phone) == 0 {
		err = ErrFormParamInvalid
		return
	} else if f.Phone != nil && len(*f.Phone) != 0 {
		if len(RegPhoneNum.FindString(*f.Phone)) == 0 {
			err = ErrPhoneInvalid
			return
		}
		ret["phone"] = *f.Phone
	}
	if f.Email != nil && len(*f.Email) == 0 {
		err = ErrFormParamInvalid
		return
	} else if f.Email != nil && len(*f.Email) != 0 {
		if len(RegPhoneNum.FindString(*f.Email)) == 0 {
			err = ErrEmailInvalid
			return
		}
		ret["email"] = *f.Email
	}
	if f.Username != nil && len(*f.Username) == 0 {
		err = ErrFormParamInvalid
		return
	} else if f.Username != nil {
		ret["username"] = *f.Username
	}
	if f.Nickname != nil && len(*f.Nickname) == 0 {
		err = ErrFormParamInvalid
		return
	} else if f.Nickname != nil {
		ret["nickname"] = *f.Nickname
	}

	if f.Sign != nil {
		ret["sign"] = *f.Sign
	}
	return
}

// FormBankCard
type FormBankCard struct {
	Number    string `json:"number" from:"number" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required" `          // 用户姓名
	Bank      string `json:"bank" form:"bank" binding:"required" `          // 银行卡名称
	BankStart string `json:"bankStart" form:"bankStart" binding:"required"` // 开户银行名称
}

// Valid
func (f *FormBankCard) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	return
}

// FormIDCard
type FormIDCard struct {
	Name     string `json:"name" form:"name" binding:"required"`
	CardId   string `json:"cardId" from:"idCard" binding:"required" `    // 身份证号
	IssueOrg string `json:"issueOrg" from:"issueOrg" binding:"required"` // 身份证发证机关
	// Birthday    string `json:"birthday" from:"birthday" binding:"required"`       // 出生日期
	ValidPeriod string `json:"validPeriod" from:"validPeriod" binding:"required"` // 有效时期
	CardIdAddr  string `json:"cardIdAddr"from:"idCardAddr" binding:"required" `   // 身份证地址
	// Sex         int    `json:"sex" form:"sex" binding:"required"`
	Nation string `json:"nation" form:"nation" binding:"required"`
}

func (f *FormIDCard) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	if len(f.CardId) == 0 {
		err = ErrFormParamInvalid
		return
	} else {
		if len(RegIDCard.FindString(f.CardId)) == 0 {
			err = ErrIDCardInvalid
			return
		}
	}
	if len(f.IssueOrg) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if len(f.CardIdAddr) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if len(f.ValidPeriod) == 0 {
		err = ErrFormParamInvalid
		return
	}
	//if len(f.Birthday) == 0 {
	//	err = ErrFormParamInvalid
	//	return
	//}
	if len(f.Name) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if len(f.Nation) == 0 {
		err = ErrFormParamInvalid
		return
	}
	return
}

type FormAddr struct {
	Country  string  `json:"country" form:"country" binding:"required"`
	Province string  `json:"province" form:"province" binding:"required"`
	City     string  `json:"city" form:"city" binding:"required"`
	District *string `json:"district" form:"district"` // 普通用户个人资料填写的地址所在区域
	Address  *string `json:"address" form:"address"`   // 普通用户个人资料填写的详细地址
	ID       *int64  `json:"id"`
}

func (f *FormAddr) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}
	if len(f.Country) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if len(f.Province) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if len(f.City) == 0 {
		err = ErrFormParamInvalid
		return
	}
	if f.Address == nil {
		f.Address = new(string)
	}
	if f.District == nil {
		f.District = new(string)
	}
	return
}

// FormUserOnline
type FormUserOnline struct {
	LoginIp     string `json:"loginIp"`   // 用户登入ip
	LoginTime   int64  `json:"loginTime"` // 用户登入时间
	IsOnline    *bool  `json:"isOnline"`  // 是否在线
	OfflineTime int64  `json:"offlineTime"`
	Device      string `json:"device"`
}

func (f *FormUserOnline) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInvalid
		return
	}

	return
}

// AddUserForm
type AddUserForm struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
	Code     string `form:"code"`
}

// FormUserList
type FormUserList struct {
	ID          string `json:"id" form:"id"`
	Username    string `json:"username" form:"username" `
	CreatedTime string `json:"createdTime" form:"createdTime" `
	Status      string `json:"status" form:"status"`
	PageSize    int    `json:"pageSize" form:"pageSize" `
	PageNum     int    `json:"pageNum" form:"pageNum" `
}
