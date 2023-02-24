package main

import (
	"context"

	"github.com/simple/douyin/cmd/relation/service"
	"github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/relation"
	"github.com/simple/douyin/pkg/errno"
)

type RelationServiceImpl struct{}

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)
	// 判断 token、ToUserId和 ActionType 有效
	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType == 0 {
		resp.BaseResp = pack.NewRelationBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.NewRelationBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.NewRelationBaseResp(errno.Success)
	return resp, nil
}

func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {
	resp = new(relation.FollowListResponse)
	// 判断 token和UserId 有效
	if len(req.Token) == 0 || req.UserId == 0 {
		resp.BaseResp = pack.NewRelationBaseResp(errno.ParamErr)
		resp.UserList = nil
		return resp, nil
	}

	followList, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.NewRelationBaseResp(err)
		resp.UserList = nil
		return resp, nil
	}

	resp.BaseResp = pack.NewRelationBaseResp(errno.Success)
	resp.UserList = followList
	return resp, nil
}

func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	resp = new(relation.FollowerListResponse)
	// 判断 token和UserId 有效
	if len(req.Token) == 0 || req.UserId == 0 {
		resp.BaseResp = pack.NewRelationBaseResp(errno.ParamErr)
		resp.UserList = nil
		return resp, nil
	}

	followerList, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.NewRelationBaseResp(err)
		resp.UserList = nil
		return resp, nil
	}

	resp.BaseResp = pack.NewRelationBaseResp(errno.Success)
	resp.UserList = followerList
	return resp, nil
}
