package user

import (
	"fmt"
	"go-projects-server/internal/app/dto"
	"go-projects-server/internal/app/model/user"
)

// Register 用户注册
func Register(data *dto.RegisterStruct) error {
	var user user.User
	// if err := gconv.Struct(data, &user); err != nil {
	// 	return err
	// }
	user.ID = 1
	fmt.Println("data...", user)
	user.GetUser()
	return nil
}
