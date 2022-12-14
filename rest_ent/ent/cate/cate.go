// Code generated by ent, DO NOT EDIT.

package cate

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cate type in the database.
	Label = "cate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeBlogs holds the string denoting the blogs edge name in mutations.
	EdgeBlogs = "blogs"
	// Table holds the table name of the cate in the database.
	Table = "cates"
	// BlogsTable is the table that holds the blogs relation/edge. The primary key declared below.
	BlogsTable = "cate_blogs"
	// BlogsInverseTable is the table name for the Blog entity.
	// It exists in this package in order to avoid circular dependency with the "blog" package.
	BlogsInverseTable = "blogs"
)

// Columns holds all SQL columns for cate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

var (
	// BlogsPrimaryKey and BlogsColumn2 are the table columns denoting the
	// primary key for the blogs relation (M2M).
	BlogsPrimaryKey = []string{"cate_id", "blog_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
