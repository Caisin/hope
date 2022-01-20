package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysjob/v1"
)

type SysJobService struct {
	pb.UnimplementedSysJobServer
	uc  *biz.SysJobUseCase
	log *log.Helper
}

func NewSysJobService(uc *biz.SysJobUseCase, logger log.Logger) *SysJobService {
	return &SysJobService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysJobService) GetPageSysJob(ctx context.Context, req *pb.SysJobPageReq) (*pb.SysJobPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysJobReply, 0)
	for i := range datas {
		items = append(items, convert.SysJobData2Reply(datas[i]))
	}
	reply := &pb.SysJobPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysJobService) GetSysJob(ctx context.Context, req *pb.SysJobReq) (*pb.SysJobReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysJobData2Reply(daya), err
}
func (s *SysJobService) UpdateSysJob(ctx context.Context, req *pb.SysJobUpdateReq) (*pb.SysJobUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysJobData2UpdateReply(data), err
}
func (s *SysJobService) CreateSysJob(ctx context.Context, req *pb.SysJobCreateReq) (*pb.SysJobCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysJobData2CreateReply(data), err
}
func (s *SysJobService) DeleteSysJob(ctx context.Context, req *pb.SysJobDeleteReq) (*pb.SysJobDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysJobDeleteReply{Result: err == nil}, err
}
func (s *SysJobService) BatchDeleteSysJob(ctx context.Context, req *pb.SysJobBatchDeleteReq) (*pb.SysJobDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysJobDeleteReply{Result: err == nil && num > 0}, err
}
