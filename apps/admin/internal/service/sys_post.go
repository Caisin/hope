package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/syspost/v1"
)

type SysPostService struct {
	pb.UnimplementedSysPostServer
	uc  *biz.SysPostUseCase
	log *log.Helper
}

func NewSysPostService(uc *biz.SysPostUseCase, logger log.Logger) *SysPostService {
	return &SysPostService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysPostService) GetPageSysPost(ctx context.Context, req *pb.SysPostPageReq) (*pb.SysPostPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysPost")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysPostData, 0)
	for i := range datas {
		items = append(items, convert.SysPostData2Reply(datas[i]))
	}
	reply := &pb.SysPostPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysPostService) GetSysPost(ctx context.Context, req *pb.SysPostReq) (*pb.SysPostReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysPost")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysPostReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysPostData2Reply(data),
	}
	return reply, nil
}
func (s *SysPostService) UpdateSysPost(ctx context.Context, req *pb.SysPostUpdateReq) (*pb.SysPostUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysPost")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysPostUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysPostData2Reply(data),
	}
	return reply, nil
}
func (s *SysPostService) CreateSysPost(ctx context.Context, req *pb.SysPostCreateReq) (*pb.SysPostCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysPost")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysPostCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysPostData2Reply(data),
	}
	return reply, err
}
func (s *SysPostService) DeleteSysPost(ctx context.Context, req *pb.SysPostDeleteReq) (*pb.SysPostDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysPost")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysPostDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysPostService) BatchDeleteSysPost(ctx context.Context, req *pb.SysPostBatchDeleteReq) (*pb.SysPostDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysPost")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysPostDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
