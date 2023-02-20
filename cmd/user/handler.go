package main

import (
	"context"

	"github.com/simple/douyin/cmd/user/service"
	user "github.com/simple/douyin/kitex_gen/user"
	"github.com/simple/douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = errno.ActionTypeErr.ErrCode
		resp.StatusMsg = errno.ActionTypeErr.ErrMsg
		// resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = errno.UserAlreadyExistErr.ErrCode
		resp.StatusMsg = errno.UserAlreadyExistErr.ErrMsg
		// resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	// resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = errno.ActionTypeErr.ErrCode
		resp.StatusMsg = errno.ActionTypeErr.ErrMsg
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.StatusCode = errno.LoginErr.ErrCode
		resp.StatusMsg = errno.LoginErr.ErrMsg
		return resp, nil
	}

	resp.UserId = uid
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MGetUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = errno.ActionTypeErr.ErrCode
		resp.StatusMsg = errno.ActionTypeErr.ErrMsg
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.StatusCode = errno.ActionTypeErr.ErrCode
		resp.StatusMsg = errno.ActionTypeErr.ErrMsg
		return resp, nil
	}
	if len(users) == 0 {
		resp.StatusCode = errno.ParamParseErr.ErrCode
		resp.StatusMsg = errno.ParamParseErr.ErrMsg
		return resp, nil
	}

	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.User = users[0]
	return resp, nil
}
