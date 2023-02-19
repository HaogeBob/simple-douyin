package service

import (
	"context"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/dal/pack"

	"github.com/simple/douyin/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
