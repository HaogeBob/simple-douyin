package pack

import (
	"errors"
	"time"

	"github.com/simple/douyin/kitex_gen/comment"
	"github.com/simple/douyin/kitex_gen/favorite"
	"github.com/simple/douyin/kitex_gen/feed"
	"github.com/simple/douyin/kitex_gen/publish"
	"github.com/simple/douyin/kitex_gen/relation"
	"github.com/simple/douyin/kitex_gen/user"
	"github.com/simple/douyin/pkg/errno"
)

func NewFeedBaseResp(err error) *feed.BaseResp {
	if err == nil {
		return feedbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return feedbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return feedbaseResp(s)
}

func feedbaseResp(err errno.ErrNo) *feed.BaseResp {
	return &feed.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func NewPublishBaseResp(err error) *publish.BaseResp {
	if err == nil {
		return publishbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishbaseResp(s)
}

func publishbaseResp(err errno.ErrNo) *publish.BaseResp {
	return &publish.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func NewUserBaseResp(err error) *user.BaseResp {
	if err == nil {
		return userbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userbaseResp(s)
}

func userbaseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func NewFavoriteBaseResp(err error) *favorite.BaseResp {
	if err == nil {
		return favoritebaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoritebaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoritebaseResp(s)
}

func favoritebaseResp(err errno.ErrNo) *favorite.BaseResp {
	return &favorite.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func NewCommentBaseResp(err error) *comment.BaseResp {
	if err == nil {
		return commentbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentbaseResp(s)
}

func commentbaseResp(err errno.ErrNo) *comment.BaseResp {
	return &comment.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func NewRelationBaseResp(err error) *relation.BaseResp {
	if err == nil {
		return relationbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationbaseResp(s)
}

func relationbaseResp(err errno.ErrNo) *relation.BaseResp {
	return &relation.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
