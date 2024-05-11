package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/lfgp-ent/pkg/enum"
)

type Creation struct {
	ent.Schema
}

// Fields of the Creation.
func (Creation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("creation_type").GoType(enum.Txt2img).StructTag(`json:"creation_type"`).Default(string(enum.Txt2img)).Comment("创作类型"),
		field.String("parameter").Default("").StructTag(`json:"parameter"`).Comment("参数"),
		field.String("url").Default("").StructTag(`json:"url"`).Comment("作品链接"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("用户 id"),
	}
}

// Indexes of Creation
func (Creation) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Creation) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("creations").Field("user_id").Unique().Required(),
	}
}

// Mixin of User
func (Creation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
