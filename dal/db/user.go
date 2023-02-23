package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/simple/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password      string `gorm:"column:password;type:varchar(32);not null"`
	FollowCount   int64  `gorm:"column:follow_count;default:0"`
	FollowerCount int64  `gorm:"column:follower_count;default:0"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func QueryUserByIds(ctx context.Context, userIds []int64) ([]*User, error) {
	var users []*User
	err := db.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return users, nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := db.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return db.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := db.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
