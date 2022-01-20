package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/activity/v1"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
	uc  *biz.ActivityUseCase
	log *log.Helper
}

func NewActivityService(uc *biz.ActivityUseCase, logger log.Logger) *ActivityService {
	return &ActivityService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ActivityService) GetPageActivity(ctx context.Context, req *pb.ActivityPageReq) (*pb.ActivityPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ActivityReply, 0)
	for i := range datas {
		items = append(items, convert.ActivityData2Reply(datas[i]))
	}
	reply := &pb.ActivityPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ActivityService) GetActivity(ctx context.Context, req *pb.ActivityReq) (*pb.ActivityReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ActivityData2Reply(daya), err
}
func (s *ActivityService) UpdateActivity(ctx context.Context, req *pb.ActivityUpdateReq) (*pb.ActivityUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ActivityData2UpdateReply(data), err
}
func (s *ActivityService) CreateActivity(ctx context.Context, req *pb.ActivityCreateReq) (*pb.ActivityCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ActivityData2CreateReply(data), err
}
func (s *ActivityService) DeleteActivity(ctx context.Context, req *pb.ActivityDeleteReq) (*pb.ActivityDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ActivityDeleteReply{Result: err == nil}, err
}
func (s *ActivityService) BatchDeleteActivity(ctx context.Context, req *pb.ActivityBatchDeleteReq) (*pb.ActivityDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ActivityDeleteReply{Result: err == nil && num > 0}, err
}
