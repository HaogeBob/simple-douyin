package pack

import (
	"errors"

	"github.com/simple/douyin/kitex_gen/publish"
	"github.com/simple/douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *publish.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *publish.BaseResp {
	return &publish.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
