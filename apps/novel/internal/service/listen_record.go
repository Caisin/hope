package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/listenrecord/v1"
)

type ListenRecordService struct {
	pb.UnimplementedListenRecordServer
	uc  *biz.ListenRecordUseCase
	log *log.Helper
}

func NewListenRecordService(uc *biz.ListenRecordUseCase, logger log.Logger) *ListenRecordService {
	return &ListenRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ListenRecordService) GetPageListenRecord(ctx context.Context, req *pb.ListenRecordPageReq) (*pb.ListenRecordPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ListenRecordReply, 0)
	for i := range datas {
		items = append(items, convert.ListenRecordData2Reply(datas[i]))
	}
	reply := &pb.ListenRecordPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ListenRecordService) GetListenRecord(ctx context.Context, req *pb.ListenRecordReq) (*pb.ListenRecordReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ListenRecordData2Reply(daya), err
}
func (s *ListenRecordService) UpdateListenRecord(ctx context.Context, req *pb.ListenRecordUpdateReq) (*pb.ListenRecordUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ListenRecordData2UpdateReply(data), err
}
func (s *ListenRecordService) CreateListenRecord(ctx context.Context, req *pb.ListenRecordCreateReq) (*pb.ListenRecordCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ListenRecordData2CreateReply(data), err
}
func (s *ListenRecordService) DeleteListenRecord(ctx context.Context, req *pb.ListenRecordDeleteReq) (*pb.ListenRecordDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ListenRecordDeleteReply{Result: err == nil}, err
}
func (s *ListenRecordService) BatchDeleteListenRecord(ctx context.Context, req *pb.ListenRecordBatchDeleteReq) (*pb.ListenRecordDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ListenRecordDeleteReply{Result: err == nil && num > 0}, err
}
