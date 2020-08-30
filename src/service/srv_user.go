package service

import (
	"fmt"
	"github.com/olongfen/user_base/src/models"
	"github.com/olongfen/user_base/src/pkg/query"
	"github.com/olongfen/user_base/src/utils"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// AddUser 添加哟个用户
func AddUser(form *utils.AddUserForm) (ret *models.UserBase, err error) {
	var (
		u = new(models.UserBase)
	)
	if len(utils.RegPhoneNum.FindString(form.Phone)) == 0 {
		err = utils.ErrPhoneInvalid
		return
	}
	u.Phone = form.Phone
	u.Username = form.Username
	if _d, _err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		u.LoginPwd = string(_d)
	}
	u.Uid = uuid.NewV4().String()
	if err = u.Insert(); err != nil {
		return
	}
	return u, nil
}

// EditUserBySelf 修改用户信息
func EditUserBySelf(uid string,form *utils.FormEditUser) (ret *models.UserBase, err error) {
	var (
		dataMap map[string]interface{}
		data    = &models.UserBase{}
	)
	if dataMap, err = form.Valid(); err != nil {
		return
	}
	if err = data.GetByUId(uid);err!=nil{
		return
	}

	if err = data.Update(form.Uid, dataMap); err != nil {
		return nil, err
	}

	//
	ret = data
	return
}



// EditUserByRole 修改用户信息
func EditUserByRole(uid string,form *utils.FormEditUser) (ret *models.UserBase, err error) {
	var (
		dataMap map[string]interface{}
		role = new(models.UserBase)
		data    = &models.UserBase{}
	)
	if dataMap, err = form.Valid(); err != nil {
		return
	}
	if err = role.GetByUId(uid);err!=nil{
		return
	}

	if err = data.GetByUId(form.Uid); err != nil {
		return
	}
	if data.Role >= role.Role{
		err = utils.ErrActionNotAllow
		return
	}
	if err = data.Update(form.Uid, dataMap); err != nil {
		return nil, err
	}

	//
	ret = data
	return
}

// ChangePwd 修改密码
func ChangePwd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetByUId(uid); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPwd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateOne(uid, "login_pwd", string(_d))
	}
}

// ChangePayPwd 修改密码
func ChangePayPwd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetByUId(uid); err != nil {
		return
	}
	if len(data.PayPwd) == 0 {
		err = utils.ErrPayPwdNotSet
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.PayPwd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateOne(uid, "pay_pwd", string(_d))
	}
}

// SetUserPayPwd
func SetUserPayPwd(uid string, pwd string) (err error) {
	u := new(models.UserBase)
	if err = u.GetByUId(uid); err != nil {
		return
	}
	if len(u.PayPwd) > 0 {
		err = utils.ErrActionNotAllow
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return u.UpdateOne(uid, "pay_pwd", string(_d))
	}
}

// GetUserList
func GetUserList(form *utils.FormUserList) (ret []*models.UserBase, err error) {
	cond := map[string]interface{}{}
	if form.Username != "" {
		cond["$and$username"] = utils.SpiltInterfaceList(form.Username, ",")
	}
	if form.Status != "" {
		cond["$and$status"] = utils.SpiltInterfaceList(form.Status, ",")
	}
	if form.ID != "" {
		var d []interface{}
		for _, v := range strings.Split(form.ID, ",") {
			d = append(d, v)
		}
		cond["$and$id"] = d

	}
	if form.CreatedTime != "" {
		cond["$and$created_at"] = utils.SpiltInterfaceList(form.CreatedTime, ",")
	}
	var (
		q *query.Query
	)
	if q, err = query.NewQuery(form.PageNum, form.PageSize).ValidCond(cond); err != nil {
		return
	}
	return models.GetUserList(q)
}

func GetUserCount(uid string)(ret int64,err error)  {
	data:=new(models.UserBase)
	if err = data.GetByUId(uid);err!=nil{
		return
	}
	return models.GetUserTotal(fmt.Sprintf(`role< %v`,data.Role))
}