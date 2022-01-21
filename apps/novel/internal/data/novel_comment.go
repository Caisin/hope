package data
import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelcomment/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelcomment"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"hope/pkg/pagin"
	"time"
)

type novelCommentRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelCommentRepo .
func NewNovelCommentRepo(data *Data, logger log.Logger) biz.NovelCommentRepo {
	return &novelCommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelComment 创建
func (r *novelCommentRepo) CreateNovelComment(ctx context.Context, req *v1.NovelCommentCreateReq) (*ent.NovelComment, error) {
	now := time.Now()
	return r.data.db.NovelComment.Create().
    SetNovelId(req.NovelId).
    SetUserId(req.UserId).
    SetAvatar(req.Avatar).
    SetUserName(req.UserName).
    SetRepUserId(req.RepUserId).
    SetRepUserName(req.RepUserName).
    SetContent(req.Content).
    SetScore(req.Score).
    SetPId(req.PId).
    SetIsTop(req.IsTop).
    SetState(novelcomment.State(req.State)).
    SetIsHighlight(req.IsHighlight).
    SetIsHot(req.IsHot).
    SetRemark(req.Remark).
	SetCreatedAt(now).
	SetUpdatedAt(now).
	Save(ctx)

}

// DeleteNovelComment 删除
func (r *novelCommentRepo) DeleteNovelComment(ctx context.Context, req *v1.NovelCommentDeleteReq) error {
	return r.data.db.NovelComment.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelComment 批量删除
func (r *novelCommentRepo) BatchDeleteNovelComment(ctx context.Context, req *v1.NovelCommentBatchDeleteReq) (int, error) {
	return r.data.db.NovelComment.Delete().Where(novelcomment.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelComment 更新
func (r *novelCommentRepo) UpdateNovelComment(ctx context.Context, req *v1.NovelCommentUpdateReq) (*ent.NovelComment, error) {
	return r.data.db.NovelComment.UpdateOne(convert.NovelCommentUpdateReq2Data(req)).Save(ctx)
}

// GetNovelComment 根据Id查询
func (r *novelCommentRepo) GetNovelComment(ctx context.Context, req *v1.NovelCommentReq) (*ent.NovelComment, error) {
	return r.data.db.NovelComment.Get(ctx, req.Id)
}

// PageNovelComment 分页查询
func (r *novelCommentRepo) PageNovelComment(ctx context.Context, req *v1.NovelCommentPageReq) ([]*ent.NovelComment, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin=&pagin.Pagination{}
	}
	query := r.data.db.NovelComment.
		Query().
		Where(
			//查询条件构造
			r.genCondition(req.Param)...,
		)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	req.GetPagin().Total = int64(count)
	if count == 0 {
		return nil, nil
	}
	query.Limit(int(p.GetPage())).
		Offset(int(p.GetOffSet()))
	if p.NeedOrder() {
		if p.IsDesc() {
			query.Order(ent.Desc(p.GetField()))
		} else {
			query.Order(ent.Asc(p.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *novelCommentRepo) genCondition(req *v1.NovelCommentReq) []predicate.NovelComment {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelComment, 0)
	if req.Id > 0 {
		list = append(list, novelcomment.ID(req.Id))
	}
	if req.NovelId > 0 {
		list = append(list, novelcomment.NovelId(req.NovelId))
	}
	if req.UserId > 0 {
		list = append(list, novelcomment.UserId(req.UserId))
	}
	if str.IsBlank(req.Avatar) {
		list = append(list, novelcomment.AvatarContains(req.Avatar))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, novelcomment.UserNameContains(req.UserName))
	}
	if req.RepUserId > 0 {
		list = append(list, novelcomment.RepUserId(req.RepUserId))
	}
	if str.IsBlank(req.RepUserName) {
		list = append(list, novelcomment.RepUserNameContains(req.RepUserName))
	}
	if str.IsBlank(req.Content) {
		list = append(list, novelcomment.ContentContains(req.Content))
	}
	if req.Score > 0 {
		list = append(list, novelcomment.Score(req.Score))
	}
	if req.PId > 0 {
		list = append(list, novelcomment.PId(req.PId))
	}
	state := novelcomment.State(req.State)
	if novelcomment.StateValidator(state)==nil {
		list = append(list, novelcomment.StateEQ(state))
	}
if str.IsBlank(req.Remark) {
		list = append(list, novelcomment.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelcomment.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelcomment.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelcomment.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelcomment.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelcomment.TenantId(req.TenantId))
	}
	
	return list
}
