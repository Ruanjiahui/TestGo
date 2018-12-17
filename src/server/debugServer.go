package main

import (
	"errors"
	"server/Http"
	"server/db"
)

const PI = 3.14

var PO = "goland"

type PK string

//声明结构
type User struct {
	name    string
	age     int
	address string
}

//声明接口
type golang interface {
	call() string
	unCall() int
}

func (user User) call() (string, error) {
	return "address=" + user.address + " age=" + (string)(user.age) + " name=" + user.name, errors.New("错误日志")
}

func (user User) unCall() string {
	return "address=" + user.address + " age=" + (string)(user.age) + " name=" + user.name
}

//定义一个bool类型
type b bool

func main() {
	//config.Loadconfig()
	//
	//fmt.Println("Hello word!")
	//
	//user := new(User)
	//
	//user.address = "广东省深圳市龙岗区"
	//user.age = 18
	//user.name = "帅的一比"
	//
	//var msg string
	//var err error
	//
	//msg , err = user.call()
	//fmt.Println(msg , err)
	//fmt.Println(user.unCall())
	//
	//i := 0
	//for i = 0; i < 10; i++ {
	//	fmt.Println("i = ", i)
	//}
	//
	//var aa = 20
	//var ip *int
	//
	//ip = &aa
	//
	//fmt.Printf("%x\n", &aa)
	//fmt.Printf("%x\n", ip)
	//fmt.Println(*ip)
	//
	//var a int
	//var b = 2
	//
	////var bl bool
	//
	//for a = 2; a < 3; a++ {
	//	b++
	//	fmt.Println(b)
	//}
	//
	//var c = total(5, 6)
	//fmt.Println(c)

	//初始化数据库
	db.InitDB()

	//初始化http
	Http.HttpInit()
	//bugCommitReq := new(Req.BugCommitReq)
	//fmt.Println(Req.ToString(*bugCommitReq))
}

/**
两个int相加
*/
func total(a int, b int) int {
	return a + b
}
