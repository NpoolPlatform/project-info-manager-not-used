package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/NpoolPlatform/project-info-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

type CoinDescription struct {
	ent.Schema
}

func (CoinDescription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (CoinDescription) Fields() []ent.Field {
	const MessageMaxLen = 2048
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.String("title"),
		field.String("message").MaxLen(MessageMaxLen),
		field.String("used_for"),
	}
}

func (CoinDescription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "coin_type_id", "used_for").
			Unique(),
	}
}
