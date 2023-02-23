package service

import (
	"context"
	"errors"

	"github.com/simple/douyin/dal/db"
	// "github.com/simple/douyin/dal/pack"
	"github.com/simple/douyin/kitex_gen/relation"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *relation.FollowerListRequest) ([]*relation.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	if currentId != req.UserId {
		return nil, errors.New("Token compare error")
	}
	// 检查当前用户是否存在
	users, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("User not exist")
	}
	// 获取粉丝联系 -> 粉丝ID -> 粉丝信息
	followers, err := db.QueryFollowerById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	followerIds := make([]int64, 0)
	for _, follower := range followers {
		followerIds = append(followerIds, follower.UserId)
	}

	followerUsers, err := db.QueryUserByIds(s.ctx, followerIds)
	if err != nil {
		return nil, err
	}

	// 打包返回值
	userList := make([]*relation.User, 0)
	for _, user := range followerUsers {
		userList = append(userList, &relation.User{
			Id:            int64(user.ID),
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		})
	}
	return userList, nil
}