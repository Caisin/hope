package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysrole/v1"
)

type SysRoleService struct {
	pb.UnimplementedSysRoleServer
	uc  *biz.SysRoleUseCase
	log *log.Helper
}

func NewSysRoleService(uc *biz.SysRoleUseCase, logger log.Logger) *SysRoleService {
	return &SysRoleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysRoleService) GetSysRolePage(ctx context.Context, req *pb.SysRolePageReq) (*pb.SysRolePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysRolePage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysRoleData, 0)
	for i := range datas {
		role := datas[i]
		reply := convert.SysRoleData2Reply(role)
		iDs, err := role.QueryMenus().IDs(ctx)
		if err != nil {
			return nil, err
		}
		reply.MenuIds = iDs
		items = append(items, reply)
	}
	reply := &pb.SysRolePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysRoleService) GetSysRole(ctx context.Context, req *pb.SysRoleReq) (*pb.SysRoleReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysRole")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysRoleReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysRoleData2Reply(data),
	}
	return reply, nil
}
func (s *SysRoleService) UpdateSysRole(ctx context.Context, req *pb.SysRoleUpdateReq) (*pb.SysRoleUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysRole")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysRoleUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysRoleData2Reply(data),
	}
	return reply, nil
}
func (s *SysRoleService) CreateSysRole(ctx context.Context, req *pb.SysRoleCreateReq) (*pb.SysRoleCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysRole")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysRoleCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysRoleData2Reply(data),
	}
	return reply, err
}
func (s *SysRoleService) DeleteSysRole(ctx context.Context, req *pb.SysRoleDeleteReq) (*pb.SysRoleDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysRole")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysRoleDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysRoleService) BatchDeleteSysRole(ctx context.Context, req *pb.SysRoleBatchDeleteReq) (*pb.SysRoleDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysRole")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysRoleDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
