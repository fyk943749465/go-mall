package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/score"
	"rpc-common/user"
	"time"
	"userapi/internal/errorx"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	_ "github.com/dtm-labs/dtmdriver-gozero"
)

var dtmServer = "etcd://localhost:2379/dtmservice"

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

	//ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelFunc()
	//userResponse, err := l.svcCtx.UserRpc.Save(ctx, &user.UserRequest{
	//	Name:   req.Name,
	//	Gender: req.Gender,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//userId, _ := strconv.ParseInt(userResponse.Id, 10, 64)
	//userScore, err := l.svcCtx.UserScoreRpc.SaveScore(context.Background(), &score.UserScoreRequest{
	//	UserId: userId,
	//	Score:  10,
	//})

	gid := dtmgrpc.MustGenGid(dtmServer)
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	userServer, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, err
	}

	userScoreServer, err := l.svcCtx.Config.UserScoreRpc.BuildTarget()
	if err != nil {
		return nil, err
	}

	userReq := &user.UserRequest{
		Id:     req.Id,
		Name:   req.Name,
		Gender: req.Gender,
	}
	// call save method
	sagaGrpc.Add(userServer+"/user.User/save", userServer+"/user.User/saveCallback", userReq)
	// 这个地方，应该是传入一个User，因为远程调用拿不到返回值。暂且先写死，为了测试效果。
	userScoreReq := &score.UserScoreRequest{
		UserId: req.Id,
		Score:  10,
	}

	// 这里出现问题，不能调用saveScoreCallback。这里调用，是逻辑补偿。显然，user 微服务的方法就不会被回滚了。
	// sagaGrpc.Add(userScoreServer+"/userscore.UserScore/saveScore", userScoreServer+"/userscore.UserScore/saveScoreCallback", userScoreReq)
	sagaGrpc.Add(userScoreServer+"/userscore.UserScore/saveScore", "", userScoreReq)
	sagaGrpc.WaitResult = true
	err = sagaGrpc.Submit()
	if err != nil {
		fmt.Println("---------------------------")
		fmt.Println(err)
		return nil, err
	}
	//fmt.Sprintf("register add score %d \n", userScore.Score)
	return &types.Response{
		Message: "success",
		Data:    "",
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
