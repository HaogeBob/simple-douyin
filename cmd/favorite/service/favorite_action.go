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

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/favorite"
	"github.com/simple/douyin/pkg/constants"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// todo:
func TokenToUserId(token string) (int64, error) {
	return 0xffff, nil
}

// FavoriteAction create favorite info
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	userId, err := TokenToUserId(req.Token)
	if err != nil {
		return err
	}

	videoId := req.VideoId
	favoriteId, err := db.QueryFavoriteAction(s.ctx, userId, videoId)

	favoriteModel := &db.Favorite{
		FavoriteId:      favoriteId,
		UserId:          userId,
		FavoriteVideoId: videoId,
	}

	// there has been record before
	if err == nil {
		if req.ActionType == constants.Like {
			return nil
		} else {
			return db.DeleteFavoriteAction(s.ctx, favoriteId)
		}
	}

	// no record before
	if req.ActionType == constants.Unlike {
		return nil
	} else {
		return db.CreateFavoriteAction(s.ctx, []*db.Favorite{favoriteModel})
	}
}
