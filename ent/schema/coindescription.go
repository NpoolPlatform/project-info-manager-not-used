package schema

import "entgo.io/ent"

// CoinDescription holds the schema definition for the CoinDescription entity.
type CoinDescription struct {
	ent.Schema
}

// Fields of the CoinDescription.
func (CoinDescription) Fields() []ent.Field {
	return nil
}

// Edges of the CoinDescription.
func (CoinDescription) Edges() []ent.Edge {
	return nil
}
