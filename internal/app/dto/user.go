package dto

import "go-projects-server/internal/app/model/user"

// RegisterStruct 用户注册结构体
type RegisterStruct struct {
	Username   string `json:"userName"`
	Password   string `json:"passWord"`
	RePassword string `json:"rePassword"`
	NickName   string `json:"nickName"`
}

// UserResponse 响应
type UserResponse struct {
	User user.User `json:"user"`
}
