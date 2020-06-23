package srv_user

import (
	"encoding/json"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// AddUser 添加哟个用户
func AddUser(form *utils.AddUserForm) (ret *models.UserBase, err error) {
	var (
		u = new(models.UserBase)
	)

	u.Phone = form.Phone
	u.Username = form.Phone
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

// EditUser 修改用户信息
func EditUser(uid string, form *utils.FormEditUser) (ret *models.UserBase, err error) {
	var (
		dataMap map[string]interface{}
		data    = &models.UserBase{}
	)
	if dataMap, err = form.Valid(); err != nil {
		return
	}
	//
	if _d, _err := json.Marshal(dataMap); _err != nil {
		return
	} else {
		if err = json.Unmarshal(_d, data); err != nil {
			return
		}
	}
	if err = data.UpdateUser(uid); err != nil {
		return nil, err
	}

	//
	ret = data
	return
}

// ChangePasswd 修改密码
func ChangePasswd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetUserByUId(uid); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPasswd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateUserOneColumn(uid, "login_passwd", string(_d))
	}
}

// ChangePayPasswd 修改密码
func ChangePayPasswd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetUserByUId(uid); err != nil {
		return
	}
	if len(data.PayPasswd) == 0 {
		err = utils.ErrPayPasswdNotSet
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.PayPasswd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateUserOneColumn(uid, "pay_passwd", string(_d))
	}
}
