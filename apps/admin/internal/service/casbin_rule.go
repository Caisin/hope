package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/casbinrule/v1"
)

type CasbinRuleService struct {
	pb.UnimplementedCasbinRuleServer
	uc  *biz.CasbinRuleUseCase
	log *log.Helper
}

func NewCasbinRuleService(uc *biz.CasbinRuleUseCase, logger log.Logger) *CasbinRuleService {
	return &CasbinRuleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CasbinRuleService) GetPageCasbinRule(ctx context.Context, req *pb.CasbinRulePageReq) (*pb.CasbinRulePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageCasbinRule")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.CasbinRuleReply, 0)
	for i := range datas {
		items = append(items, convert.CasbinRuleData2Reply(datas[i]))
	}
	reply := &pb.CasbinRulePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *CasbinRuleService) GetCasbinRule(ctx context.Context, req *pb.CasbinRuleReq) (*pb.CasbinRuleReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCasbinRule")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CasbinRuleData2Reply(daya), err
}
func (s *CasbinRuleService) UpdateCasbinRule(ctx context.Context, req *pb.CasbinRuleUpdateReq) (*pb.CasbinRuleUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateCasbinRule")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CasbinRuleData2UpdateReply(data), err
}
func (s *CasbinRuleService) CreateCasbinRule(ctx context.Context, req *pb.CasbinRuleCreateReq) (*pb.CasbinRuleCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateCasbinRule")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CasbinRuleData2CreateReply(data), err
}
func (s *CasbinRuleService) DeleteCasbinRule(ctx context.Context, req *pb.CasbinRuleDeleteReq) (*pb.CasbinRuleDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteCasbinRule")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CasbinRuleDeleteReply{Result: err == nil}, err
}
func (s *CasbinRuleService) BatchDeleteCasbinRule(ctx context.Context, req *pb.CasbinRuleBatchDeleteReq) (*pb.CasbinRuleDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteCasbinRule")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CasbinRuleDeleteReply{Result: err == nil && num > 0}, err
}
