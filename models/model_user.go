package models

import "gorm.io/gorm"

// UserBase 用户信息
type UserBase struct {
	gorm.Model
	Uid         string
	Username    string
	Phone       string
	LoginPasswd string
	PayPasswd   string
	Email       string
	NickName    string
	Status      int
}

// InsertUserData 插入一条数据
func (u *UserBase) InsertUserData() (err error) {
	return db.Create(u).Error
}

// UpdateUser 更新数据
func (u *UserBase) UpdateUser() error {
	return db.Updates(u).Error
}

// UpdateUserOneColumn 更新一个字段
func (u *UserBase) UpdateUserOneColumn(column string, val interface{}) error {
	return db.Model(u).Update(column, val).Error
}

// GetUserById 通过id获取数据
func (u *UserBase) GetUserById(id int64) error {
	return db.First(u, "id = ?", id).Error
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
