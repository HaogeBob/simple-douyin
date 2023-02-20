package db

import (
	"context"

	"github.com/simple/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Follow_count   int64  `json:"follow_count"`
	Follower_count int64  `json:"follower_count"`
	Is_follow      bool   `json:"is_follow"`
}

type Videos struct {
	gorm.Model
	Id             int64  `json:"id"`
	Author_id      int64  `json:"author_id"`
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	Favorite_count int64  `json:"favorite_count"`
	Comment_count  int64  `json:"comment_count"`
	Is_favorite    bool   `json:"is_favorite"`
	Title          string `json:"title"`
	Token          string `json:"token"`
	Release_time   int64  `json:"release_time"`
}

func (v *Videos) TableName() string {
	return constants.VideoTableName
}

func (u *User) TableName() string {
	return constants.UserinfoName
}

func Query_user(ctx context.Context, token string) (User, error) {

	var res User
	if err := DB.WithContext(ctx).Where("token == ?", token).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil

}

func Query_vedio(ctx context.Context, author_id int64, time int64) ([]*Videos, error) {

	res := make([]*Videos, 0)
	if err := DB.WithContext(ctx).Where("author_id = ? and release_time >= ?", author_id, time).Order("release_time").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil

}
