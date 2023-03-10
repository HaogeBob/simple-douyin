package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/simple/douyin/dal/db"

	"github.com/simple/douyin/kitex_gen/user"
	"github.com/simple/douyin/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNameValidationErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.PasswordValidationErr
	}
	return int64(u.ID), nil
}
