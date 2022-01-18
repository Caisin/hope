package biz

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/novel/v1"
	"hope/apps/novel/internal/data/ent"
)

func ReqToData(v *v1.NovelCreateReq) *ent.Novel {
	return &ent.Novel{
		NovelId:       v.NovelId,
		ClassifyId:    v.ClassifyId,
		ClassifyName:  v.ClassifyName,
		AuthorId:      v.AuthorId,
		Title:         v.Title,
		Summary:       v.Summary,
		Author:        v.Author,
		Anchor:        v.Anchor,
		Hits:          v.Hits,
		Keywords:      v.Keywords,
		Source:        v.Source,
		Score:         v.Score,
		Cover:         v.Cover,
		TagIds:        v.TagIds,
		WordNum:       v.WordNum,
		FreeNum:       v.FreeNum,
		OnlineState:   v.OnlineState,
		Price:         v.Price,
		Publish:       v.Publish,
		OriginalPrice: v.OriginalPrice,
		ChapterPrice:  v.ChapterPrice,
		ChapterCount:  v.ChapterCount,
		SignType:      v.SignType,
		SignDate:      v.SignDate.AsTime(),
		LeadingMan:    v.LeadingMan,
		LeadingLady:   v.LeadingLady,
		Remark:        v.Remark,
		MediaKey:      v.MediaKey,
	}
}
func Data2Req(v *ent.Novel) *v1.NovelCreateReq {
	return &v1.NovelCreateReq{
		NovelId:       v.NovelId,
		ClassifyId:    v.ClassifyId,
		ClassifyName:  v.ClassifyName,
		AuthorId:      v.AuthorId,
		Title:         v.Title,
		Summary:       v.Summary,
		Author:        v.Author,
		Anchor:        v.Anchor,
		Hits:          v.Hits,
		Keywords:      v.Keywords,
		Source:        v.Source,
		Score:         v.Score,
		Cover:         v.Cover,
		TagIds:        v.TagIds,
		WordNum:       v.WordNum,
		FreeNum:       v.FreeNum,
		OnlineState:   v.OnlineState,
		Price:         v.Price,
		Publish:       v.Publish,
		OriginalPrice: v.OriginalPrice,
		ChapterPrice:  v.ChapterPrice,
		ChapterCount:  v.ChapterCount,
		SignType:      v.SignType,
		SignDate:      timestamppb.New(v.SignDate),
		LeadingMan:    v.LeadingMan,
		LeadingLady:   v.LeadingLady,
		Remark:        v.Remark,
		MediaKey:      v.MediaKey,
	}
}
