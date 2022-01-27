package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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

func (s *AgreementLogService) GetAgreementLogPage(ctx context.Context, req *pb.AgreementLogPageReq) (*pb.AgreementLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAgreementLogPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AgreementLogData, 0)
	for i := range datas {
		items = append(items, convert.AgreementLogData2Reply(datas[i]))
	}
	reply := &pb.AgreementLogPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *AgreementLogService) GetAgreementLog(ctx context.Context, req *pb.AgreementLogReq) (*pb.AgreementLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAgreementLog")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AgreementLogReply{
		Code:    200,
		Message: "success",
		Result:  convert.AgreementLogData2Reply(data),
	}
	return reply, nil
}
func (s *AgreementLogService) UpdateAgreementLog(ctx context.Context, req *pb.AgreementLogUpdateReq) (*pb.AgreementLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAgreementLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AgreementLogUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AgreementLogData2Reply(data),
	}
	return reply, nil
}
func (s *AgreementLogService) CreateAgreementLog(ctx context.Context, req *pb.AgreementLogCreateReq) (*pb.AgreementLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAgreementLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AgreementLogCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AgreementLogData2Reply(data),
	}
	return reply, err
}
func (s *AgreementLogService) DeleteAgreementLog(ctx context.Context, req *pb.AgreementLogDeleteReq) (*pb.AgreementLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAgreementLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AgreementLogDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *AgreementLogService) BatchDeleteAgreementLog(ctx context.Context, req *pb.AgreementLogBatchDeleteReq) (*pb.AgreementLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAgreementLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AgreementLogDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
