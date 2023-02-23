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

func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

type User_info struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Follow_count   int64  `json:"follow_count"`
	Follower_count int64  `json:"follower_count"`
	Is_follow      bool   `json:"is_follow"`
}

func (u *User_info) TableName() string {
	return constants.UserTableName
}

func Query_user_info(ctx context.Context, userid int64) (User_info, error) {

	var res User_info

	if err := DB.WithContext(ctx).Where("id = ?", userid).Find(&res).Error; err != nil {

		return res, err
	}

	return res, nil

}
