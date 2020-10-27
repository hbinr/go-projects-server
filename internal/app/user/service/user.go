package service

import (
	"go-projects-server/internal/app/user/dao"
	"go-projects-server/internal/app/user/model"

	"github.com/gogf/gf/util/gconv"
)

func (r *UserService) Insert(user *model.UserDto) (err error) {
	var u dao.UserModel
	if err := gconv.Struct(user, &u); err != nil {
		return err
	}
	return r.Dao.Insert(&u)
}

func (r *UserService) Delete(id int64) bool {
	return r.Dao.Delete(id)
}

func (r *UserService) Update(user *model.UserDto) error {
	var uModel model.User
	if err := gconv.Struct(user, &uModel); err != nil {
		return err
	}
	return r.Dao.Update(&uModel)
}

func (r *UserService) SlectByName(userName string) (*model.UserDto, error) {
	uModel, err := r.Dao.SlectByName(userName)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}

func (r *UserService) SelectById(id int64) (*model.UserDto, error) {
	uModel, err := r.Dao.SelectById(id)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}
