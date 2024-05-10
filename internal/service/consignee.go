// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IConsignee interface {
		Create(ctx context.Context, in model.ConsigneeAddInput) (out model.ConsigneeAddOutput, err error)
		Update(ctx context.Context, in model.ConsigneeUpdateInput) (out model.ConsigneeUpdateOutput, err error)
		Delete(ctx context.Context, in model.ConsigneeDeleteInput) (out model.ConsigneeDeleteOutput, err error)
		List(ctx context.Context, in model.ConsigneeListInput) (out model.ConsigneeListOutput, err error)
	}
)

var (
	localConsignee IConsignee
)

func Consignee() IConsignee {
	if localConsignee == nil {
		panic("implement not found for interface IConsignee, forgot register?")
	}
	return localConsignee
}

func RegisterConsignee(i IConsignee) {
	localConsignee = i
}
