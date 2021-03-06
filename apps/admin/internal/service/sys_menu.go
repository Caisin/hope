package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysmenu/v1"
)

type SysMenuService struct {
	pb.UnimplementedSysMenuServer
	uc  *biz.SysMenuUseCase
	log *log.Helper
}

func NewSysMenuService(uc *biz.SysMenuUseCase, logger log.Logger) *SysMenuService {
	return &SysMenuService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysMenuService) GetSysMenuPage(ctx context.Context, req *pb.SysMenuPageReq) (*pb.SysMenuPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysMenuPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysMenuData, 0)
	for i := range datas {
		items = append(items, convert.SysMenuData2Reply(datas[i]))
	}
	reply := &pb.SysMenuPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysMenuService) GetSysMenu(ctx context.Context, req *pb.SysMenuReq) (*pb.SysMenuReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysMenu")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysMenuReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysMenuData2Reply(data),
	}
	return reply, nil
}
func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *pb.SysMenuUpdateReq) (*pb.SysMenuUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysMenu")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysMenuUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysMenuData2Reply(data),
	}
	return reply, nil
}
func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *pb.SysMenuCreateReq) (*pb.SysMenuCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysMenu")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysMenuCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysMenuData2Reply(data),
	}
	return reply, err
}
func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *pb.SysMenuDeleteReq) (*pb.SysMenuDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysMenu")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysMenuDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysMenuService) BatchDeleteSysMenu(ctx context.Context, req *pb.SysMenuBatchDeleteReq) (*pb.SysMenuDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysMenu")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysMenuDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
