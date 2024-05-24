package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rpc-common/score"
	"user-score/internal/model"

	"user-score/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"

	_ "github.com/dtm-labs/dtmdriver-gozero"
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

	barrier, err2 := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err2 != nil {
		// internal error 代表重试
		return nil, status.Error(codes.Internal, err2.Error())
	}
	userScore := &model.UserScore{
		UserId: in.UserId,
		Score:  int(in.Score),
	}

	err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {

		if userScore.Score == 10 {
			return errors.New("模拟积分服务出现问题")
		}
		err := l.svcCtx.UserScoreRepo.SaveUserScore(tx, context.Background(), userScore)
		if err != nil {
			return err
		}
		return nil
	})
	// 这里返回错误，代表回滚 Aborted
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &score.UserScoreResponse{
		UserId: userScore.UserId,
		Score:  int32(userScore.Score),
	}, nil
}
