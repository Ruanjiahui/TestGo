package Controller

import (
	"fmt"
	"io"
	"net/http"
	"server/Presenter"
	"server/bean/Req"
	"strconv"
	"github.com/json-iterator/go"
	"io/ioutil"
)

/**
添加bug
*/
func BugCommit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")

	var err error
	bugCommitReq := new(Req.BugCommitReq)

	if request.Method == "POST" && request.Header.Get("Content-Type") == "application/json; charset=UTF-8" {
		var body []byte
		//解析接收的数据
		body, err = ioutil.ReadAll(request.Body)
		fmt.Println(string(body))
		err = jsoniter.Unmarshal(body, bugCommitReq)

		//int 格式转换错误
		if err != nil {
			fmt.Println(err)
			response.WriteHeader(200)
			response.Write([]byte("400 Bad Request"))
			return
		}

		msg := Presenter.BugCommit(bugCommitReq)
		//返回数据
		response.WriteHeader(200)
		response.Write([]byte(msg))
		return
	} else if request.Method == "POST" {

		bugCommitReq.Bround = request.PostFormValue("bround")
		bugCommitReq.Model = request.PostFormValue("model")
		bugCommitReq.BugData = request.PostFormValue("bugData")
		bugCommitReq.DebugOS = request.PostFormValue("debugOS")
		bugCommitReq.DebugOSVersion = request.PostFormValue("debugOSVersion")
		if request.PostFormValue("debugLon") != "" {
			bugCommitReq.DebugLon, err = strconv.ParseFloat(request.PostFormValue("debugLon"), 64)
		}
		if request.PostFormValue("debugLat") != "" {
			bugCommitReq.DebugLat, err = strconv.ParseFloat(request.PostFormValue("debugLat"), 64)
		}
		bugCommitReq.AppPackage = request.PostFormValue("appPackage")
		bugCommitReq.AppVersionName = request.PostFormValue("appVersionName")
		if request.PostFormValue("appVersionCode") != "" {
			bugCommitReq.AppVersionCode, err = strconv.Atoi(request.PostFormValue("appVersionCode"))
		}
		bugCommitReq.AppInstallDate = request.PostFormValue("appInstallDate")
		bugCommitReq.AppInstallUpdateDate = request.PostFormValue("appInstallUpdateDate")
		bugCommitReq.PhoneType = request.PostFormValue("phoneType")
		bugCommitReq.User = request.PostFormValue("user")
		bugCommitReq.Ip = request.PostFormValue("ip")
		bugCommitReq.Port = request.PostFormValue("port")

		var b []byte
		b, err = jsoniter.Marshal(bugCommitReq)
		fmt.Println(string(b))

		//int 格式转换错误
		if err != nil {
			fmt.Println(err)
			response.WriteHeader(200)
			response.Write([]byte("400 Bad Request"))
			return
		}

		msg := Presenter.BugCommit(bugCommitReq)
		//返回数据
		response.WriteHeader(200)
		response.Write([]byte(msg))
		return
	} else if request.Method == "OPTIONS" {
		response.WriteHeader(200)
		response.Write([]byte(""))
		return
	} else {
		response.WriteHeader(200)
		response.Write([]byte("405"))
		return
	}
}

/**
分页获取bug接口
*/
func BugPageGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")

	//解析接收的数据
	bugPageGetReq := new(Req.BugPageGetReq)

	var err error
	if request.Method == "GET" {
		bugPageGetReq.Page, err = strconv.Atoi(request.FormValue("page"))
		bugPageGetReq.Pagecount, err = strconv.Atoi(request.FormValue("pagecount"))
	} else if request.Method == "POST" {
		bugPageGetReq.Page, err = strconv.Atoi(request.PostFormValue("page"))
		bugPageGetReq.Pagecount, err = strconv.Atoi(request.PostFormValue("pagecount"))
	} else if request.Method == "OPTIONS" {
		io.WriteString(response, "")
	}

	//int 格式转换错误
	if err != nil {
		fmt.Println(err)
		io.WriteString(response, "400 Bad Request")
		return
	}

	msg := Presenter.BugPageGet(bugPageGetReq)
	//返回数据
	io.WriteString(response, msg)
}

/**
获取指定bug接口
*/
func BugDetail(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")

	//解析接收的数据
	bugDetailReq := new(Req.BugDetailReq)

	var err error
	if request.Method == "GET" {
		bugDetailReq.DebugID, err = strconv.Atoi(request.FormValue("debugID"))
	} else if request.Method == "POST" {
		bugDetailReq.DebugID, err = strconv.Atoi(request.PostFormValue("debugID"))
	} else if request.Method == "OPTIONS" {
		io.WriteString(response, "")
	}
	//int 格式转换错误
	if err != nil {
		fmt.Println(err)
		io.WriteString(response, "400 Bad Request")
		return
	}

	msg := Presenter.BugDetail(bugDetailReq)
	//返回数据
	io.WriteString(response, msg)
}
