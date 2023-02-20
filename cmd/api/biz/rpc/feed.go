package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/simple/douyin/kitex_gen/feed"
	"github.com/simple/douyin/kitex_gen/feed/feedservice"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/errno"
)

var feedClient feedservice.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := feedservice.NewClient(
		constants.FeedServiceName,
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
	feedClient = c

}

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp, nil
}
