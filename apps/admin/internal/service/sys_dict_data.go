package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysdictdata/v1"
)

type SysDictDataService struct {
	pb.UnimplementedSysDictDataServer
	uc  *biz.SysDictDataUseCase
	log *log.Helper
}

func NewSysDictDataService(uc *biz.SysDictDataUseCase, logger log.Logger) *SysDictDataService {
	return &SysDictDataService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDictDataService) GetPageSysDictData(ctx context.Context, req *pb.SysDictDataPageReq) (*pb.SysDictDataPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysDictData")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysDictDataData, 0)
	for i := range datas {
		items = append(items, convert.SysDictDataData2Reply(datas[i]))
	}
	reply := &pb.SysDictDataPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysDictDataService) GetSysDictData(ctx context.Context, req *pb.SysDictDataReq) (*pb.SysDictDataReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysDictData")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysDictDataReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysDictDataData2Reply(data),
	}
	return reply, nil
}
func (s *SysDictDataService) UpdateSysDictData(ctx context.Context, req *pb.SysDictDataUpdateReq) (*pb.SysDictDataUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysDictData")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysDictDataUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysDictDataData2Reply(data),
	}
	return reply, nil
}
func (s *SysDictDataService) CreateSysDictData(ctx context.Context, req *pb.SysDictDataCreateReq) (*pb.SysDictDataCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysDictData")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysDictDataCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysDictDataData2Reply(data),
	}
	return reply, err
}
func (s *SysDictDataService) DeleteSysDictData(ctx context.Context, req *pb.SysDictDataDeleteReq) (*pb.SysDictDataDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysDictData")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysDictDataDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysDictDataService) BatchDeleteSysDictData(ctx context.Context, req *pb.SysDictDataBatchDeleteReq) (*pb.SysDictDataDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysDictData")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysDictDataDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
