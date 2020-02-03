package base

import (
	"fmt"
	"github.com/olefen/contrib"
	"strings"
)

const prefix = "userDetail"

const (
	GORMPGUniqueErr  = `pq: duplicate key value violates unique constraint`
	GORMPGPrimaryErr = ``
)

// TransformGORMErr 转换数据库错误
func TransformGORMErr(gormErr error) (err error) {
	// 唯一错误处理
	if strings.Contains(gormErr.Error(), GORMPGUniqueErr) {
		s := strings.Split(gormErr.Error(), " ")
		err = contrib.NewError(2, fmt.Sprintf(`%s is exist`, s[len(s)-1]), prefix)
		return
	}
	return gormErr
}

var (
	ErrFormParamInValid    = contrib.NewError(1, "form param invalid", prefix)
	ErrUserNotExist        = contrib.NewError(2, "user not exist", prefix)
	ErrActionNotAllow      = contrib.NewError(3, "action not allow", prefix)
	ErrNotRealNameVerified = contrib.NewError(4, "user not verified real name", prefix)
	ErrCapOfBankCard       = contrib.NewError(5, "cap of user bank card ", prefix)
	ErrUserEmailExist      = contrib.NewError(6, "user email exist", prefix)
	ErrUserVerified        = contrib.NewError(8, "user already verified", prefix)
	ErrUserAccountFroze    = contrib.NewError(7, "user account froze", prefix)
)
