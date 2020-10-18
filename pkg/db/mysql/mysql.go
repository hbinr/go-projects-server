package mysql

import (
	"fmt"
	"go-projects-server/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Init 初始化mysql连接
func Init(cfg *setting.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = DB.Close()
}
