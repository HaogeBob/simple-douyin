package pack

import (
	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/publish"
	"github.com/simple/douyin/kitex_gen/user"
)

func Video(v *db.Video, u *user.User, is_favorite bool) *publish.Video {
	if v == nil {
		return nil
	}
	user := publish.User{
		Id:            int64(u.Id),
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	return &publish.Video{
		Id:            int64(v.ID),
		User:          &user,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavorateCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    is_favorite,
	}
}

func Videos(vs []*db.Video, u *user.User, a []bool) []*publish.Video {
	videos := make([]*publish.Video, 0)
	for i := 0; i < len(vs); i++ {
		video := Video(vs[i], u, a[i])
		videos = append(videos, video)
	}
	return videos
}
