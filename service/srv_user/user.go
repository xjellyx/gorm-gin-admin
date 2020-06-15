package srv_user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olongfen/user_base/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	if err = u.GetUserByPhone(form.Phone); err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return
	} else if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {
		err = nil
	} else {
		err = fmt.Errorf("该手机号码已经注册")
		return
	}
	u.Phone = form.Phone
	if _d, _err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		u.LoginPasswd = string(_d)
	}
	u.Uid = uuid.NewV4().String()
	if err = u.InsertUserData(); err != nil {
		return
	}
	return u, nil
}
