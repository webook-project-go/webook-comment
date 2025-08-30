package repository

import (
	"context"
	"database/sql"
	"github.com/kisara71/GoTemplate/slice"
	"github.com/webook-project-go/webook-comment/domain"
	"github.com/webook-project-go/webook-comment/repository/cache"
	"github.com/webook-project-go/webook-comment/repository/dao"
	"time"
)

type Repository interface {
	GetList(ctx context.Context, biz string, bizID int64, minID int64, limit int) ([]domain.Comment, error)
	GetReplies(ctx context.Context, pid int64, minID int64, limit int) ([]domain.Comment, error)

	Create(ctx context.Context, comment domain.Comment) (int64, error)
	CreateReply(ctx context.Context, comment domain.Comment) (int64, error)
	Remove(ctx context.Context, id int64, uid int64) error
}

type repository struct {
	d     dao.Dao
	cache cache.Cache
}

func NewRepository(d dao.Dao, cache cache.Cache) Repository {
	return &repository{
		d:     d,
		cache: cache,
	}
}
func toDomain(comment dao.Comment) domain.Comment {
	return domain.Comment{
		ID:      comment.ID,
		UID:     comment.UID,
		PID:     comment.PID.Int64,
		RID:     comment.RID.Int64,
		Biz:     comment.Biz,
		BizID:   comment.BizID,
		Content: comment.Content,
		Ctime:   time.UnixMilli(comment.Ctime),
	}
}
func toEntity(comment domain.Comment) dao.Comment {
	return dao.Comment{
		ID:  comment.ID,
		UID: comment.UID,
		RID: sql.NullInt64{
			Int64: comment.RID,
			Valid: comment.RID > 0,
		},
		PID: sql.NullInt64{
			Int64: comment.PID,
			Valid: comment.PID > 0,
		},
		Biz:     comment.Biz,
		BizID:   comment.BizID,
		Content: comment.Content,
		Ctime:   comment.Ctime.UnixMilli(),
	}
}
func (r *repository) GetList(ctx context.Context, biz string, bizID int64, minID int64, limit int) ([]domain.Comment, error) {
	if minID == 0 {
		comments, err := r.cache.GetList(ctx, biz, bizID)
		if err == nil && len(comments) > 0 {
			if len(comments) > limit {
				return comments[:limit], nil
			}
			return comments, nil
		}
	}

	res, err := r.d.FindByBiz(ctx, biz, bizID, minID, limit)
	if err != nil {
		return nil, err
	}
	comments, err := slice.Map[dao.Comment, domain.Comment](0, len(res), res, toDomain)
	if err != nil {
		return nil, err
	}
	_ = r.cache.SetList(ctx, biz, bizID, comments)
	return comments, nil
}

func (r *repository) GetReplies(ctx context.Context, pid int64, minID int64, limit int) ([]domain.Comment, error) {
	if minID == 0 {
		comments, err := r.cache.GetReplies(ctx, pid)
		if err == nil && len(comments) > 0 {
			if len(comments) > limit {
				return comments[:limit], nil
			}
			return comments, nil
		}
	}
	res, err := r.d.FindByPid(ctx, pid, minID, limit)
	if err != nil {
		return nil, err
	}
	comments, err := slice.Map[dao.Comment, domain.Comment](0, len(res), res, toDomain)
	if err != nil {
		return nil, err
	}
	_ = r.cache.SetReplies(ctx, pid, comments)
	return comments, nil
}

func (r *repository) Create(ctx context.Context, comment domain.Comment) (int64, error) {
	entity := toEntity(comment)
	return r.d.Create(ctx, entity)
}

func (r *repository) CreateReply(ctx context.Context, comment domain.Comment) (int64, error) {
	entity := toEntity(comment)
	return r.d.Create(ctx, entity)
}

func (r *repository) Remove(ctx context.Context, id int64, uid int64) error {
	return r.d.Remove(ctx, id, uid)
}
