package srv_card

import (
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
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
	if err = u.GetUserByUId(uid); err != nil {
		return
	}
	// 已经实名认证，返回
	if err = data.GetIDCardByUid(uid); err == nil && data.ID != 0 {
		err = utils.ErrUserIsVerified
		return
	} else if err != nil {
		return
	}
	data.Uid = uid
	data.Name = f.Name
	data.CardIdAddr = f.CardIdAddr
	data.IssueOrg = f.IssueOrg
	data.Birthday = f.Birthday
	data.Sex = f.Sex
	data.CardId = f.CardId
	data.Nation = f.Nation
	data.ValidPeriod = f.ValidPeriod
	if err = data.InsertIDCard(); err != nil {
		return
	}

	//
	ret = data
	return
}
