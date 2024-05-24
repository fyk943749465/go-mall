package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/score"
	"rpc-common/user"
	"strconv"
	"time"
	"userapi/internal/errorx"

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

	userId, _ := strconv.ParseInt(userResponse.Id, 10, 64)
	userScore, err := l.svcCtx.UserScoreRpc.SaveScore(context.Background(), &score.UserScoreRequest{
		UserId: userId,
		Score:  10,
	})
	fmt.Sprintf("register add score %d \n", userScore.Score)
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserLogic) GetUser(t *types.IdRequest) (resp *types.Response, err error) {

	userId := l.ctx.Value("userId")
	logx.Infof("get token : %s \n", userId)

	if t.Id == "1" {
		return nil, errorx.ParamsError
	}

	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: t.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{
		Message: "success",
		Data:    userResponse,
	}
	return
}

func (l *UserLogic) getToken(secretKey string, iat, seconds int64, userId int) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserLogic) Login(t *types.LoginRequest) (string, error) {

	logx.Info("login executing....")
	userId := 100
	auth := l.svcCtx.Config.Auth
	return l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, userId)
}
