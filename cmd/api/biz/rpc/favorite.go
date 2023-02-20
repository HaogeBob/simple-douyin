package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/simple/douyin/kitex_gen/favorite"
	"github.com/simple/douyin/kitex_gen/favorite/favoriteservice"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/errno"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		// client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// FavoriteAction favorite or no favorite
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	fmt.Println("rpc here, ready to invoke client")
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// FavoriteList list user favorite list
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}
