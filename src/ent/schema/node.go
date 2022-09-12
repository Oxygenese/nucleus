package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("hostname"),
		field.Int("status"),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return nil
}

// Mixin BaseModel
func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

func (Node) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		// 唯一约束索引
		index.Fields("hostname", "id").
			Unique(),
	}
}
