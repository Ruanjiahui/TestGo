package Status

import "server/bean"

/**
服务器错误返回信息
*/
func Server_error_fun() string {
	var respBase = new(bean.RespBase)
	respBase.Code = Server_error
	respBase.Msg = Server_error_msg
	return bean.ToString(*respBase)
}

/**
成功返回信息
*/
func Success_fun() string {
	var respBase = new(bean.RespBase)
	respBase.Code = Success
	respBase.Msg = Success_msg
	return bean.ToString(*respBase)
}

/**
响应失败返回信息
*/
func Fail_fun() string {
	var respBase = new(bean.RespBase)
	respBase.Code = Fail
	respBase.Msg = Fail_msg
	return bean.ToString(*respBase)
}

/**
空值返回信息
*/
func Empty_fun() string {
	var respBase = new(bean.RespBase)
	respBase.Code = Empty
	respBase.Msg = Empty_msg
	return bean.ToString(*respBase)
}

/**
非法参数返回信息
*/
func Illegal_parameter_fun() string {
	var respBase = new(bean.RespBase)
	respBase.Code = Illegal_parameter
	respBase.Msg = Illegal_parameter_msg
	return bean.ToString(*respBase)
}
