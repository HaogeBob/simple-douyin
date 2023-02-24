package service

import (
	"context"
	"errors"

	"github.com/simple/douyin/dal/db"
	"github.com/simple/douyin/kitex_gen/relation"
	"github.com/simple/douyin/pkg/constants"
	"github.com/simple/douyin/pkg/jwt"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return err
	}
	// 检验用户是否存在
	users, err := db.QueryUserByIds(s.ctx, []int64{currentId})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("User not exist")
	}
	
	// 检验关注对象是否存在
	users, err = db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("ToUser not exist")
	}

	// 关注、取关
	if req.ActionType == 1 {
		err := db.Create(s.ctx, currentId, req.ToUserId)
		if err != nil {
			return err
		}
		return nil
	} else if req.ActionType == 2 {
		err := db.Delete(s.ctx, currentId, req.ToUserId)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("ActionType not valid")
	}
}
