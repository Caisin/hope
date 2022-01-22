package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/datasource/v1"
)

type DataSourceService struct {
	pb.UnimplementedDataSourceServer
	uc  *biz.DataSourceUseCase
	log *log.Helper
}

func NewDataSourceService(uc *biz.DataSourceUseCase, logger log.Logger) *DataSourceService {
	return &DataSourceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *DataSourceService) GetPageDataSource(ctx context.Context, req *pb.DataSourcePageReq) (*pb.DataSourcePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageDataSource")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.DataSourceData, 0)
	for i := range datas {
		items = append(items, convert.DataSourceData2Reply(datas[i]))
	}
	reply := &pb.DataSourcePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *DataSourceService) GetDataSource(ctx context.Context, req *pb.DataSourceReq) (*pb.DataSourceReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetDataSource")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.DataSourceReply{
		Code:    200,
		Message: "success",
		Result:  convert.DataSourceData2Reply(data),
	}
	return reply, nil
}
func (s *DataSourceService) UpdateDataSource(ctx context.Context, req *pb.DataSourceUpdateReq) (*pb.DataSourceUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateDataSource")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.DataSourceUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.DataSourceData2Reply(data),
	}
	return reply, nil
}
func (s *DataSourceService) CreateDataSource(ctx context.Context, req *pb.DataSourceCreateReq) (*pb.DataSourceCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateDataSource")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.DataSourceCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.DataSourceData2Reply(data),
	}
	return reply, err
}
func (s *DataSourceService) DeleteDataSource(ctx context.Context, req *pb.DataSourceDeleteReq) (*pb.DataSourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteDataSource")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DataSourceDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *DataSourceService) BatchDeleteDataSource(ctx context.Context, req *pb.DataSourceBatchDeleteReq) (*pb.DataSourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteDataSource")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DataSourceDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
