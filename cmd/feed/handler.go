package main

import (
	"context"

	"github.com/simple/douyin/cmd/feed/service"
	"github.com/simple/douyin/kitex_gen/feed"
	"github.com/simple/douyin/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {

	videos, time, err := service.NewFeedService(ctx).Feed(req.Token, req.LatestTime)

	var res feed.BaseResp

	next_time := time
	if err != nil {

		res.StatusCode = errno.ServiceErr.ErrCode
		res.StatusMessage = errno.ServiceErr.ErrMsg
		res.ServiceTime = next_time

		return &feed.FeedResponse{
			BaseResp:  &res,
			VideoList: nil,
			NextTime:  next_time,
		}, nil
	}
	//视频加载没有出错
	res.StatusCode = errno.Success.ErrCode
	res.StatusMessage = errno.Success.ErrMsg
	res.ServiceTime = next_time

	return &feed.FeedResponse{
		BaseResp:  &res,
		VideoList: videos,
		NextTime:  next_time,
	}, nil

}
