package models

import "gorm.io/gorm"

const (
	UserStatusRegister = iota // 0 注册状态
	UserStatusLogin           // 1 登录状态
	UserStatusLogout          // 2 登出状态
	UserStatusLock            // 3 封锁状态
)

// UserBase 用户信息
type UserBase struct {
	gorm.Model
	Uid         string `json:"uid" gorm:"type:varchar(36); unique_index"`
	Username    string `json:"username" gorm:"type:varchar(16); unique_index"`
	Phone       string `json:"phone" gorm:"type:varchar(11); unique_index"`
	LoginPasswd string `json:"-" gorm:"type:varchar(16)"`
	PayPasswd   string `json:"-" gorm:"type:varchar(16)"`
	Email       string `json:"email" gorm:"type:varchar(32)"`
	Nickname    string `json:"nickname" gorm:"type:varchar(12)"`
	HeadIcon    string `json:"headIcon"`
	Sign        string `json:"sign" gorm:"type:varchar(256)"`
	Status      int    `json:"status"`

	// 外键
	UserCard UserCard `json:"userCard" gorm:"foreignkey:ID"`
}

// InsertUserData 插入一条数据
func (u *UserBase) InsertUserData() (err error) {
	return db.Create(u).Error
}

func (u *UserBase) UpdateUserInterface(uid string, data interface{}) error {
	return db.Model(u).Where("uid = ?", uid).Updates(data).Error
}

// UpdateUser 更新数据
func (u *UserBase) UpdateUser(uid string) error {
	return db.Model(u).Where("uid = ?", uid).Updates(u).Error
}

// UpdateUserOneColumn 更新一个字段
func (u *UserBase) UpdateUserOneColumn(uid string, column string, val interface{}) error {
	return db.Model(u).Where("uid = ?", uid).Update(column, val).Error
}

// GetUserById 通过id获取数据
func (u *UserBase) GetUserById(id int64) error {
	return db.First(u, "id = ?", id).Error
}

// GetUserByUId 通过id获取数据
func (u *UserBase) GetUserByUId(uid string) error {
	return db.First(u, "uid = ?", uid).Error
}

// GetUserByUsername 通过username获取用户信息
func (u *UserBase) GetUserByUsername(username string) error {
	return db.First(u, "username = ?", username).Error
}

// GetUserByPhone 通过phone获取信息
func (u *UserBase) GetUserByPhone(phone string) error {
	return db.First(u, "phone = ?", phone).Error
}

// GetUserList 获取用户列表
func GetUserList(pageNum int, pageSize int, cond interface{}) (ret []*UserBase, err error) {
	if err = db.Where(cond).Offset(pageNum).Limit(pageSize).Find(&ret).Error; err != nil {
		return
	}
	return
}

// GetUserTotal 获取总数
func GetUserTotal(cond interface{}) (int64, error) {
	var count int64
	if err := db.Model(&UserBase{}).Where(cond).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
