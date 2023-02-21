// Copyright 2022 Tik-Tok Authors
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

package main

import (
	"context"

	"github.com/simple/douyin/cmd/favorite/pack"
	"github.com/simple/douyin/cmd/favorite/service"
	"github.com/simple/douyin/kitex_gen/favorite"
	"github.com/simple/douyin/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteAction interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteList interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)

	// if err = req.IsValid(); err != nil {
	// 	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	// 	return resp, nil
	// }

	videos, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
