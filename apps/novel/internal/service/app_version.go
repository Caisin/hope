package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/appversion/v1"
)

type AppVersionService struct {
	pb.UnimplementedAppVersionServer
	uc  *biz.AppVersionUseCase
	log *log.Helper
}

func NewAppVersionService(uc *biz.AppVersionUseCase, logger log.Logger) *AppVersionService {
	return &AppVersionService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AppVersionService) GetPageAppVersion(ctx context.Context, req *pb.AppVersionPageReq) (*pb.AppVersionPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.AppVersionReply, 0)
	for i := range datas {
		items = append(items, convert.AppVersionData2Reply(datas[i]))
	}
	reply := &pb.AppVersionPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AppVersionService) GetAppVersion(ctx context.Context, req *pb.AppVersionReq) (*pb.AppVersionReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.AppVersionData2Reply(daya), err
}
func (s *AppVersionService) UpdateAppVersion(ctx context.Context, req *pb.AppVersionUpdateReq) (*pb.AppVersionUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.AppVersionData2UpdateReply(data), err
}
func (s *AppVersionService) CreateAppVersion(ctx context.Context, req *pb.AppVersionCreateReq) (*pb.AppVersionCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.AppVersionData2CreateReply(data), err
}
func (s *AppVersionService) DeleteAppVersion(ctx context.Context, req *pb.AppVersionDeleteReq) (*pb.AppVersionDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.AppVersionDeleteReply{Result: err == nil}, err
}
func (s *AppVersionService) BatchDeleteAppVersion(ctx context.Context, req *pb.AppVersionBatchDeleteReq) (*pb.AppVersionDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.AppVersionDeleteReply{Result: err == nil && num > 0}, err
}
