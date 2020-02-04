package model

import userBase "github.com/olefen/userDetail"

// AddressDetail 地址信息
type AddressDetail struct {
	ID uint64 `gorm:"primary_key"`
	// Aid      string `json:"aid" gorm:"primary_key;size:36;index"`
	Uid      string `json:"uid" gorm:"index;size:36"`
	Country  string `json:"country" gorm:"index;size:128"`
	Province string `json:"province" gorm:"index;size:36"`
	City     string `json:"city" gorm:"index;size:36"`
	District string `json:"district" gorm:"size:128"` // 普通用户个人资料填写的地址所在区域
	Address  string `json:"address" gorm:"size:128"`  // 普通用户个人资料填写的详细地址
	TimeData
}

// TableName
func (a *AddressDetail) TableName() string {
	return "address_detail"
}

func NewAddress(in *AddressDetail) (ret *AddressDetail) {
	if in != nil {
		ret = in
	} else {
		ret = new(AddressDetail)
	}
	return
}

// PubAddressAdd 添加用户地址
func PubAddressAdd(uid string, f *userBase.FormAddr) (ret *AddressDetail, err error) {
	var (
		data     *AddressDetail
		dataUser *UserDetail
	)
	if err = f.Valid(); err != nil {
		return
	}
	if dataUser, err = PubUserGet(uid); err != nil {
		return
	}
	if Database.Model(dataUser).Association("Addr").Count() == 6 {
		err = userBase.ErrCapOfAddress
		return
	}
	data = NewAddress(&AddressDetail{
		Uid:      uid,
		Country:  f.Country,
		Province: f.Province,
		City:     f.City,
		District: *f.District,
		Address:  *f.Address,
	})

	if err = Database.Model(&AddressDetail{}).Create(data).Error; err != nil {
		return
	}

	ret = data
	return
}

// PubAddressGetList
func PubAddressGetList(uid string) (ret []*AddressDetail, err error) {
	var (
		data     []*AddressDetail
		dataUser *UserDetail
	)
	if dataUser, err = PubUserGet(uid); err != nil {
		return
	}
	if err = Database.Model(dataUser).Association("Addr").Find(&data).Error; err != nil {
		return
	}

	ret = data
	return
}

// PubAddressGetOne
func PubAddressGetOne(id int64) (ret *AddressDetail, err error) {
	ret = new(AddressDetail)
	if err = Database.Model(&AddressDetail{}).First(ret, "id=?", id).Error; err != nil {
		return
	}
	return
}

// PubAddressUpdate
func PubAddressUpdate(uid string, f *userBase.FormAddr) (ret *AddressDetail, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	if f.ID == nil || f.ID != nil && *f.ID < 0 {
		err = userBase.ErrFormParamInvalid
		return
	}
	var (
		data *AddressDetail
	)

	if data, err = PubAddressGetOne(*f.ID); err != nil {
		return
	}
	if data.Uid != uid {
		err = userBase.ErrActionNotAllow
		return
	}
	data.Country = f.Country
	data.Province = f.Province
	data.City = f.City
	data.District = *f.District
	data.Address = *f.Address
	if err = Database.Model(&AddressDetail{}).Where("id=?", data.ID).Updates(data).Error; err != nil {
		return
	}

	ret = data
	return
}

// PubAddressDelete
func PubAddressDelete(uid string, id int64) (err error) {

	var (
		data *AddressDetail
	)

	if data, err = PubAddressGetOne(id); err != nil {
		return
	}
	if data.Uid != uid {
		err = userBase.ErrActionNotAllow
		return
	}

	if err = Database.Delete(&AddressDetail{}, "id=?", data.ID).Error; err != nil {
		return
	}

	return
}

//
