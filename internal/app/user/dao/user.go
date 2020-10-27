package dao

import (
	"errors"
	"go-projects-server/pkg/log"

	"gorm.io/gorm"
)

func (r *UserDao) Insert(user *UserModel) (err error) {
	return r.DB.Create(&user).Error
}

func (r *UserDao) Delete(id int64) bool {
	return r.DB.Where("id = ?", id).Delete(&UserModel{}).RowsAffected > 0
}

func (r *UserDao) Update(user *UserModel) error {
	return r.DB.Where("id = ?").Updates(&user).Error
}

func (r *UserDao) SlectByName(userName string) (*UserModel, error) {
	var user UserModel
	if err := r.DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) SelectById(id int64) (*UserModel, error) {
	var (
		user UserModel
		err  error
	)
	if err = r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("gorm.ErrRecordNotFound", err)
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}
