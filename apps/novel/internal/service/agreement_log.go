package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/agreementlog/v1"
)

type AgreementLogService struct {
	pb.UnimplementedAgreementLogServer
	uc  *biz.AgreementLogUseCase
	log *log.Helper
}

func NewAgreementLogService(uc *biz.AgreementLogUseCase, logger log.Logger) *AgreementLogService {
	return &AgreementLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AgreementLogService) GetPageAgreementLog(ctx context.Context, req *pb.AgreementLogPageReq) (*pb.AgreementLogPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.AgreementLogReply, 0)
	for i := range datas {
		items = append(items, convert.AgreementLogData2Reply(datas[i]))
	}
	reply := &pb.AgreementLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AgreementLogService) GetAgreementLog(ctx context.Context, req *pb.AgreementLogReq) (*pb.AgreementLogReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.AgreementLogData2Reply(daya), err
}
func (s *AgreementLogService) UpdateAgreementLog(ctx context.Context, req *pb.AgreementLogUpdateReq) (*pb.AgreementLogUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.AgreementLogData2UpdateReply(data), err
}
func (s *AgreementLogService) CreateAgreementLog(ctx context.Context, req *pb.AgreementLogCreateReq) (*pb.AgreementLogCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.AgreementLogData2CreateReply(data), err
}
func (s *AgreementLogService) DeleteAgreementLog(ctx context.Context, req *pb.AgreementLogDeleteReq) (*pb.AgreementLogDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.AgreementLogDeleteReply{Result: err == nil}, err
}
func (s *AgreementLogService) BatchDeleteAgreementLog(ctx context.Context, req *pb.AgreementLogBatchDeleteReq) (*pb.AgreementLogDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.AgreementLogDeleteReply{Result: err == nil && num > 0}, err
}
