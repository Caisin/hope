package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *UserAnalysisStatisticsService) GetPageUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsPageReq) (*pb.UserAnalysisStatisticsPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserAnalysisStatisticsReply, 0)
	for i := range datas {
		items = append(items, convert.UserAnalysisStatisticsData2Reply(datas[i]))
	}
	reply := &pb.UserAnalysisStatisticsPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserAnalysisStatisticsService) GetUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsReq) (*pb.UserAnalysisStatisticsReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserAnalysisStatisticsData2Reply(daya), err
}
func (s *UserAnalysisStatisticsService) UpdateUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsUpdateReq) (*pb.UserAnalysisStatisticsUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserAnalysisStatisticsData2UpdateReply(data), err
}
func (s *UserAnalysisStatisticsService) CreateUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsCreateReq) (*pb.UserAnalysisStatisticsCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserAnalysisStatisticsData2CreateReply(data), err
}
func (s *UserAnalysisStatisticsService) DeleteUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsDeleteReq) (*pb.UserAnalysisStatisticsDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserAnalysisStatisticsDeleteReply{Result: err == nil}, err
}
func (s *UserAnalysisStatisticsService) BatchDeleteUserAnalysisStatistics(ctx context.Context, req *pb.UserAnalysisStatisticsBatchDeleteReq) (*pb.UserAnalysisStatisticsDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserAnalysisStatisticsDeleteReply{Result: err == nil && num > 0}, err
}
