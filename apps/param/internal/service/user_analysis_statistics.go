package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/useranalysisstatistics/v1"
)

type UserAnalysisStatisticsService struct {
	pb.UnimplementedUserAnalysisStatisticsServer
	uc  *biz.UserAnalysisStatisticsUseCase
	log *log.Helper
}

func NewUserAnalysisStatisticsService(uc *biz.UserAnalysisStatisticsUseCase, logger log.Logger) *UserAnalysisStatisticsService {
	return &UserAnalysisStatisticsService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserAnalysisStatisticsService) GetUserAnalysisStatisticsPage(ctx context.Context, req *pb.UserAnalysisStatisticsPageReq) (*pb.UserAnalysisStatisticsPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserAnalysisStatisticsPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserAnalysisStatisticsData, 0)
	for i := range datas {
		items = append(items, convert.UserAnalysisStatisticsData2Reply(datas[i]))
	}
	reply := &pb.UserAnalysisStatisticsPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserAnalysisStatisticsService) GetUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsReq) (*pb.UserAnalysisStatisticsReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserAnalysisStatistics")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserAnalysisStatisticsReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserAnalysisStatisticsData2Reply(data),
	}
	return reply, nil
}
func (s *UserAnalysisStatisticsService) UpdateUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsUpdateReq) (*pb.UserAnalysisStatisticsUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserAnalysisStatistics")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserAnalysisStatisticsUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserAnalysisStatisticsData2Reply(data),
	}
	return reply, nil
}
func (s *UserAnalysisStatisticsService) CreateUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsCreateReq) (*pb.UserAnalysisStatisticsCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserAnalysisStatistics")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserAnalysisStatisticsCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserAnalysisStatisticsData2Reply(data),
	}
	return reply, err
}
func (s *UserAnalysisStatisticsService) DeleteUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsDeleteReq) (*pb.UserAnalysisStatisticsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserAnalysisStatistics")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserAnalysisStatisticsDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserAnalysisStatisticsService) BatchDeleteUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsBatchDeleteReq) (*pb.UserAnalysisStatisticsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserAnalysisStatistics")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserAnalysisStatisticsDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
