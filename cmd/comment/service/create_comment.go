package service

import (
	"context"
	"errors"
	"sync"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/comment"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type CreateCommentService struct {
	ctx context.Context
}

// 创建一个CreateCommentService对象并返回
func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

// Create Comment
func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("video not exist")
	}

	commentRaw := &db.CommentRaw{
		UserId:   currentId,
		VideoId:  req.VideoId,
		Contents: req.CommentText,
	}
	//创建评论记录并增加视频评论数
	err := db.CreateComment(s.ctx, commentRaw)
	if err != nil {
		return nil, err
	}
	//获取当前用户信息
	users, err := db.QueryUserByIds(s.ctx, []int64{currentId})
	if err != nil {
		return nil, err
	}
	comment := pack.CommentInfo(commentRaw, users[0])
	return comment, nil
}
