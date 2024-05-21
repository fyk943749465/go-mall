package logic

import (
	"context"
	"rpc-common/user"
	"time"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	userResponse, err := l.svcCtx.UserRpc.Save(ctx, &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}
