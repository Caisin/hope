package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/usermsg/v1"
)

type UserMsgService struct {
	pb.UnimplementedUserMsgServer
	uc  *biz.UserMsgUseCase
	log *log.Helper
}

func NewUserMsgService(uc *biz.UserMsgUseCase, logger log.Logger) *UserMsgService {
	return &UserMsgService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserMsgService) GetUserMsgPage(ctx context.Context, req *pb.UserMsgPageReq) (*pb.UserMsgPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserMsgPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserMsgData, 0)
	for i := range datas {
		items = append(items, convert.UserMsgData2Reply(datas[i]))
	}
	reply := &pb.UserMsgPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserMsgService) GetUserMsg(ctx context.Context, req *pb.UserMsgReq) (*pb.UserMsgReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserMsg")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserMsgReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserMsgData2Reply(data),
	}
	return reply, nil
}
func (s *UserMsgService) UpdateUserMsg(ctx context.Context, req *pb.UserMsgUpdateReq) (*pb.UserMsgUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserMsg")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserMsgUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserMsgData2Reply(data),
	}
	return reply, nil
}
func (s *UserMsgService) CreateUserMsg(ctx context.Context, req *pb.UserMsgCreateReq) (*pb.UserMsgCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserMsg")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserMsgCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserMsgData2Reply(data),
	}
	return reply, err
}
func (s *UserMsgService) DeleteUserMsg(ctx context.Context, req *pb.UserMsgDeleteReq) (*pb.UserMsgDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserMsg")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserMsgDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserMsgService) BatchDeleteUserMsg(ctx context.Context, req *pb.UserMsgBatchDeleteReq) (*pb.UserMsgDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserMsg")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserMsgDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
