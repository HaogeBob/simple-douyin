package main

import (
	"context"

	"github.com/simple/douyin/cmd/publish/service"
	"github.com/simple/douyin/dal/pack"
	publish "github.com/simple/douyin/kitex_gen/publish"
	"github.com/simple/douyin/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(publish.PublishActionResponse)

	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(publish.PublishListResponse)

	videos, err := service.NewMGetVideoService(ctx).MGetVideo(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
