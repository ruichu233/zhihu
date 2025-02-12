package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
	"time"
	"zhihu/app/follow/internal/svc"
	"zhihu/app/follow/model"
	"zhihu/app/follow/pb/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *FollowActionLogic) FollowAction(in *follow.FollowActionRequest) (*follow.FollowActionResponse, error) {
	// 构建缓存键
	followKey := GetFollowingKey(in.FollowerId)

	// 查询数据库中的关注状态
	isFollowing, err := l.checkDBFollowStatus(in.FollowerId, in.FolloweeId)
	if err != nil {
		return nil, fmt.Errorf("db get error: %w", err)
	}

	// 处理关注操作
	switch in.ActionType {
	case follow.FollowActionRequest_FOLLOW:
		if isFollowing {
			return &follow.FollowActionResponse{Success: true}, nil // 如果已经关注，直接返回
		}
		if err := l.followUser(in.FollowerId, in.FolloweeId, followKey); err != nil {
			return nil, err
		}
	case follow.FollowActionRequest_UNFOLLOW:
		if !isFollowing {
			return &follow.FollowActionResponse{Success: true}, nil // 如果已经取消关注，直接返回
		}
		if err := l.unfollowUser(in.FollowerId, in.FolloweeId, followKey); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown action type: %d", in.ActionType)
	}

	return &follow.FollowActionResponse{Success: true}, nil
}

// 查询数据库中的关注状态
func (l *FollowActionLogic) checkDBFollowStatus(followerId, followeeId int64) (bool, error) {
	var f model.Follow
	err := l.svcCtx.DB.Where("follower_id = ? AND followee_id = ?", followerId, followeeId).First(&f).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if f.DeletedAt.Valid {
		return false, nil
	}
	return true, nil
}

// 关注用户的操作
func (l *FollowActionLogic) followUser(followerId, followeeId int64, followKey string) error {
	if followerId <= 0 || followeeId <= 0 {
		return fmt.Errorf("invalid followerId or followeeId")
	}
	if followKey == "" {
		return fmt.Errorf("followKey cannot be empty")
	}

	return l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// 检查数据库中是否有记录
		f := model.Follow{
			BaseModel: model.BaseModel{
				Id: idgen.NextId(),
			},
			FollowerID: followerId,
			FolloweeID: followeeId,
		}
		// 保存关注状态到数据库
		if err := tx.Save(&f).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("database save error: %w", err)
		}
		// 更新缓存
		if err := l.svcCtx.RDB.ZAdd(l.ctx, followKey, redis.Z{
			Member: followeeId,
			Score:  float64(time.Now().Unix()),
		}).Err(); err != nil {
			tx.Rollback()
			return fmt.Errorf("redis sadd error: %w", err)
		}

		return nil
	})
}

// 取消关注用户的操作
func (l *FollowActionLogic) unfollowUser(followerId, followeeId int64, followKey string) error {
	return l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 查找数据库记录
		var f model.Follow
		if err := tx.Where("follower_id = ? AND followee_id = ?", followerId, followeeId).First(&f).Error; err != nil {
			return fmt.Errorf("database query error: %w", err)
		}

		if err := tx.Where("id = ?", f.Id).Delete(&f).Error; err != nil {
			return fmt.Errorf("database save error: %w", err)
		}

		// 从缓存中移除关注状态
		_ = l.svcCtx.RDB.ZRem(l.ctx, followKey, followeeId).Err()
		return nil
	})
}

// 获取关注列表的缓存键
func GetFollowingKey(userId int64) string {
	return fmt.Sprintf("following_%d", userId)
}

// 获取粉丝列表的缓存键
func GetFollowerKey(userId int64) string {
	return fmt.Sprintf("follower_%d", userId)
}

// 获取朋友列表的缓存键
func GetFriendKey(userId int64) string {
	return fmt.Sprintf("friend_%d", userId)
}
