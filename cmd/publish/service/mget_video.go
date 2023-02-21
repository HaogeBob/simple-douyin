package service

import (
	"context"

	"github.com/simple/douyin/cmd/publish/rpc"
	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/publish"
	"github.com/simple/douyin/kitex_gen/user"
)

type MGetVideoService struct {
	ctx context.Context
}

func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

func (s *MGetVideoService) MGetVideo(req *publish.PublishListRequest) ([]*publish.Video, error) {
	videoModels, err := db.MGetVideos(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	favorite_array := make([]bool, 0)
	for i := 0; i < len(videoModels); i++ {
		is_favorite, err := db.IsFavorite(s.ctx, req.NowUserId, int64(videoModels[i].ID))
		if err != nil {
			return nil, err
		}
		favorite_array[i] = is_favorite
	}
	user, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserId: req.UserId})
	videos := pack.Videos(videoModels, user, favorite_array)
	return videos, nil
}
