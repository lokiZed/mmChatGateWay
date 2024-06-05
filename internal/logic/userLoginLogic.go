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

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (*types.UserLoginRes, error) {
	rpcUser := user.NewUser(rpcClient.GetRpcClient())
	rpcReq := &mmChatUserRpc.UserLoginReq{
		UserName: req.UserName,
		UserPass: req.UserPass,
	}
	rpcRes, err := rpcUser.UserLogin(l.ctx, rpcReq)
	if err != nil {
		l.Errorw("userLoginRpc", logx.Field("err", err.Error()))
		return nil, errors.New(define.ResponseCodePanic, "登陆失败")
	}
	if rpcRes.Msg != "" {
		l.Infow("userLoginRpc", logx.Field("msg", rpcRes.Msg))
		return nil, errors.New(define.ResponseCodeBusError, rpcRes.Msg)
	}
	out := &types.UserLoginRes{
		Token: rpcRes.Data.Token,
	}
	return out, nil
}
