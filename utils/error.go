package utils

import (
	"fmt"
	"github.com/olongfen/contrib"
	"strings"
)

const prefix = "user_base"

const (
	GORMPGUniqueErr  = `pq: duplicate key value violates unique constraint`
	GORMPGPrimaryErr = ``
)

// TransformGORMErr 转换数据库错误
func TransformGORMErr(gormErr error) (err error) {
	// 唯一错误处理
	if strings.Contains(gormErr.Error(), GORMPGUniqueErr) {
		s := strings.Split(gormErr.Error(), " ")
		err = contrib.NewError(0, fmt.Sprintf(`%s is exist`, s[len(s)-1]), prefix)
		return
	}
	return gormErr
}

var (
	ErrFormParamInvalid    = contrib.NewError(1, "form param invalid", prefix)
	ErrUserNotExist        = contrib.NewError(2, "user dose not exist", prefix)
	ErrActionNotAllow      = contrib.NewError(3, "action dose not allow", prefix)
	ErrNotRealNameVerified = contrib.NewError(4, "user dose not verified real name", prefix)
	ErrCapOfBankCard       = contrib.NewError(5, "cap of user bank card ", prefix)
	ErrUserEmailExist      = contrib.NewError(6, "user email exist", prefix)
	ErrUserVerified        = contrib.NewError(8, "user already verified", prefix)
	ErrUserAccountFroze    = contrib.NewError(7, "user account froze", prefix)
	ErrCapOfAddress        = contrib.NewError(9, "cap of user address", prefix)
	ErrUserNotOnline       = contrib.NewError(10, "user dose not online", prefix)
	ErrPayPwdNotSet        = contrib.NewError(11, "user pay password dose not set", prefix)
	ErrImageSizeToBig      = contrib.NewError(12, "image size to big", prefix)
	ErrImagePixelToBig     = contrib.NewError(13, "image pixel to big ", prefix)
	ErrIPAddressInvalid    = contrib.NewError(14, "ip address invalid", prefix)
	ErrInsertDataFailed    = contrib.NewError(15, "insert data failed", prefix)
	ErrGetDataFailed       = contrib.NewError(16, "get data failed", prefix)
	ErrUpdateDataFailed    = contrib.NewError(17, "update data failed", prefix)
	ErrDeleteDataFailed    = contrib.NewError(18, "delete data failed ", prefix)
	ErrUserIsVerified      = contrib.NewError(19, "the user has been authenticated by real name", prefix)
	ErrPhoneInvalid        = contrib.NewError(20, "phone number invalid", prefix)
	ErrEmailInvalid        = contrib.NewError(21, "email invalid", prefix)
	ErrIDCardInvalid       = contrib.NewError(22, "id card number invalid", prefix)
	ErrCaptchaVerifyFail   = contrib.NewError(23, "captcha verify failed ", prefix)
)
