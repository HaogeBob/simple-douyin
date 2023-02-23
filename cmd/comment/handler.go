package main

import (
	"context"

	"github.com/simple/douyin/cmd/comment/service"
	"github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/comment"
	"github.com/simple/douyin/pkg/errno"
)

// implement IDL service interface of comment
type CommentServiceImpl struct{}

// Create Comment
func (*CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	resp = new(comment.CreateCommentResponse)
	// 判断 token和videoid 有效
	if len(req.Token) == 0 || req.VideoId == 0 {
		resp.BaseResp = pack.NewCommentBaseResp(errno.ParamErr)
		resp.Comment = nil
		return resp, nil
	}

	comment, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.NewCommentBaseResp(err)
		resp.Comment = nil
		return resp, nil
	}

	resp.BaseResp = pack.NewCommentBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// Delete Comment
func (*CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (resp *comment.DeleteCommentResponse, err error) {
	resp = new(comment.DeleteCommentResponse)
	// 判断 token、videoid和commentid 有效
	if len(req.Token) == 0 || req.VideoId == 0 || req.CommentId == 0 {
		resp.BaseResp = pack.NewCommentBaseResp(errno.ParamErr)
		resp.Comment = nil
		return resp, nil
	}

	comment, err := service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.NewCommentBaseResp(err)
		resp.Comment = nil
		return resp, nil
	}

	resp.BaseResp = pack.NewCommentBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// Return Comment List
func (*CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = new(comment.CommentListResponse)

	if len(req.Token) == 0 || req.VideoId == 0 {
		resp.BaseResp = pack.NewCommentBaseResp(errno.ParamErr)
		resp.CommentList = nil
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.NewCommentBaseResp(err)
		resp.CommentList = nil
		return resp, nil
	}

	resp.BaseResp = pack.NewCommentBaseResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}
