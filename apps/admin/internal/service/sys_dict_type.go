package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysdicttype/v1"
)

type SysDictTypeService struct {
	pb.UnimplementedSysDictTypeServer
	uc  *biz.SysDictTypeUseCase
	log *log.Helper
}

func NewSysDictTypeService(uc *biz.SysDictTypeUseCase, logger log.Logger) *SysDictTypeService {
	return &SysDictTypeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDictTypeService) GetPageSysDictType(ctx context.Context, req *pb.SysDictTypePageReq) (*pb.SysDictTypePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysDictType")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysDictTypeReply, 0)
	for i := range datas {
		items = append(items, convert.SysDictTypeData2Reply(datas[i]))
	}
	reply := &pb.SysDictTypePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysDictTypeService) GetSysDictType(ctx context.Context, req *pb.SysDictTypeReq) (*pb.SysDictTypeReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysDictType")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysDictTypeData2Reply(daya), err
}
func (s *SysDictTypeService) UpdateSysDictType(ctx context.Context, req *pb.SysDictTypeUpdateReq) (*pb.SysDictTypeUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysDictType")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysDictTypeData2UpdateReply(data), err
}
func (s *SysDictTypeService) CreateSysDictType(ctx context.Context, req *pb.SysDictTypeCreateReq) (*pb.SysDictTypeCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysDictType")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysDictTypeData2CreateReply(data), err
}
func (s *SysDictTypeService) DeleteSysDictType(ctx context.Context, req *pb.SysDictTypeDeleteReq) (*pb.SysDictTypeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysDictType")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysDictTypeDeleteReply{Result: err == nil}, err
}
func (s *SysDictTypeService) BatchDeleteSysDictType(ctx context.Context, req *pb.SysDictTypeBatchDeleteReq) (*pb.SysDictTypeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysDictType")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysDictTypeDeleteReply{Result: err == nil && num > 0}, err
}
