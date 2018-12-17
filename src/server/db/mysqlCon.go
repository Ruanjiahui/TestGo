package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

//数据库配置
const (
	userName = "root"
	password = "pama123456"
	ip       = "192.168.101.10"
	port     = "3306"
	dbName   = "debug"
)

//Db数据库连接池
var DB *sql.DB

//错误日志
var err error

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err = sql.Open("mysql", path)
	//是否打开数据库成功
	if err != nil {
		fmt.Println("open database fail : ", err)
		return
	}
	//设置数据库最大连接数
	DB.SetMaxOpenConns(10)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(10 * time.Minute)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail : ", err)
		return
	}
	fmt.Println("connnect success")
}
