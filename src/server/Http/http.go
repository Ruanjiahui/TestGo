package Http

import (
	"net/http"
	"server/Controller"
	"time"
	"fmt"
)

func HttpInit() {
	s := http.Server{
		//监听端口
		Addr: ":8890",
		//读取超时
		ReadTimeout: 10 * time.Second,
		//写入超时
		WriteTimeout: 10 * time.Second,
	}
	//创建路由
	routerList := http.NewServeMux()
	//监听路由
	routerList.HandleFunc("/BugCommit", Controller.BugCommit)
	routerList.HandleFunc("/BugPageGet", Controller.BugPageGet)
	routerList.HandleFunc("/BugGet", Controller.BugDetail)
	//配置路由
	s.Handler = routerList

	//s.ListenAndServe()

	//监听服务
	err := s.ListenAndServeTLS("/opt/ssl/214793394590511.pem", "/opt/ssl/214793394590511.key")
	if err != nil {
		fmt.Println(err)
	}
}
