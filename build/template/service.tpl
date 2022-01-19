package service

import (
	"context"
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

func New{{.name}}Service(uc *biz.{{.name}}UseCase, log *log.Helper) *{{.name}}Service {
	return &{{.name}}Service{uc: uc, log: log}
}

func (s *{{.name}}Service) GetPage{{.name}}(ctx context.Context, req *pb.{{.name}}PageReq) (*pb.{{.name}}PageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.{{.name}}Reply, 0)
	for i := range datas {
		items = append(items, convert.{{.name}}Data2Reply(datas[i]))
	}
	reply := &pb.{{.name}}PageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *{{.name}}Service) Get{{.name}}(ctx context.Context, req *pb.{{.name}}Req) (*pb.{{.name}}Reply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.{{.name}}Data2Reply(daya), err
}
func (s *{{.name}}Service) Update{{.name}}(ctx context.Context, req *pb.{{.name}}UpdateReq) (*pb.{{.name}}UpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.{{.name}}Data2UpdateReply(data), err
}
func (s *{{.name}}Service) Create{{.name}}(ctx context.Context, req *pb.{{.name}}CreateReq) (*pb.{{.name}}CreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.{{.name}}Data2CreateReply(data), err
}
func (s *{{.name}}Service) Delete{{.name}}(ctx context.Context, req *pb.{{.name}}DeleteReq) (*pb.{{.name}}DeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.{{.name}}DeleteReply{Result: err == nil}, err
}
func (s *{{.name}}Service) BatchDelete{{.name}}(ctx context.Context, req *pb.{{.name}}BatchDeleteReq) (*pb.{{.name}}DeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.{{.name}}DeleteReply{Result: err == nil && num > 0}, err
}
