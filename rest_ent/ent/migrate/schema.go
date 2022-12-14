// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "user_blogs", Type: field.TypeUUID, Nullable: true},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blogs_users_blogs",
				Columns:    []*schema.Column{BlogsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CatesColumns holds the columns for the "cates" table.
	CatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// CatesTable holds the schema information for the "cates" table.
	CatesTable = &schema.Table{
		Name:       "cates",
		Columns:    CatesColumns,
		PrimaryKey: []*schema.Column{CatesColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "content", Type: field.TypeString},
		{Name: "blog_comments", Type: field.TypeUUID, Nullable: true},
		{Name: "comment_replies", Type: field.TypeUUID, Nullable: true},
		{Name: "user_comments", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_blogs_comments",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{BlogsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_comments_replies",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Default: "Unknown"},
		{Name: "pref", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// CateBlogsColumns holds the columns for the "cate_blogs" table.
	CateBlogsColumns = []*schema.Column{
		{Name: "cate_id", Type: field.TypeUUID},
		{Name: "blog_id", Type: field.TypeUUID},
	}
	// CateBlogsTable holds the schema information for the "cate_blogs" table.
	CateBlogsTable = &schema.Table{
		Name:       "cate_blogs",
		Columns:    CateBlogsColumns,
		PrimaryKey: []*schema.Column{CateBlogsColumns[0], CateBlogsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cate_blogs_cate_id",
				Columns:    []*schema.Column{CateBlogsColumns[0]},
				RefColumns: []*schema.Column{CatesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cate_blogs_blog_id",
				Columns:    []*schema.Column{CateBlogsColumns[1]},
				RefColumns: []*schema.Column{BlogsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlogsTable,
		CatesTable,
		CommentsTable,
		UsersTable,
		CateBlogsTable,
	}
)

func init() {
	BlogsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = BlogsTable
	CommentsTable.ForeignKeys[1].RefTable = CommentsTable
	CommentsTable.ForeignKeys[2].RefTable = UsersTable
	CateBlogsTable.ForeignKeys[0].RefTable = CatesTable
	CateBlogsTable.ForeignKeys[1].RefTable = BlogsTable
}
