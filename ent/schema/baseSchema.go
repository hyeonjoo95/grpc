package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Comment("생성 일자"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("업데이트 일자"),
	}
}
