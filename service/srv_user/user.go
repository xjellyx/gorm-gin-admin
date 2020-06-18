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
		data    = &models.UserBase{Uid: uid}
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
	if err = data.UpdateUser(); err != nil {
		return nil, err
	}

	//
	ret = data
	return
}
