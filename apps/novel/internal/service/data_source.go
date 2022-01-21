package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/datasource/v1"
)

type DataSourceService struct {
	pb.UnimplementedDataSourceServer
	uc  *biz.DataSourceUseCase
	log *log.Helper
}

func NewDataSourceService(uc *biz.DataSourceUseCase, logger log.Logger) *DataSourceService {
	return &DataSourceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *DataSourceService) GetPageDataSource(ctx context.Context, req *pb.DataSourcePageReq) (*pb.DataSourcePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageDataSource")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.DataSourceReply, 0)
	for i := range datas {
		items = append(items, convert.DataSourceData2Reply(datas[i]))
	}
	reply := &pb.DataSourcePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *DataSourceService) GetDataSource(ctx context.Context, req *pb.DataSourceReq) (*pb.DataSourceReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetDataSource")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.DataSourceData2Reply(daya), err
}
func (s *DataSourceService) UpdateDataSource(ctx context.Context, req *pb.DataSourceUpdateReq) (*pb.DataSourceUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateDataSource")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.DataSourceData2UpdateReply(data), err
}
func (s *DataSourceService) CreateDataSource(ctx context.Context, req *pb.DataSourceCreateReq) (*pb.DataSourceCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateDataSource")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.DataSourceData2CreateReply(data), err
}
func (s *DataSourceService) DeleteDataSource(ctx context.Context, req *pb.DataSourceDeleteReq) (*pb.DataSourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteDataSource")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DataSourceDeleteReply{Result: err == nil}, err
}
func (s *DataSourceService) BatchDeleteDataSource(ctx context.Context, req *pb.DataSourceBatchDeleteReq) (*pb.DataSourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteDataSource")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DataSourceDeleteReply{Result: err == nil && num > 0}, err
}
