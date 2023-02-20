package service

import (
	"context"
	"errors"
	"math"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/feed"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type FeedService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s FeedService) Feed(token string, latestTime int64) ([]*feed.Video, int64, error) {

	Jwt := jwt.NewJWT([]byte(constants.SecretKey))

	currentId, _ := Jwt.CheckToken(token)
	if currentId <= 0 {
		return nil, 0, errors.New("token compare error")
	}

	user, err := db.Query_user_info(s.ctx, currentId)

	if err != nil {
		return nil, 0, err
	}

	videos, err := db.Query_vedio(s.ctx, user.Id, latestTime)
	if err != nil {
		return nil, 0, err
	}

	var req = make([]*feed.Video, len(videos))

	user_info := feed.User{Id: user.Id, FollowCount: user.Follow_count,
		FollowerCount: user.Follower_count, IsFollow: user.Is_follow}

	var time = int64(math.MaxInt64)
	var temp feed.Video

	for i := 0; i < len(videos); i++ {
		temp.Author = &user_info
		temp.Id = videos[i].Id
		temp.Title = videos[i].Title

		temp.CommentCount = videos[i].Comment_count
		temp.FavoriteCount = videos[i].Favorite_count
		temp.CoverUrl = videos[i].Cover_url
		temp.PlayUrl = videos[i].Play_url
		temp.IsFavorite = videos[i].Is_favorite

		if videos[i].Release_time < time {
			time = videos[i].Release_time
		}

		req = append(req, &temp)
	}

	return req, time, nil
}
