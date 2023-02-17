// Code generated by hertz generator. DO NOT EDIT.

package Favorite

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	favorite "github.com/simple/douyin/biz/handler/favorite"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action := _favorite.Group("/action", _actionMw()...)
				_action.POST("/", append(_favorite_ctionMw(), favorite.FavoriteAction)...)
			}
			{
				_list := _favorite.Group("/list", _listMw()...)
				_list.GET("/", append(_favoritelistMw(), favorite.FavoriteList)...)
			}
		}
	}
}
