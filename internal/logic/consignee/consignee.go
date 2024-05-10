package consignee

import (
	"context"
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
)

const tableName = "consignee_info"

type sConsignee struct{}

func init() {
	service.RegisterConsignee(New())
}

func New() *sConsignee {
	return &sConsignee{}
}

func (s *sConsignee) Create(ctx context.Context, in model.ConsigneeAddInput) (out model.ConsigneeAddOutput, err error) {
	lastInsertID, err := g.Model(tableName).Data(g.Map{
		"user_id":    in.UserId,
		"is_default": in.IsDefault,
		"name":       in.Name,
		"city":       in.City,
		"detail":     in.Detail,
		"phone":      in.Phone,
		"province":   in.Province,
		"town":       in.Town,
		"street":     in.Street,
	}).InsertAndGetId()
	if err != nil {
		return out, errors.New("consignee create err:" + err.Error())
	}
	return model.ConsigneeAddOutput{Id: uint(lastInsertID)}, nil
}

func (s *sConsignee) Update(ctx context.Context, in model.ConsigneeUpdateInput) (out model.ConsigneeUpdateOutput, err error) {
	result, err := g.Model(tableName).Data(g.Map{
		"is_default": in.IsDefault,
		"name":       in.Name,
		"city":       in.City,
		"detail":     in.Detail,
		"phone":      in.Phone,
		"province":   in.Province,
		"town":       in.Town,
		"street":     in.Street,
	}).Where("id", string(in.Id)).Update()
	if err != nil {
		return out, errors.New("consignee update err:" + err.Error())
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return out, errors.New("consignee update err:" + err.Error())
	}
	return model.ConsigneeUpdateOutput{Id: uint(lastInsertID)}, nil
}

func (s *sConsignee) Delete(ctx context.Context, in model.ConsigneeDeleteInput) (out model.ConsigneeDeleteOutput, err error) {
	result, err := g.Model(tableName).Where("id", in.Id).Unscoped().Delete()
	if err != nil {
		return out, errors.New("consignee delete err at logic:" + err.Error())
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return out, errors.New("consignee delete err at logic:" + err.Error())
	}
	return model.ConsigneeDeleteOutput{Id: uint(lastInsertID)}, nil
}

func (s *sConsignee) List(ctx context.Context, in model.ConsigneeListInput) (out model.ConsigneeListOutput, err error) {
	err = g.Model(tableName).Where("user_id", in.UserId).Scan(&out.List)
	if err != nil {
		return out, errors.New("consignee list err:" + err.Error())
	}
	return out, nil
}
