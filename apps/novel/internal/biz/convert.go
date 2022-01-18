package biz

import (
	v1 "hope/api/novel/novel/v1"
	"hope/apps/novel/internal/data/ent"
)

func ReqToData(req *v1.NovelCreateReq) *ent.Novel {
	return &ent.Novel{
		NovelId:       req.NovelId,
		ClassifyId:    req.ClassifyId,
		ClassifyName:  req.ClassifyName,
		AuthorId:      req.AuthorId,
		Title:         req.Title,
		Summary:       req.Summary,
		Author:        req.Author,
		Anchor:        req.Anchor,
		Hits:          req.Hits,
		Keywords:      req.Keywords,
		Source:        req.Source,
		Score:         req.Score,
		Cover:         req.Cover,
		TagIds:        req.TagIds,
		WordNum:       req.WordNum,
		FreeNum:       req.FreeNum,
		OnlineState:   req.OnlineState,
		Price:         req.Price,
		Publish:       req.Publish,
		OriginalPrice: req.OriginalPrice,
		ChapterPrice:  req.ChapterPrice,
		ChapterCount:  req.ChapterCount,
		SignType:      req.SignType,
		SignDate:      req.SignDate.AsTime(),
		LeadingMan:    req.LeadingMan,
		LeadingLady:   req.LeadingLady,
		Remark:        req.Remark,
		MediaKey:      req.MediaKey,
	}
}
