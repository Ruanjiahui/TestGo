package Req

/**
提交代码的结构体
*/
type BugCommitReq struct {

	//手机品牌
	Bround string

	//手机型号
	Model string

	//错误日志
	BugData string

	//手机系统
	DebugOS string

	//手机系统版本
	DebugOSVersion string

	//经度
	DebugLon float64

	//纬度
	DebugLat float64

	//app 包名
	AppPackage string

	//app版本名称
	AppVersionName string

	//app版本号
	AppVersionCode int

	//安装日期
	AppInstallDate string

	//app 更新日期
	AppInstallUpdateDate string

	//安装类型
	PhoneType string

	//使用者
	User string

	Ip string

	Port string
}
