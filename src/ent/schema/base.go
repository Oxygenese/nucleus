package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Base holds the schema definition for the BaseModel entity.
type Base struct {
	ent.Schema
}

// Fields of the Base.
func (Base) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Unique().
			Immutable().
			Comment("primary key"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("creation time of item"),

		field.Time("updated_at").
			UpdateDefault(time.Now).
			Optional().
			Nillable().
			Comment("update time of item"),
	}
}

// Edges of the BaseModel.
func (Base) Edges() []ent.Edge {
	return nil
}
