package service

import (
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
	"strconv"
	"time"
)

// AddIDCard
func AddIDCard(uid string, f *utils.FormIDCard) (ret *models.UserCard, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		u    = new(models.UserBase)
		data = new(models.UserCard)
	)
	if err = u.GetByUId(uid); err != nil {
		return
	}
	// 已经实名认证，返回
	if err = data.GetByUid(uid); err == nil && data.ID != 0 {
		err = utils.ErrUserIsVerified
		return
	} else if err != nil {
		return
	}
	data.Uid = uid
	data.Name = f.Name
	data.CardIdAddr = f.CardIdAddr
	data.IssueOrg = f.IssueOrg
	data.Birthday = string(f.CardId[6:10]) + "/" + string(f.CardId[10:12]) + "/" + string(f.CardId[12:14])
	year, _ := strconv.ParseInt(f.CardId[6:10], 10, 64)
	data.Sex = time.Now().Year() - int(year)
	data.CardId = f.CardId
	data.Nation = f.Nation
	data.ValidPeriod = f.ValidPeriod
	if err = data.Insert(); err != nil {
		return
	}

	//
	ret = data
	return
}

func IDCardList(pageNum, pageSize int, cond interface{}) {

}
