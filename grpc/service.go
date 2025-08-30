package grpc

import (
	"context"
	"github.com/kisara71/GoTemplate/slice"
	"github.com/webook-project-go/webook-apis/gen/go/apis/comment/v1"
	"github.com/webook-project-go/webook-comment/domain"
	"github.com/webook-project-go/webook-comment/service"
	"time"
)

type Service struct {
	svc service.Service
	v1.UnimplementedCommentServiceServer
}

func NewService(svc service.Service) *Service {
	return &Service{
		svc: svc,
	}
}
func toGrpcComment(comment domain.Comment) *v1.Comment {
	return &v1.Comment{
		Id:      comment.ID,
		Uid:     comment.UID,
		Pid:     comment.PID,
		Rid:     comment.RID,
		Biz:     comment.Biz,
		BizId:   comment.BizID,
		Content: comment.Content,
		Ctime:   comment.Ctime.UnixMilli(),
	}
}
func toDomain(comment *v1.Comment) domain.Comment {
	return domain.Comment{
		ID:      comment.Id,
		UID:     comment.Uid,
		PID:     comment.Pid,
		RID:     comment.Rid,
		Biz:     comment.Biz,
		BizID:   comment.BizId,
		Content: comment.Content,
		Ctime:   time.UnixMilli(comment.Ctime),
	}
}
func (s *Service) GetList(ctx context.Context, request *v1.GetListRequest) (*v1.GetListResponse, error) {
	res, err := s.svc.GetList(ctx, request.GetBiz(), request.GetBizId(), request.GetMinId(), int(request.GetLimit()))
	if err != nil {
		return nil, err
	}
	comments, err := slice.Map(0, len(res), res, toGrpcComment)
	if err != nil {
		return nil, err
	}
	return &v1.GetListResponse{
		Comments: comments,
	}, nil
}

func (s *Service) GetReplies(ctx context.Context, request *v1.GetRepliesRequest) (*v1.GetRepliesResponse, error) {
	res, err := s.svc.GetReplies(ctx, request.GetPid(), request.GetMinId(), int(request.GetLimit()))
	if err != nil {
		return nil, err
	}
	comments, err := slice.Map(0, len(res), res, toGrpcComment)
	if err != nil {
		return nil, err
	}
	return &v1.GetRepliesResponse{Comments: comments}, nil
}

func (s *Service) Create(ctx context.Context, request *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	res, err := s.svc.Create(ctx, toDomain(request.GetComment()))
	if err != nil {
		return nil, err
	}
	return &v1.CreateCommentResponse{
		Id: res,
	}, nil
}

func (s *Service) Reply(ctx context.Context, request *v1.ReplyCommentRequest) (*v1.ReplyCommentResponse, error) {
	res, err := s.svc.Create(ctx, toDomain(request.GetComment()))
	if err != nil {
		return nil, err
	}
	return &v1.ReplyCommentResponse{
		Id: res,
	}, nil
}

func (s *Service) Delete(ctx context.Context, request *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	err := s.svc.Delete(ctx, request.GetId(), request.GetUid())
	if err != nil {
		return nil, err
	}
	return &v1.DeleteCommentResponse{}, nil
}

func (s *Service) mustEmbedUnimplementedCommentServiceServer() {
	//TODO implement me
	panic("implement me")
}
