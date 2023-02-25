package service

import (
	"context"
	"errors"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/relation"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

func (s *FollowListService) FollowList(req *relation.FollowListRequest) ([]*relation.User, error) {
	// Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	// currentId, _ := Jwt.CheckToken(req.Token)

	// if currentId != req.UserId {
	// 	return nil, errors.New("Token compare error")
	// }
	// 检查当前用户是否存在
	users, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("User not exist")
	}
	// 获取关注联系 -> 关注者ID -> 关注者信息
	follows, err := db.QueryFollowById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	followIds := make([]int64, 0)
	for _, follow := range follows {
		followIds = append(followIds, follow.ToUserId)
	}

	followUsers, err := db.QueryUserByIds(s.ctx, followIds)
	if err != nil {
		return nil, err
	}

	// 打包返回值
	userList := make([]*relation.User, 0)
	for _, user := range followUsers {
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
