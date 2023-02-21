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

type User struct {
	gorm.Model
	PassWord      string `json:"pass_word"`
	UserName      string `json:"user_name"`
	FollowCount   int64  `json:"follow_account"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (n *User) TableName() string {
	return constants.UserTableName
}

// QueryUserInfo query user info
func QueryUserInfo(ctx context.Context, userIds []int64) ([]*User, error) {
	var res []*User
	if err := DB.WithContext(ctx).Where("ID = ", userIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
