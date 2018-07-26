package Resp

import (
	"server/bean"
)

type BugGetPageResp struct {
	bean.RespBase

	Count int

	Pagecount int

	Pagetotal int

	Page int

	Data []BugBean
}
