package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("请设置昵称").StructTag(`json:"name"`).Comment("用户名"),
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("头像"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("用户的手机号"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.String("email").Default("").StructTag(`json:"email"'`).Comment("邮箱"),
	}
}

// Indexes of User
func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("creations", Creation.Type),
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户表"),
	}
}
