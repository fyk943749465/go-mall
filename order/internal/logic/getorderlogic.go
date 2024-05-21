package logic

import (
	"context"
	"rpc-common/user"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	// get userid by orderId
	userId := l.getOrderById(req.Id)
	// get user by userid
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return &types.OrderReply{
		Id:       req.Id,
		Name:     "hello order name",
		UserName: userResponse.GetName(),
	}, nil
}

func (l *GetOrderLogic) getOrderById(id string) string {
	return "1"
}
