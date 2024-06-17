package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Comment("이메일"),
		field.String("password").NotEmpty().Comment("패스워드"),
		field.String("name").NotEmpty().Comment("이름"),
		field.Uint32("age").Comment("나이"),
		field.String("phone_number").NotEmpty().Comment("핸드폰 번호"),
		field.Bool("is_used").Default(true).Comment("현재 사용 여부"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
