package service

import (
	"go-projects-server/internal/app/user/dao"
	"go-projects-server/internal/app/user/model"

	"github.com/google/wire"
)

// 验证接口是否实现
var _ IUserService = (*UserService)(nil)

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)))

//func NewUserService(db *gorm.DB) IUserService {
//	return &UserService{
//		Dao: dao.NewUserDao(db),
//	}
//}

type IUserService interface {
	Insert(user *model.UserDto) error
	Delete(int64) bool
	Update(user *model.UserDto) error
	SelectById(id int64) (*model.UserDto, error)
	SlectByName(userName string) (*model.UserDto, error)
}

type UserService struct {
	Dao dao.IUserDao
}
