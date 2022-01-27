package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/listenrecord/v1"
)

type ListenRecordService struct {
	pb.UnimplementedListenRecordServer
	uc  *biz.ListenRecordUseCase
	log *log.Helper
}

func NewListenRecordService(uc *biz.ListenRecordUseCase, logger log.Logger) *ListenRecordService {
	return &ListenRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ListenRecordService) GetListenRecordPage(ctx context.Context, req *pb.ListenRecordPageReq) (*pb.ListenRecordPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetListenRecordPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.ListenRecordData, 0)
	for i := range datas {
		items = append(items, convert.ListenRecordData2Reply(datas[i]))
	}
	reply := &pb.ListenRecordPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *ListenRecordService) GetListenRecord(ctx context.Context, req *pb.ListenRecordReq) (*pb.ListenRecordReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetListenRecord")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListenRecordReply{
		Code:    200,
		Message: "success",
		Result:  convert.ListenRecordData2Reply(data),
	}
	return reply, nil
}
func (s *ListenRecordService) UpdateListenRecord(ctx context.Context, req *pb.ListenRecordUpdateReq) (*pb.ListenRecordUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateListenRecord")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListenRecordUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ListenRecordData2Reply(data),
	}
	return reply, nil
}
func (s *ListenRecordService) CreateListenRecord(ctx context.Context, req *pb.ListenRecordCreateReq) (*pb.ListenRecordCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateListenRecord")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListenRecordCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ListenRecordData2Reply(data),
	}
	return reply, err
}
func (s *ListenRecordService) DeleteListenRecord(ctx context.Context, req *pb.ListenRecordDeleteReq) (*pb.ListenRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteListenRecord")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ListenRecordDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *ListenRecordService) BatchDeleteListenRecord(ctx context.Context, req *pb.ListenRecordBatchDeleteReq) (*pb.ListenRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteListenRecord")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ListenRecordDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
