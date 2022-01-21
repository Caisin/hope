package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/socialuser/v1"
)

type SocialUserService struct {
	pb.UnimplementedSocialUserServer
	uc  *biz.SocialUserUseCase
	log *log.Helper
}

func NewSocialUserService(uc *biz.SocialUserUseCase, logger log.Logger) *SocialUserService {
	return &SocialUserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SocialUserService) GetPageSocialUser(ctx context.Context, req *pb.SocialUserPageReq) (*pb.SocialUserPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSocialUser")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SocialUserReply, 0)
	for i := range datas {
		items = append(items, convert.SocialUserData2Reply(datas[i]))
	}
	reply := &pb.SocialUserPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SocialUserService) GetSocialUser(ctx context.Context, req *pb.SocialUserReq) (*pb.SocialUserReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSocialUser")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SocialUserData2Reply(daya), err
}
func (s *SocialUserService) UpdateSocialUser(ctx context.Context, req *pb.SocialUserUpdateReq) (*pb.SocialUserUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSocialUser")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SocialUserData2UpdateReply(data), err
}
func (s *SocialUserService) CreateSocialUser(ctx context.Context, req *pb.SocialUserCreateReq) (*pb.SocialUserCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSocialUser")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SocialUserData2CreateReply(data), err
}
func (s *SocialUserService) DeleteSocialUser(ctx context.Context, req *pb.SocialUserDeleteReq) (*pb.SocialUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSocialUser")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SocialUserDeleteReply{Result: err == nil}, err
}
func (s *SocialUserService) BatchDeleteSocialUser(ctx context.Context, req *pb.SocialUserBatchDeleteReq) (*pb.SocialUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSocialUser")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SocialUserDeleteReply{Result: err == nil && num > 0}, err
}
