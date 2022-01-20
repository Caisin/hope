package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.CasbinRuleData2Reply(daya), err
}
func (s *CasbinRuleService) UpdateCasbinRule(ctx context.Context, req *pb.CasbinRuleUpdateReq) (*pb.CasbinRuleUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.CasbinRuleData2UpdateReply(data), err
}
func (s *CasbinRuleService) CreateCasbinRule(ctx context.Context, req *pb.CasbinRuleCreateReq) (*pb.CasbinRuleCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.CasbinRuleData2CreateReply(data), err
}
func (s *CasbinRuleService) DeleteCasbinRule(ctx context.Context, req *pb.CasbinRuleDeleteReq) (*pb.CasbinRuleDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.CasbinRuleDeleteReply{Result: err == nil}, err
}
func (s *CasbinRuleService) BatchDeleteCasbinRule(ctx context.Context, req *pb.CasbinRuleBatchDeleteReq) (*pb.CasbinRuleDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.CasbinRuleDeleteReply{Result: err == nil && num > 0}, err
}
