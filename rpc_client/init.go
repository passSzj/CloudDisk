package shorturlservice

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-cloud-disk/conf"
)

var ShortUrlClient ShortUrlService

func InitGrpcClient() {

	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{conf.RpcHost}, // 你的 gRPC 服务地址
		NonBlock:  true,
	})

	ShortUrlClient = NewShortUrlService(cli)

}
