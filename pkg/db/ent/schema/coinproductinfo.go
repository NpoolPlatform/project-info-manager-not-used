package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/NpoolPlatform/project-info-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

type CoinProductInfo struct {
	ent.Schema
}

func (CoinProductInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (CoinProductInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.String("product_page"),
	}
}

func (CoinProductInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "coin_type_id").
			Unique(),
	}
}
