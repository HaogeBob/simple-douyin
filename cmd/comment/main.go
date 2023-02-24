package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	tracer1 "github.com/kitex-contrib/tracer-opentracing"
	"github.com/simple/douyin/dal"
	comment "github.com/simple/douyin/kitex_gen/comment/commentservice"
	"github.com/simple/douyin/pkg/bound"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/middleware"
	tracer2 "github.com/simple/douyin/pkg/tracer"
)

func Init() {
	tracer2.InitJaeger(constants.CommentServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), 
		server.WithMuxTransport(),                                          
		server.WithSuite(tracer1.NewDefaultServerSuite()),                  
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                
		server.WithRegistry(r),                                             
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
