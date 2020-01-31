package base

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
		err = ErrFormParamInValid
		return
	}
	if len(f.Phone) == 0 {
		err = ErrFormParamInValid.SetVars("phone")
		return
	}
	if len(f.Password) == 0 {
		err = ErrFormParamInValid.SetVars("password")
		return
	}

	return
}

// LoginForm 登入参数
type LoginForm struct {
	Username string  `json:"username" form:"username" binding:"required"`
	Password string  `json:"password" form:"password" binding:"required"`
	IP       string  `json:"ip" form:"ip" binding:"required"`
	DeviceId *string `json:"deviceId" form:"deviceId"`
}

func (f *LoginForm) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInValid
		return
	}
	return
}

// UpdateUserProfile
type UpdateUserProfile struct {
	Nickname *string `json:"nickname" form:"nickname"`
	Username *string `json:"username" form:"username"`
	LocNum   *string `json:"locNum" form:"locNum"`
	Phone    *string `json:"phone" form:"phone"`
	HeadIcon *string `json:"headIcon" form:"headIcon"`
	Sign     *string `json:"sign" form:"sign"`
}

func (f *UpdateUserProfile) Valid() (ret map[string]interface{}, err error) {
	if f == nil {
		err = ErrFormParamInValid
		return
	}
	ret = map[string]interface{}{}
	if f.Phone != nil && len(*f.Phone) == 0 {
		if f.LocNum == nil {
			f.LocNum = new(string)
			*f.LocNum = "86"
		}
		err = ErrFormParamInValid
		return
	} else if f.Phone != nil && f.LocNum != nil {
		ret["phone"] = *f.Phone
		ret["loc_num"] = *f.LocNum
	}

	if f.Username != nil && len(*f.Username) == 0 {
		err = ErrFormParamInValid
		return
	} else if f.Username != nil {
		ret["username"] = *f.Username
	}
	if f.Nickname != nil && len(*f.Nickname) == 0 {
		err = ErrFormParamInValid
		return
	} else if f.Nickname != nil {
		ret["nick_name"] = *f.Nickname
	}
	if f.HeadIcon != nil && len(*f.HeadIcon) == 0 {
		err = ErrFormParamInValid
		return
	} else if f.HeadIcon != nil {
		ret["head_icon"] = *f.HeadIcon
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
		err = ErrFormParamInValid
		return
	}
	return
}

// FormIDCard
type FormIDCard struct {
	Name        string `json:"name" form:"name" binding:"required"`
	IDCard      string `json:"idCard" from:"idCard" binding:"required" `          // 身份证号
	IssueOrg    string `json:"issueOrg" from:"issueOrg" binding:"required"`       // 身份证发证机关
	Birthday    string `json:"birthday" from:"birthday" binding:"required"`       // 出生日期
	ValidPeriod string `json:"validPeriod" from:"validPeriod" binding:"required"` // 有效时期
	IDCardAddr  string `json:"idCardAddr"from:"idCardAddr" binding:"required" `   // 身份证地址
	Sex         int    `json:"sex" form:"sex" binding:"required"`
	Nation      string `json:"nation" form:"nation" binding:"required"`
}

func (f *FormIDCard) Valid() (err error) {
	if f == nil {
		err = ErrFormParamInValid
		return
	}
	if len(f.IDCard) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.IssueOrg) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.IDCardAddr) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.ValidPeriod) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.Birthday) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.Name) == 0 {
		err = ErrFormParamInValid
		return
	}
	if len(f.Nation) == 0 {
		err = ErrFormParamInValid
		return
	}
	return
}
