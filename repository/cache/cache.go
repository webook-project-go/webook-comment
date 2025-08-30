package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/webook-project-go/webook-comment/domain"
	"time"
)

type Cache interface {
	GetList(ctx context.Context, biz string, bizID int64) ([]domain.Comment, error)
	SetList(ctx context.Context, biz string, bizID int64, comments []domain.Comment) error

	GetReplies(ctx context.Context, pid int64) ([]domain.Comment, error)
	SetReplies(ctx context.Context, pid int64, comments []domain.Comment) error
}
type redisCache struct {
	client redis.Cmdable
}

func NewRedisCache(client redis.Cmdable) Cache {
	return &redisCache{client: client}
}

func (r *redisCache) keyBiz(biz string, bizID int64) string {
	return fmt.Sprintf("comment:biz:%s:%d", biz, bizID)
}
func (r *redisCache) keyReplies(pid int64) string {
	return fmt.Sprintf("comment:replies:%d", pid)
}

func (r *redisCache) GetList(ctx context.Context, biz string, bizID int64) ([]domain.Comment, error) {
	key := r.keyBiz(biz, bizID)
	data, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if e := json.Unmarshal(data, &comments); e != nil {
		return nil, e
	}
	return comments, nil
}

func (r *redisCache) SetList(ctx context.Context, biz string, bizID int64, comments []domain.Comment) error {
	key := r.keyBiz(biz, bizID)
	data, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, time.Minute*10).Err()
}

func (r *redisCache) GetReplies(ctx context.Context, pid int64) ([]domain.Comment, error) {
	key := r.keyReplies(pid)
	data, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if e := json.Unmarshal(data, &comments); e != nil {
		return nil, e
	}
	return comments, nil
}

func (r *redisCache) SetReplies(ctx context.Context, pid int64, comments []domain.Comment) error {
	key := r.keyReplies(pid)
	data, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, time.Minute*10).Err()
}
