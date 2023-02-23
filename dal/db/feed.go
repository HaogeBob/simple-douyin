package db

import (
	"context"
	"fmt"
	"time"

	"github.com/simple/douyin/pkg/constants"
)

type Videos struct {
	//gorm.Model
	Id           int64  `json:"id"`
	Author_id    int64  `json:"author_id"`
	Play_url     string `json:"play_url"`
	Cover_url    string `json:"cover_url"`
	Title        string `json:"title"`
	Release_time string `json:"release_time"`
}

func (v *Videos) TableName() string {
	return constants.VideoTableName
}

func Query_vedio(ctx context.Context, author_id int64, time_d int64) ([]*Videos, error) {

	// 将时间戳转化为时间查
	t := time.Unix(time_d, 0).Format(constants.TimeFormat)
	res := make([]*Videos, 0)
	fmt.Println(t)
	if err := DB.WithContext(ctx).Where("author_id = ? and release_time >= ?", author_id, t).Order("release_time").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil

}

/*

	Favorite_count int64  `json:"favorite_count"`
	Comment_count  int64  `json:"comment_count"`
	Is_favorite    bool   `json:"is_favorite"`

*/
