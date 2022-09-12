package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseModel holds the schema definition for the BaseModel entity.
type BaseModel struct {
	ent.Schema
}

// Fields of the BaseModel.
func (BaseModel) Fields() []ent.Field {
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
func (BaseModel) Edges() []ent.Edge {
	return nil
}
