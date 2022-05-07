package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

type Description struct {
	ent.Schema
}

func (Description) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Description) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.String("title"),
		field.String("message").MaxLen(2048),
		field.String("used_for"),
	}
}

func (Description) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "used_for").
			Unique(),
	}
}
