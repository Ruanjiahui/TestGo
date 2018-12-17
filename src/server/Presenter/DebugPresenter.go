package Presenter

import (
	"fmt"
	"github.com/json-iterator/go"
	"server/Status"
	"server/bean/Req"
	"server/bean/Resp"
	"server/db"
)

/**
处理添加错误日志的业务
*/
func BugCommit(bugCommitReq *Req.BugCommitReq) string {
	//判断参数是否为空
	//if bugCommitReq.BugData == "" {
	//	return Status.Illegal_parameter_fun()
	//}
	//将数据插入数据库
	result, _, err := db.InsertBugCommit(bugCommitReq)

	fmt.Println("result : ", result, " err : ", err)

	//返回错误则返回服务器错误
	if err != nil {
		return Status.Server_error_fun()
	}
	//result 返回false 则提示插入失败
	if !result {
		return Status.Fail_fun()
	}
	//返回成功
	return Status.Success_fun()
}

/**
处理添加错误日志的业务
*/
func BugPageGet(bugCommitReq *Req.BugPageGetReq) string {
	//判断参数是否为空
	if bugCommitReq.Page <= 0 {
		return Status.Illegal_parameter_fun()
	}
	bugCommitReq.StartFlag = (bugCommitReq.Page - 1) * bugCommitReq.Pagecount
	//查询数据
	bugBeanList, err := db.SelectBugPageGet(bugCommitReq)

	//返回错误则返回服务器错误
	if err != nil {
		return Status.Server_error_fun()
	}

	//将数据插入数据库
	count, err := db.BugGetTotal()

	bugGetPageResp := new(Resp.BugGetPageResp)
	bugGetPageResp.Data = bugBeanList
	bugGetPageResp.Msg = Status.Success_msg
	bugGetPageResp.Code = Status.Success
	bugGetPageResp.Count = count
	bugGetPageResp.Page = bugCommitReq.Page
	bugGetPageResp.Pagecount = bugCommitReq.Pagecount
	bugGetPageResp.Pagetotal = count / bugCommitReq.Pagecount
	if count % bugCommitReq.Pagecount != 0 {
		bugGetPageResp.Pagetotal = count / bugCommitReq.Pagecount + 1
	}

	msg, err := jsoniter.Marshal(*bugGetPageResp)
	//返回错误则返回服务器错误
	if err != nil {
		return Status.Server_error_fun()
	}
	//返回成功
	return string(msg)
}

/**
获取指定debug的详细信息
*/
func BugDetail(bugDetailReq *Req.BugDetailReq) string {
	//判断参数是否为空
	if bugDetailReq.DebugID <= 0 {
		return Status.Illegal_parameter_fun()
	}
	//将数据插入数据库
	bugBean, err := db.SelectBugDetail(bugDetailReq)

	//返回错误则返回服务器错误
	if err != nil {
		return Status.Server_error_fun()
	}

	bugDetailResp := new(Resp.BugDetailResp)
	bugDetailResp.Msg = Status.Success_msg
	bugDetailResp.Code = Status.Success
	bugDetailResp.Data = bugBean

	msg, err := jsoniter.Marshal(*bugDetailResp)
	//返回错误则返回服务器错误
	if err != nil {
		return Status.Server_error_fun()
	}
	//返回成功
	return string(msg)
}
