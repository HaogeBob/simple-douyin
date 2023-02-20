package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/simple/douyin/kitex_gen/user"
	"github.com/simple/douyin/kitex_gen/user/userservice"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/errno"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// client.WithMiddleware(mw.CommonMiddleware),
		// client.WithInstanceMW(mw.ClientMiddleware),
		// client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c

}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (int64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *user.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	// fmt.Println(resp.UserId)
	return resp.UserId, nil
}

// MGetUser get user message
func MGetUser(ctx context.Context, req *user.MGetUserRequest) (*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.User, nil
}
