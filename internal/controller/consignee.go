package controller

import (
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop-v2/api/frontend"
	"goframe-shop-v2/internal/consts"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
	"golang.org/x/net/context"
)

type cConsignee struct{}

var Consignee cConsignee

func (c *cConsignee) Create(ctx context.Context, req *frontend.ConsigneeAddReq) (res *frontend.ConsigneeAddRes, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	out, err := service.Consignee().Create(ctx, model.ConsigneeAddInput{
		UserId: userId,
		CommonConsignee: model.CommonConsignee{
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
	})
	if err != nil {
		return res, errors.New("consignee create err:" + err.Error())
	}
	return &frontend.ConsigneeAddRes{Id: out.Id}, nil
}

func (c *cConsignee) Update(ctx context.Context, req *frontend.ConsigneeUpdateReq) (res *frontend.ConsigneeUpdateRes, err error) {
	out, err := service.Consignee().Update(ctx, model.ConsigneeUpdateInput{
		CommonConsignee: model.CommonConsignee{
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
		Id: req.Id,
	})
	if err != nil {
		return res, errors.New("consignee update err:" + err.Error())
	}
	return &frontend.ConsigneeUpdateRes{Id: out.Id}, nil
}

func (c *cConsignee) Delete(ctx context.Context, req *frontend.ConsigneeDeleteReq) (res *frontend.ConsigneeDeleteRes, err error) {
	out, err := service.Consignee().Delete(ctx, model.ConsigneeDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, errors.New("consignee delete err at controller:" + err.Error())
	}
	return &frontend.ConsigneeDeleteRes{Id: out.Id}, nil
}

func (c *cConsignee) List(ctx context.Context, req *frontend.ConsigneeListReq) (res *frontend.ConsigneeListRes, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	out, err := service.Consignee().List(ctx, model.ConsigneeListInput{
		UserId: userId,
		Page:   req.Page,
		Size:   req.Size,
	})
	if err != nil {
		return res, errors.New("consignee list err:" + err.Error())
	}
	return &frontend.ConsigneeListRes{List: out.List, Page: out.Page, Size: req.Size}, nil
}
