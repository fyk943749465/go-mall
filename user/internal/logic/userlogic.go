package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rpc-common/user"
	"strconv"
	"user/internal/model"
	"user/internal/svc"

	_ "github.com/dtm-labs/dtmdriver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line

	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserRepo.FindById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   u.Name,
		Gender: u.Gender,
	}, nil
}

func (l *UserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {

	barrier, err2 := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err2 != nil {
		// internal error 代表重试
		return nil, status.Error(codes.Internal, err2.Error())
	}

	data := &model.User{
		Id:     in.GetId(),
		Name:   in.Name,
		Gender: in.Gender,
	}

	err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		err := l.svcCtx.UserRepo.Save(tx, context.Background(), data)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &user.UserResponse{
		Id:     strconv.FormatInt(data.Id, 10),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}

func (l *UserLogic) SaveCallback(in *user.UserRequest) (*user.UserResponse, error) {

	fmt.Println("-----------------------------------------")
	fmt.Println("save user callback .............")
	fmt.Println("------------------------------------------")
	return &user.UserResponse{}, nil
}
