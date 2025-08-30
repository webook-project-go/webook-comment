package service

import (
	"context"
	"github.com/webook-project-go/webook-comment/domain"
	"github.com/webook-project-go/webook-comment/repository"
)

type Service interface {
	GetList(ctx context.Context, biz string, bizID int64, minID int64, limit int) ([]domain.Comment, error)
	GetReplies(ctx context.Context, pid int64, minID int64, limit int) ([]domain.Comment, error)

	Create(ctx context.Context, comment domain.Comment) (int64, error)
	Reply(ctx context.Context, comment domain.Comment) (int64, error)
	Delete(ctx context.Context, id int64, uid int64) error
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

type service struct {
	repo repository.Repository
}

func (s *service) GetList(ctx context.Context, biz string, bizID int64, minID int64, limit int) ([]domain.Comment, error) {
	return s.repo.GetList(ctx, biz, bizID, minID, limit)
}

func (s *service) GetReplies(ctx context.Context, pid int64, minID int64, limit int) ([]domain.Comment, error) {
	return s.repo.GetReplies(ctx, pid, minID, limit)
}

func (s *service) Create(ctx context.Context, comment domain.Comment) (int64, error) {
	return s.repo.Create(ctx, comment)
}

func (s *service) Reply(ctx context.Context, comment domain.Comment) (int64, error) {
	return s.repo.Create(ctx, comment)
}

func (s *service) Delete(ctx context.Context, id int64, uid int64) error {
	return s.repo.Remove(ctx, id, uid)
}
