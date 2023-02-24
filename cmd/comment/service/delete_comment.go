package service

import (
	"context"
	"errors"
	// "sync"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/comment"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type DeleteCommentService struct {
	ctx context.Context
}

// 创建一个NewDeleteCommentService对象并返回
func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

// Delete Comment
func (s *DeleteCommentService) DeleteComment(req *comment.DeleteCommentRequest) (*comment.Comment, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}
	if currentId <= 0 {
		return nil, errors.New("Token compare error")
	}
	
	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("videoId not exist")
	}
	comments, err := db.QueryCommentByCommentIds(s.ctx, []int64{req.CommentId})
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return nil, errors.New("commentId not exist")
	}

	var commentRaw *db.CommentRaw
	//删除评论记录并减少视频评论数
	commentRaw, err = db.DeleteComment(s.ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	//获取用户信息
	users, err := db.QueryUserByIds(s.ctx, []int64{currentId})
	if err != nil {
		return nil, err
	}
	comment := pack.CommentInfo(commentRaw, users[0])
	return comment, nil
}
