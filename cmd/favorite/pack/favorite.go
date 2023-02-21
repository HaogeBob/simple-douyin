// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"github.com/simple/douyin/dal/db"
)

// VideoIds pack Videos info
// in: []*db.Favorite
// out:[]int64 videoIds
func VideoIds(m []*db.Favorite) []int64 {
	videoIds := make([]int64, 0)
	if len(m) == 0 {
		return videoIds
	}

	for i := 0; i < len(m); i++ {
		videoIds = append(videoIds, m[i].FavoriteVideoId)
	}
	return videoIds
}

// UserIds pack Users info
// in:[]int64 videoIds
// out:[]int64 userIds
func UserIds(videos []*db.Video) []int64 {
	userIds := make([]int64, 0)
	if len(videos) == 0 {
		return userIds
	}

	for i := 0; i < len(videos); i++ {
		userIds = append(userIds, videos[i].AuthorId)
	}
	return userIds
}
