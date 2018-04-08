package main

import (
	"config"
	"fmt"
)

const PI  = 3.14

var PO  = "goland"

type PK string

//声明结构
type gopher struct {}

//声明接口
type golang interface {}

//定义一个bool类型
type b bool

func main() {
	config.Loadconfig()

	fmt.Println("Hello word!")


	var a int
	var b = 2

	//var bl bool

	for a = 2; a < 3; a++ {
		b++
		fmt.Println(b)
	}

	var c = total(5 , 6)
	fmt.Println(c)
}

/**
	两个int相加
 */
func total(a int , b int) int {
	return a + b
}


