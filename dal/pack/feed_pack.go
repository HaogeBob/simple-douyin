package pack

import (
	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/feed"
)

func Feed_pack(video *db.Videos) (feed.Video, error) {

	var tt feed.Video

	tt.Id = video.Id
	tt.Title = video.Title
	tt.CoverUrl = video.Cover_url
	tt.PlayUrl = video.Play_url

	tt.CommentCount = 0
	tt.FavoriteCount = 0

	tt.IsFavorite = true

	return tt, nil

}
