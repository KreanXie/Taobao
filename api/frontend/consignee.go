package frontend

import "github.com/gogf/gf/v2/frame/g"

type CommonConsignee struct {
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Town      string `json:"town"`
	Street    string `json:"street"`
	Detail    string `json:"detail"`
}

type ConsigneeAddReq struct {
	g.Meta `path:"/consignee/add" method:"post" summary:"增加收货地址接口"`
	CommonConsignee
}

type ConsigneeAddRes struct {
	Id uint `json:"id"`
}

type ConsigneeUpdateReq struct {
	g.Meta `path:"/consignee/update" method:"post" summary:"更新收货地址接口"`
	CommonConsignee
	Id uint `json:"id"`
}

type ConsigneeUpdateRes struct {
	Id uint `json:"id"`
}

type ConsigneeDeleteReq struct {
	g.Meta `path:"/consignee/delete" method:"post" summary:"删除收货地址接口"`
	Id     uint `json:"id"`
}

type ConsigneeDeleteRes struct {
	Id uint `json:"id"`
}

type ConsigneeListReq struct {
	g.Meta `path:"/consignee/list" method:"post" summary:"收货地址列表接口"`
	UserId uint `json:"user_id" description:"用户ID"`
	CommonPaginationReq
}

type ConsigneeListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
