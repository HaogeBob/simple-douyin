package db

import (
	"context"

	"github.com/simple/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavorateCount int64  `json:"favorate_count"`
	CommentCount  int64  `json:"comment_count"`
}

type Relation struct {
	gorm.Model
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
}

type Favorate struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

type Comment struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}

func (f *Favorate) TableName() string {
	return constants.FavoriteTableName
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

func QueryFollowCount(ctx context.Context, userID int64) (int64, error) {
	var count int64
	if err := DB.Model(&Relation{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func QueryFollowerCount(ctx context.Context, userID int64) (int64, error) {
	var count int64
	if err := DB.Model(&Relation{}).Where("to_user_id = ?", userID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func IsFavorite(ctx context.Context, userID int64, videoID int64) (bool, error) {
	var relation *Relation
	if err := DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&relation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func QueryFavoriteCount(ctx context.Context, videoID int64) (int64, error) {
	var count int64
	if err := DB.Model(&Favorate{}).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func QueryCommentCount(ctx context.Context, videoID int64) (int64, error) {
	var count int64
	if err := DB.Model(&Comment{}).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
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
