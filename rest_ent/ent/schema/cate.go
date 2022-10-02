package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cate holds the schema definition for the Cate entity.
type Cate struct {
	ent.Schema
}

// Fields of the Cate.
func (Cate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Cate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UUIDMixin{},
	}
}

// Edges of the Cate.
func (Cate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blogs", Blog.Type),
	}
}
