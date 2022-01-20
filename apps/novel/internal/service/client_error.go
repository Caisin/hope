package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/clienterror/v1"
)

type ClientErrorService struct {
	pb.UnimplementedClientErrorServer
	uc  *biz.ClientErrorUseCase
	log *log.Helper
}

func NewClientErrorService(uc *biz.ClientErrorUseCase, logger log.Logger) *ClientErrorService {
	return &ClientErrorService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ClientErrorService) GetPageClientError(ctx context.Context, req *pb.ClientErrorPageReq) (*pb.ClientErrorPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ClientErrorReply, 0)
	for i := range datas {
		items = append(items, convert.ClientErrorData2Reply(datas[i]))
	}
	reply := &pb.ClientErrorPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ClientErrorService) GetClientError(ctx context.Context, req *pb.ClientErrorReq) (*pb.ClientErrorReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ClientErrorData2Reply(daya), err
}
func (s *ClientErrorService) UpdateClientError(ctx context.Context, req *pb.ClientErrorUpdateReq) (*pb.ClientErrorUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ClientErrorData2UpdateReply(data), err
}
func (s *ClientErrorService) CreateClientError(ctx context.Context, req *pb.ClientErrorCreateReq) (*pb.ClientErrorCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ClientErrorData2CreateReply(data), err
}
func (s *ClientErrorService) DeleteClientError(ctx context.Context, req *pb.ClientErrorDeleteReq) (*pb.ClientErrorDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ClientErrorDeleteReply{Result: err == nil}, err
}
func (s *ClientErrorService) BatchDeleteClientError(ctx context.Context, req *pb.ClientErrorBatchDeleteReq) (*pb.ClientErrorDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ClientErrorDeleteReply{Result: err == nil && num > 0}, err
}
