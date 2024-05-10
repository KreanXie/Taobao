package model

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

type ConsigneeAddInput struct {
	CommonConsignee
	UserId uint `json:"user_id"`
}

type ConsigneeAddOutput struct {
	Id uint `json:"id"`
}

type ConsigneeUpdateInput struct {
	CommonConsignee
	Id uint `json:"id"`
}

type ConsigneeUpdateOutput struct {
	Id uint `json:"id"`
}

type ConsigneeDeleteInput struct {
	Id uint `json:"id"`
}

type ConsigneeDeleteOutput struct {
	Id uint `json:"id"`
}

type ConsigneeListInput struct {
	UserId uint `json:"user_id" description:"用户ID"`
	Page   int  // 分页号码
	Size   int  // 分页数量，最大50
}

type ConsigneeListOutput struct {
	List  []CommonConsignee `json:"list" description:"列表"`
	Page  int               `json:"page" description:"分页码"`
	Size  int               `json:"size" description:"分页数量"`
	Total int               `json:"total" description:"数据总数"`
}
