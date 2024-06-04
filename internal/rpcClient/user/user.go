package user

import (
	"context"

	"mmChat/internal/rpcClient/pb/mmChatUserRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserLoginReq     = mmChatUserRpc.UserLoginReq
	UserLoginRes     = mmChatUserRpc.UserLoginRes
	UserLoginResData = mmChatUserRpc.UserLoginResData
	UserRegisterReq  = mmChatUserRpc.UserRegisterReq
	UserRegisterRes  = mmChatUserRpc.UserRegisterRes

	User interface {
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRes, error)
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRes, error) {
	client := mmChatUserRpc.NewUserClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}

func (m *defaultUser) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginRes, error) {
	client := mmChatUserRpc.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}
