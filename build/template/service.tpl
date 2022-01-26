package service

import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/{{.model}}/internal/biz"
	"hope/apps/{{.model}}/internal/convert"

	pb "hope/api/{{.model}}/{{.pkg}}/v1"
)

type {{.name}}Service struct {
	pb.Unimplemented{{.name}}Server
	uc  *biz.{{.name}}UseCase
	log *log.Helper
}

func New{{.name}}Service(uc *biz.{{.name}}UseCase, logger log.Logger) *{{.name}}Service {
	return &{{.name}}Service{uc: uc, log: log.NewHelper(logger)}
}

func (s *{{.name}}Service) Get{{.name}}Page(ctx context.Context, req *pb.{{.name}}PageReq) (*pb.{{.name}}PageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Get{{.name}}Page")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.{{.name}}Data, 0)
	for i := range datas {
		items = append(items, convert.{{.name}}Data2Reply(datas[i]))
	}
	reply := &pb.{{.name}}PageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items: items,
	}
	return reply, nil
}
func (s *{{.name}}Service) Get{{.name}}(ctx context.Context, req *pb.{{.name}}Req) (*pb.{{.name}}Reply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Get{{.name}}")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.{{.name}}Reply{
		Code:    200,
		Message: "success",
		Result:  convert.{{.name}}Data2Reply(data),
	}
	return reply, nil
}
func (s *{{.name}}Service) Update{{.name}}(ctx context.Context, req *pb.{{.name}}UpdateReq) (*pb.{{.name}}UpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Update{{.name}}")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.{{.name}}UpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.{{.name}}Data2Reply(data),
	}
	return reply, nil
}
func (s *{{.name}}Service) Create{{.name}}(ctx context.Context, req *pb.{{.name}}CreateReq) (*pb.{{.name}}CreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Create{{.name}}")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.{{.name}}CreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.{{.name}}Data2Reply(data),
	}
	return reply, err
}
func (s *{{.name}}Service) Delete{{.name}}(ctx context.Context, req *pb.{{.name}}DeleteReq) (*pb.{{.name}}DeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Delete{{.name}}")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.{{.name}}DeleteReply{Code: 200, Message: "success",Result: err == nil}, err
}
func (s *{{.name}}Service) BatchDelete{{.name}}(ctx context.Context, req *pb.{{.name}}BatchDeleteReq) (*pb.{{.name}}DeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDelete{{.name}}")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.{{.name}}DeleteReply{Code: 200, Message: "success",Result: err == nil && num > 0}, err
}
