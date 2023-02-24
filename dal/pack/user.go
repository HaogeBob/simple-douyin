package pack

import (
	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/user"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{Id: int64(u.ID), Name: u.Username, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
