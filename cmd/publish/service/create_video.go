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
		AuthorId:      111,
		PlayUrl:       "111",
		CoverUrl:      "111",
		Title:         req.Title,
		FavorateCount: 111,
		CommentCount:  111,
		IsFavorite:    true,
	}
	return db.CreateVideo(s.ctx, []*db.Video{videoModel})
}
