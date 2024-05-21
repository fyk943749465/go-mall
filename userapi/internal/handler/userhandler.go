package handler

import (
	"userapi/internal/svc"
)

type UserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUserHandler(svcCtx *svc.ServiceContext) *UserHandler {

	return &UserHandler{
		svcCtx: svcCtx,
	}

}
