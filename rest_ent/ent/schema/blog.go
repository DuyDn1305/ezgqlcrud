package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// id UUID
// 	title string!
// 	content string!
// 	thumbnail string!
// 	created_at time.Time!
// 	updated_at time.Time!
// 	cate Cate
// 	author User!	

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.String("thumbnail"),
	}
}

func (Blog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UUIDMixin{},
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Cate.Type).Ref("blogs"),
		edge.From("author", User.Type).Ref("blogs").Unique(),
		edge.To("comments", Comment.Type),
	}
}
