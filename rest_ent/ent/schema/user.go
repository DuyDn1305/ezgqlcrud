package schema

import (
	"context"
	"restent/main/myhook"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"golang.org/x/crypto/bcrypt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
// Email string! // pk
// Password string!
// Name string!
// Pref [string!]!
// CreatedAt: time.Time!
// UpdatedAt: time.Time!
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),
		field.String("name").Default("Unknown"),
		field.Strings("pref").Optional(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UUIDMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blogs", Blog.Type),
		edge.To("comments", Comment.Type),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		myhook.On(func(ctx context.Context, m ent.Mutation) {
			if pass, ok := m.Field("password"); ok {
				// ctx.Value()
				// v := pass.(string)
				// fmt.Println(pass, v)
				if new_pass, err := bcrypt.GenerateFromPassword([]byte(pass.(string)), 10); err == nil {
					m.SetField("password", string(new_pass))
					// fmt.Println(string(new_pass))
					// fmt.Println("bcrypt success")
				}
			}
		}, ent.OpCreate|ent.OpUpdate),
	}
}
