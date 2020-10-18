package user

import (
	"encoding/json"
	"fmt"
	"go-projects-server/internal/app/model"

	"go.uber.org/zap"
)

// User 用户结构体
type User struct {
	model.BaseModel
	UserID   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Gender   int8   `json:"gender" db:"gender"`
}

// GetUser 根据ID获取用户
func (u *User) GetUser() {
	db := u.GetDB()
	sqlStr := `select id,user_id,username from user where id=?`
	if err := db.Get(u, sqlStr, u.ID); err != nil {
		fmt.Println("获取用户异常，err:", err)
		zap.L().Error("获取用户异常，err:", zap.Error(err))
		return
	}
	res, _ := json.Marshal(u)
	fmt.Println("获取用户成功:", string(res))
}
