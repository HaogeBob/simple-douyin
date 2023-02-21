package service

import (
	"context"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/publish"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *publish.PublishActionRequest) error {
	videoModel := &db.Video{
		AuthorId:      req.UserId,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		Title:         req.Title,
		FavorateCount: 0,
		CommentCount:  0,
	}
	return db.CreateVideo(s.ctx, []*db.Video{videoModel})
}
