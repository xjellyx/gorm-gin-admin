package middleware

import "github.com/olongfen/gorm-gin-admin/src/models"

func CheckUserLogin(u *models.UserBase) bool  {
	return u.Status==models.UserStatusLogin
}
