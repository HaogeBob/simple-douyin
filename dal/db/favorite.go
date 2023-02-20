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
	"fmt"

	"github.com/simple/douyin/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	FavoriteId      int64 `json:"favorite_id"`
	UserId          int64 `json:"user_id"`
	FavoriteVideoId int64 `json:"favorite_video_id"`
}

func (n *Favorite) TableName() string {
	return constants.FavoriteTableName
}

// FavoriteAction update favorite info
func FavoriteAction(ctx context.Context, favorites []*Favorite) error {
	fmt.Println("db here")
	if err := DB.WithContext(ctx).Create(favorites).Error; err != nil {
		return err
	}
	return nil
}

func FavoriteList(ctx context.Context, userID int64) ([]*Favorite, error) {
	var res []*Favorite
	if err := DB.WithContext(ctx).Where("id in ?", userID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
