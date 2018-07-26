package db

import (
	"fmt"
	"server/bean/Req"
	"server/bean/Resp"
)

/**
插入错误日志
*/
func InsertBugCommit(req *Req.BugCommitReq) (bool, int64, error) {
	//var tx *sql.Tx

	//开启事务
	//tx , err = DB.Begin()
	//if err != nil{
	//	fmt.Println("tx fail : " , err)
	//	return false , 0 , err
	//}

	//准备sql语句
	stmt, err := DB.Prepare("INSERT INTO debug.debugdata (debugSource, debugBround, debugModel, debugTime , debugOS , debugOSVersion , debugLon , debugLat , appPackage , appVersionCode , appVersionName , appInstallDate , appInstallUpdateDate , phoneType) VALUES (? , ? , ? , now() , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?)")
	if err != nil {
		fmt.Println("Prepare fail : ", err)
		return false, 0, err
	}

	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(req.BugData, req.Bround, req.Model, req.DebugOS, req.DebugOSVersion, req.DebugLon, req.DebugLat, req.AppPackage, req.AppVersionCode, req.AppVersionName, req.AppInstallDate, req.AppInstallUpdateDate, req.PhoneType)
	if err != nil {
		fmt.Println("Exec fail : ", err)
		return false, 0, err
	}
	//插入数据返回id
	debugID, err := res.LastInsertId()
	//事务提交
	//err = tx.Commit()
	//if err != nil{
	//	//事务回滚
	//	tx.Rollback()
	//	return false , debugID , err
	//}
	return true, debugID, err
}

/**
分页获取debug的数据
*/
func SelectBugPageGet(req *Req.BugPageGetReq) ([]Resp.BugBean, error) {
	stmt, err := DB.Prepare("SELECT debugID , debugBround , debugModel , debugTime , debugOS , debugOSVersion , appVersionName , appVersionCode from debug.debugdata limit ? offset ?")
	if err != nil {
		fmt.Println("Prepare fail : ", err)
		return nil, err
	}

	rows, err := stmt.Query(req.Pagecount, req.StartFlag)
	if err != nil {
		fmt.Println("stmt fail : ", err)
		return nil, err
	}

	//创建切片数组的长度
	var bugBeanList = make([]Resp.BugBean, req.Pagecount)

	num := 0
	//循环读取数据库的数据
	for rows.Next() {
		bugBean := new(Resp.BugBean)
		err = rows.Scan(&bugBean.DebugID, &bugBean.DebugBround, &bugBean.DebugModel, &bugBean.DebugTime, &bugBean.DebugOS, &bugBean.DebugOSVersion, &bugBean.AppVersionName, &bugBean.AppVersionCode)
		if err != nil {
			fmt.Println("Scan fail : ", err)
			return nil, err
		}
		bugBeanList[num] = *bugBean
		num++
	}

	var cpbugBeanList = make([]Resp.BugBean, num)
	var i int
	for i = 0; i < num; i++ {
		cpbugBeanList[i] = bugBeanList[i]
	}
	return cpbugBeanList, err
}

/**
查询debug数据的行数
*/
func BugGetTotal() (int, error) {

	stmt, err := DB.Prepare("SELECT count(*) from debug.debugdata")
	if err != nil {
		fmt.Println("bugGetTotal stmt fail : ", err)
		return 0, err
	}

	rows := stmt.QueryRow()

	var length int
	err = rows.Scan(&length)
	if err != nil {
		fmt.Println("bugGetTotal QueryRow fail : ", err)
		return 0, err
	}
	return length, err
}

/**
分页获取debug的数据
*/
func SelectBugDetail(req *Req.BugDetailReq) (Resp.BugBean, error) {
	stmt, err := DB.Prepare("SELECT * from debug.debugdata where debugID = ?")
	if err != nil {
		fmt.Println("Prepare fail : ", err)
		return Resp.BugBean{}, err
	}

	rows, err := stmt.Query(req.DebugID)
	if err != nil {
		fmt.Println("stmt fail : ", err)
		return Resp.BugBean{}, err
	}

	bugBean := new(Resp.BugBean)
	//循环读取数据库的数据
	for rows.Next() {
		err = rows.Scan(&bugBean.DebugID, &bugBean.DebugSource, &bugBean.DebugBround, &bugBean.DebugModel, &bugBean.DebugTime, &bugBean.DebugOS, &bugBean.DebugOSVersion, &bugBean.DebugLon, &bugBean.DebugLat, &bugBean.AppPackage, &bugBean.AppVersionCode, &bugBean.AppVersionName, &bugBean.AppInstallDate, &bugBean.AppInstallUpdateDate, &bugBean.PhoneType)
		if err != nil {
			fmt.Println("Scan fail : ", err)
			return Resp.BugBean{}, err
		}
	}
	return *bugBean, err
}
