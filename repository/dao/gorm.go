package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Dao interface {
	FindByBiz(ctx context.Context, biz string, bizID int64, mindID int64, limit int) ([]Comment, error)
	FindByPid(ctx context.Context, pid int64, minID int64, limit int) ([]Comment, error)
	Create(ctx context.Context, comment Comment) (int64, error)
	Remove(ctx context.Context, id int64, uid int64) error
}

func NewDao(db *gorm.DB) Dao {
	return &dao{
		db: db,
	}
}

type dao struct {
	db *gorm.DB
}

func (d *dao) Create(ctx context.Context, comment Comment) (int64, error) {
	err := d.db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

func (d *dao) Remove(ctx context.Context, id int64, uid int64) error {
	res := d.db.WithContext(ctx).Where("id = ? AND uid = ?", id, uid).Delete(&Comment{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no such comment")
	}
	return nil
}

func (d *dao) FindByPid(ctx context.Context, pid int64, minID int64, limit int) ([]Comment, error) {
	var res []Comment
	err := d.db.WithContext(ctx).Where("pid = ? AND id < ?", pid, minID).
		Limit(limit).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *dao) FindByBiz(ctx context.Context, biz string, bizID int64, mindID int64, limit int) ([]Comment, error) {
	var res []Comment
	err := d.db.WithContext(ctx).Where("biz = ? AND biz_id = ? AND id > ?", biz, bizID, mindID).
		Limit(limit).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
