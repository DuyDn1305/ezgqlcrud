package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}


func (Comment) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
		UUIDMixin{},
    }
}

// type Comment {
// 	id UUID
// 	content string!
// 	created_at time.Time!
// 	updated_at time.Time!
// 	blog Blog
// 	parentComment Comment
// 	writer User
// }
// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("writer", User.Type).Ref("comments").Unique(),
		edge.To("replies", Comment.Type).From("reply_to").Unique(),
		edge.From("belongto", Blog.Type).Ref("comments").Unique(),

	}
}
