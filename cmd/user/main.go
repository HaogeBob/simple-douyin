package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/simple/douyin/dal"
	user "github.com/simple/douyin/kitex_gen/user/userservice"
	"github.com/simple/douyin/pkg/constants"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.UserAddress)
	if err != nil {
		panic(err)
	}

	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		// server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		// server.WithMuxTransport(),
		// server.WithMiddleware(mw.CommonMiddleware),
		// server.WithMiddleware(mw.ServerMiddleware),
		// server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
