package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// Relation Gorm Data Structures
type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (RelationRaw) TableName() string {
	return "relation"
}

// 根据当前用户id和关注用户id获取关注信息
// 举例: 用来判断某评论的发表用户是否被当前用户所关注 获取 is_follow 字段值
func QueryRelationByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by userid and touserid fail " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*RelationRaw)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}

// 增加当前用户的关注总数，增加其他用户的粉丝总数，创建关注记录
func Create(ctx context.Context, currentId int64, toUserId int64) error {
	relation := &RelationRaw{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("relation").Create(&relation).Error
		if err != nil {
			klog.Error("create relation fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			klog.Error("add follow_count fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("add follower_count fail " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 减少当前用户的关注总数，减少其他用户的粉丝总数，删除关注记录
func Delete(ctx context.Context, currentId int64, toUserId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("relation").Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&RelationRaw{}).Error
		if err != nil {
			klog.Error("delete relation fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			klog.Error("minus follow_count fail " + err.Error())
			return err
		}
		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("minus follower_count fail " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 通过用户id，查询该用户关注的博主，返回关注记录
func QueryFollowById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by userid fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// 通过用户id，查询该用户的粉丝， 返回关注记录
func QueryFollowerById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by touserid fail " + err.Error())
		return nil, err
	}
	return relations, nil
}
