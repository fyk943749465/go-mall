package logic

import (
	"context"
	"fmt"
	"rpc-common/score"

	"user-score/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveScoreCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveScoreCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveScoreCallbackLogic {
	return &SaveScoreCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveScoreCallbackLogic) SaveScoreCallback(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	// todo: add your logic here and delete this line

	fmt.Println("-------------------------------------------")
	fmt.Println("user-score saveUserScoreCallback ..........")
	fmt.Println("-------------------------------------------")
	return &score.UserScoreResponse{}, nil
}
