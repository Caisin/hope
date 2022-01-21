package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/bookpackage/v1"
)

type BookPackageService struct {
	pb.UnimplementedBookPackageServer
	uc  *biz.BookPackageUseCase
	log *log.Helper
}

func NewBookPackageService(uc *biz.BookPackageUseCase, logger log.Logger) *BookPackageService {
	return &BookPackageService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BookPackageService) GetPageBookPackage(ctx context.Context, req *pb.BookPackagePageReq) (*pb.BookPackagePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageBookPackage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.BookPackageReply, 0)
	for i := range datas {
		items = append(items, convert.BookPackageData2Reply(datas[i]))
	}
	reply := &pb.BookPackagePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *BookPackageService) GetBookPackage(ctx context.Context, req *pb.BookPackageReq) (*pb.BookPackageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetBookPackage")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.BookPackageData2Reply(daya), err
}
func (s *BookPackageService) UpdateBookPackage(ctx context.Context, req *pb.BookPackageUpdateReq) (*pb.BookPackageUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateBookPackage")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.BookPackageData2UpdateReply(data), err
}
func (s *BookPackageService) CreateBookPackage(ctx context.Context, req *pb.BookPackageCreateReq) (*pb.BookPackageCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateBookPackage")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.BookPackageData2CreateReply(data), err
}
func (s *BookPackageService) DeleteBookPackage(ctx context.Context, req *pb.BookPackageDeleteReq) (*pb.BookPackageDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteBookPackage")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.BookPackageDeleteReply{Result: err == nil}, err
}
func (s *BookPackageService) BatchDeleteBookPackage(ctx context.Context, req *pb.BookPackageBatchDeleteReq) (*pb.BookPackageDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteBookPackage")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.BookPackageDeleteReply{Result: err == nil && num > 0}, err
}
