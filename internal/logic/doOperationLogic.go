package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mmChat/internal/svc"
)

type DoOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoOperationLogic {
	return &DoOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoOperationLogic) DoOperation() error {
	// todo: add your logic here and delete this line

	return nil
}
