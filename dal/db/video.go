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

package db

import (
	"context"

	"github.com/simple/douyin/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	// VideoId       int64  `json:"video_id"`
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"player_id"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	VideoTitle    string `json:"video_title"`
	CreateTime    string `json:"create_time"`
}

func (n *Video) TableName() string {
	return constants.VideoTableName
}

// QueryVideoInfo query video info
func QueryVideoInfo(ctx context.Context, videoIds []int64) ([]*Video, error) {
	var res []*Video
	if err := DB.WithContext(ctx).Where("ID = ", videoIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
