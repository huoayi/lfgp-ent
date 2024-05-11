// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CreationsColumns holds the columns for the "creations" table.
	CreationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "creation_type", Type: field.TypeEnum, Comment: "创作类型", Enums: []string{"txt2voice", "txt2img", "voice2voice"}, Default: "txt2img"},
		{Name: "parameter", Type: field.TypeString, Comment: "参数", Default: ""},
		{Name: "url", Type: field.TypeString, Comment: "作品链接", Default: ""},
		{Name: "user_id", Type: field.TypeInt64, Comment: "用户 id", Default: 0},
	}
	// CreationsTable holds the schema information for the "creations" table.
	CreationsTable = &schema.Table{
		Name:       "creations",
		Columns:    CreationsColumns,
		PrimaryKey: []*schema.Column{CreationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "creations_users_creations",
				Columns:    []*schema.Column{CreationsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Comment: "19 位雪花 ID"},
		{Name: "created_by", Type: field.TypeInt64, Comment: "创建者 ID", Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Comment: "更新者 ID", Default: 0},
		{Name: "created_at", Type: field.TypeTime, Comment: "创建时刻，带时区"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "更新时刻，带时区"},
		{Name: "deleted_at", Type: field.TypeTime, Comment: "软删除时刻，带时区"},
		{Name: "name", Type: field.TypeString, Comment: "用户名", Default: "请设置昵称"},
		{Name: "nick_name", Type: field.TypeString, Comment: "用户昵称", Default: "请设置昵称"},
		{Name: "jpg_url", Type: field.TypeString, Comment: "头像", Default: ""},
		{Name: "phone", Type: field.TypeString, Comment: "用户的手机号", Default: ""},
		{Name: "password", Type: field.TypeString, Comment: "密码", Default: ""},
		{Name: "email", Type: field.TypeString, Comment: "邮箱", Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户表",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CreationsTable,
		UsersTable,
	}
)

func init() {
	CreationsTable.ForeignKeys[0].RefTable = UsersTable
	CreationsTable.Annotation = &entsql.Annotation{}
	UsersTable.Annotation = &entsql.Annotation{}
}
