package Controller

import (
	"fmt"
	"io"
	"net/http"
	"server/Presenter"
	"server/bean/Req"
	"strconv"
)

/**
添加bug
*/
func BugCommit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")

	if request.Method == "POST" {

		var err error
		//解析接收的数据
		bugCommitReq := new(Req.BugCommitReq)
		bugCommitReq.Bround = request.PostFormValue("bround")
		bugCommitReq.Model = request.PostFormValue("model")
		bugCommitReq.BugData = request.PostFormValue("bugData")
		bugCommitReq.DebugOS = request.PostFormValue("debugOS")
		bugCommitReq.DebugOSVersion = request.PostFormValue("debugOSVersion")
		bugCommitReq.DebugLon, err = strconv.ParseFloat(request.PostFormValue("debugLon"), 64)
		bugCommitReq.DebugLat, err = strconv.ParseFloat(request.PostFormValue("debugLat"), 64)
		bugCommitReq.AppPackage = request.PostFormValue("appPackage")
		bugCommitReq.AppVersionName = request.PostFormValue("appVersionName")
		bugCommitReq.AppVersionCode, err = strconv.Atoi(request.PostFormValue("appVersionCode"))
		bugCommitReq.AppInstallDate = request.PostFormValue("appInstallDate")
		bugCommitReq.AppInstallUpdateDate = request.PostFormValue("appInstallUpdateDate")
		bugCommitReq.PhoneType = request.PostFormValue("phoneType")

		//int 格式转换错误
		if err != nil {
			io.WriteString(response, "400 Bad Request")
			return
		}

		msg := Presenter.BugCommit(bugCommitReq)
		//返回数据
		io.WriteString(response, msg)
	} else if request.Method == "OPTIONS" {
		io.WriteString(response, "")
	} else {
		io.WriteString(response, "405")
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
