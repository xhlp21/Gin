package datebles

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InitDB() *sql.DB {

	//打印连接数据库参数
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		"root", "root", "127.0.0.1:3306", "test", "utf8")

	//进行连接数据库
	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err.Error())
	} else {
		conn.SetConnMaxLifetime(7 * time.Second) //设置空闲时间，超过时间自动断开
		conn.SetMaxOpenConns(10)                 //同一时间内，最大连接限制
		conn.SetMaxIdleConns(10)                 //空闲时间内，最大连接限制
		return conn
	}
}
