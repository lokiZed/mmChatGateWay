package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"mmChat/internal/define"
	"mmChat/internal/rpcClient"
	"mmChat/internal/rpcClient/pb/mmChatUserRpc"
	"mmChat/internal/rpcClient/user"

	"mmChat/internal/svc"
	"mmChat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) error {
	registerIp := l.ctx.Value(define.CtxKeyIp)
	rpcUser := user.NewUser(rpcClient.GetRpcClient())
	rpcReq := &mmChatUserRpc.UserRegisterReq{
		Avatar:   req.Avatar,
		UserName: req.UserName,
		UserPass: req.UserPass,
		NickName: req.NickName,
		Gender:   mmChatUserRpc.Gender(req.Gender),
		Age:      req.Age,
		DeviceId: req.DeviceId,
		Ip:       registerIp.(string),
	}
	rpcRes, err := rpcUser.UserRegister(l.ctx, rpcReq)
	if err != nil {
		l.Errorw("userRegisterRpc", logx.Field("err", err.Error()))
		return errors.New(define.ResponseCodePanic, "注册失败")
	}
	if rpcRes.Msg != "" {
		l.Errorw("userRegisterRpc", logx.Field("msg", rpcRes.Msg))
		return errors.New(define.ResponseCodeBusError, rpcRes.Msg)
	}
	return errors.New(define.ResponseCodeOk, "")
}
