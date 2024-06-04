package rpcClient

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"mmChat/internal/config"
	"sync"
)

var globalClient zrpc.Client

var once sync.Once

func InitRpcClient(c config.Config) {
	once.Do(func() {
		client, err := zrpc.NewClient(c.RpcClientConf)
		if err != nil {
			log.Fatal(errorx.Wrap(err, "init rpc client faiul").Error())
		}
		globalClient = client
	})
}

func GetRpcClient() zrpc.Client {
	return globalClient
}
