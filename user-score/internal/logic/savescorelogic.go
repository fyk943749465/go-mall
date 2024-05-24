package logic

import (
	"context"
	"rpc-common/score"
	"user-score/internal/model"

	"user-score/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveScoreLogic {
	return &SaveScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveScoreLogic) SaveScore(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	userScore := &model.UserScore{
		UserId: in.UserId,
		Score:  int(in.Score),
	}
	err := l.svcCtx.UserScoreRepo.SaveUserScore(context.Background(), userScore)
	if err != nil {
		return nil, err
	}

	return &score.UserScoreResponse{
		UserId: userScore.UserId,
		Score:  int32(userScore.Score),
	}, nil
}
