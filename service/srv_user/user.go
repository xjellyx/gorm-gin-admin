package srv_user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olongfen/userDetail/models"
)

// AddUserForm
type AddUserForm struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
	Code     string `form:"code"`
}

// AddUser 添加哟个用户
func AddUser(form *AddUserForm) (ret *models.UserBase, err error) {
	var (
		u = new(models.UserBase)
	)
	if err = u.GetUserByPhone(form.Phone); err != nil && !gorm.IsRecordNotFoundError(err) {
		return
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		err = nil
	} else {
		err = fmt.Errorf("该手机号码已经注册")
		return
	}
	u.Phone = form.Phone
	u.LoginPasswd = form.Password
	if err = u.InsertUserData(); err != nil {
		return
	}
	return u, nil
}
