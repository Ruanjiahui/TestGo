package config

import "fmt"
//首字母为大写则是public
func Loadconfig() {
	fmt.Println("welcome load config")
}

//首字母为小写则是private
func loadconfig()  {
	fmt.Println("private load config")
}
