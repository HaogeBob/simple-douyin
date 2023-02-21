// Copyright 2022 Tik-Tok Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"

	"github.com/simple/douyin/cmd/favorite/pack"
	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/favorite"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList create favorite info
// todo
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	favoriteModels, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videoIds := pack.VideoIds(favoriteModels)
	videos, err := db.QueryVideoInfo(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}
	userIds := pack.UserIds(videos)
	users, err := db.QueryUserInfo(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	var res []*favorite.Video
	for i := 0; i < len(videos); i++ {
		res = append(res, &favorite.Video{
			Id: int64(videos[i].ID),
			Author: &favorite.User{
				Id:            int64(users[i].ID),
				Name:          users[i].UserName,
				FollowCount:   users[i].FollowCount,
				FollowerCount: users[i].FollowerCount,
				IsFollow:      users[i].IsFollow,
			},
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavorite,
			Title:         videos[i].VideoTitle,
		})
	}

	return res, nil
}
