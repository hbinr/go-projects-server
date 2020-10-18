package model

import (
	"go-projects-server/pkg/db/mysql"
	"time"

	"github.com/jmoiron/sqlx"
)

// BaseModel 基础model
type BaseModel struct {
	ID         int64     `json:"id" db:"id"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

// GetDB 获取sqlx 数据库引擎示例
func (base *BaseModel) GetDB() *sqlx.DB {
	return mysql.DB
}
