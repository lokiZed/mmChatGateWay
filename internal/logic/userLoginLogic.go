package logic

import (
	"context"
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

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginRes, err error) {
	// todo: add your logic here and delete this line
	rpcUser := user.NewUser(rpcClient.GetRpcClient())
	rpcReq := &mmChatUserRpc.UserLoginReq{
		UserName: req.UserName,
		UserPass: req.UserPass,
	}
	rpcRes, err := rpcUser.UserLogin(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	resp = &types.UserLoginRes{
		Code: rpcRes.Code,
		Msg:  rpcRes.Msg,
		Data: types.ResData{
			Token: rpcRes.Data.GetToken(),
		},
	}
	return resp, nil
}
