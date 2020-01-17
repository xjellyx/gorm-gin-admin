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
