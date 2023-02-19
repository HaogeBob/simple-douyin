package db

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type Video struct {
	gorm.Model
	AuthorId      int32  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavorateCount int64  `json:"favorate_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}
	return nil
}

func MGetVideos(ctx context.Context, userID int64) ([]*Video, error) {
	var res []*Video

	if err := DB.WithContext(ctx).Where("author_id = ?", userID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
