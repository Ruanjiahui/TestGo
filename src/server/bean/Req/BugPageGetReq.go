package Req

/**
获取每页的数据
*/
type BugPageGetReq struct {

	//当前页数
	Page int

	//当前页数数量
	Pagecount int

	//开始获取的标识
	StartFlag int
}
