package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SocialUser holds the schema definition for the SocialUser entity.
type SocialUser struct {
	ent.Schema
}

// Fields of the SocialUser.
func (SocialUser) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").
			Comment(`用户ID`),
		field.String("unionid").Optional().
			Comment(`只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段`),
		field.String("token").Optional().
			Comment(`第三方登录token`),
		field.String("openid").Optional().
			Comment(`用户的标识，对当前公众号唯一`),
		field.String("routineOpenid").Optional().
			Comment(`小程序唯一身份ID`),
		field.String("userName").Optional().
			Comment(`用户名`),
		field.String("nickName").Optional().
			Comment(`用户的昵称`),
		field.Time("birthday").Optional().
			Comment(`生日`),
		field.String("phone").Optional().
			Comment(`手机`),
		field.String("email").Optional().
			Comment(`邮箱`),
		field.String("password").Optional().
			Comment(`密码`),
		field.String("avatar").Optional().
			Comment(`用户头像`),
		field.Int("sex").Optional().
			Comment(`用户的性别，值为1时是男性，值为2时是女性，值为0时是未知`),
		field.String("region").Optional().
			Comment(`用户所在城市区域`),
		field.String("city").Optional().
			Comment(`用户所在城市`),
		field.String("language").Optional().
			Comment(`用户的语言，简体中文为zh_CN`),
		field.String("province").Optional().
			Comment(`用户所在省份`),
		field.String("country").Optional().
			Comment(`用户所在国家`),
		field.String("signature").Optional().
			Comment(`个性签名`),
		field.String("remark").Optional().
			Comment(`公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注`),
		field.Int("groupid").Optional().
			Comment(`用户所在的分组ID（兼容旧的用户分组接口）`),
		field.String("tagidList").Optional().
			Comment(`用户被打上的标签ID列表`),
		field.Int("subscribe").Optional().
			Comment(`用户是否订阅该公众号标识`),
		field.Int("subscribeTime").Optional().
			Comment(`关注公众号时间`),
		field.String("sessionKey").Optional().
			Comment(`小程序用户会话密匙`),
		field.String("userType").Optional().
			Comment(`用户类型`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SocialUser.
func (SocialUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", TaskLog.Type).Comment("任务日志"),
		edge.To("listenRecords", ListenRecord.Type).Comment("收听阅读记录"),
		edge.To("ads", AdChangeLog.Type).Comment("广告变化列表"),
		edge.To("bookshelves", NovelBookshelf.Type).Comment("我的书架"),
		edge.To("autoBuyNovels", NovelAutoBuy.Type).Comment("自动购买书籍"),
		edge.To("comments", NovelComment.Type).Comment("我的评论"),
		edge.To("msgs", UserMsg.Type).Comment("用户消息"),
		edge.To("orders", PayOrder.Type).Comment("订单列表"),
		edge.To("vips", VipUser.Type).Comment("会员列表"),
		edge.To("balances", AmBalance.Type).Comment("账本列表"),
		edge.To("assetLogs", AssetChangeLog.Type).Comment("资金变化列表"),
		edge.To("buyChapterRecords", NovelBuyChapterRecord.Type).Comment("章节购买记录"),
		edge.To("buyNovelRecords", NovelBuyRecord.Type).Comment("整本购买记录"),
		edge.From("channel", AdChannel.Type).Ref("users").Comment("注册渠道").Unique(),
	}
}
