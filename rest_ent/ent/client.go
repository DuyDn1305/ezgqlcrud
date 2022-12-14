// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"restent/ent/migrate"

	"restent/ent/blog"
	"restent/ent/cate"
	"restent/ent/comment"
	"restent/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Blog is the client for interacting with the Blog builders.
	Blog *BlogClient
	// Cate is the client for interacting with the Cate builders.
	Cate *CateClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Blog = NewBlogClient(c.config)
	c.Cate = NewCateClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Blog:    NewBlogClient(cfg),
		Cate:    NewCateClient(cfg),
		Comment: NewCommentClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Blog:    NewBlogClient(cfg),
		Cate:    NewCateClient(cfg),
		Comment: NewCommentClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Blog.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Blog.Use(hooks...)
	c.Cate.Use(hooks...)
	c.Comment.Use(hooks...)
	c.User.Use(hooks...)
}

// BlogClient is a client for the Blog schema.
type BlogClient struct {
	config
}

// NewBlogClient returns a client for the Blog from the given config.
func NewBlogClient(c config) *BlogClient {
	return &BlogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blog.Hooks(f(g(h())))`.
func (c *BlogClient) Use(hooks ...Hook) {
	c.hooks.Blog = append(c.hooks.Blog, hooks...)
}

// Create returns a builder for creating a Blog entity.
func (c *BlogClient) Create() *BlogCreate {
	mutation := newBlogMutation(c.config, OpCreate)
	return &BlogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Blog entities.
func (c *BlogClient) CreateBulk(builders ...*BlogCreate) *BlogCreateBulk {
	return &BlogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Blog.
func (c *BlogClient) Update() *BlogUpdate {
	mutation := newBlogMutation(c.config, OpUpdate)
	return &BlogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlogClient) UpdateOne(b *Blog) *BlogUpdateOne {
	mutation := newBlogMutation(c.config, OpUpdateOne, withBlog(b))
	return &BlogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlogClient) UpdateOneID(id uuid.UUID) *BlogUpdateOne {
	mutation := newBlogMutation(c.config, OpUpdateOne, withBlogID(id))
	return &BlogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Blog.
func (c *BlogClient) Delete() *BlogDelete {
	mutation := newBlogMutation(c.config, OpDelete)
	return &BlogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlogClient) DeleteOne(b *Blog) *BlogDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BlogClient) DeleteOneID(id uuid.UUID) *BlogDeleteOne {
	builder := c.Delete().Where(blog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlogDeleteOne{builder}
}

// Query returns a query builder for Blog.
func (c *BlogClient) Query() *BlogQuery {
	return &BlogQuery{
		config: c.config,
	}
}

// Get returns a Blog entity by its id.
func (c *BlogClient) Get(ctx context.Context, id uuid.UUID) (*Blog, error) {
	return c.Query().Where(blog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlogClient) GetX(ctx context.Context, id uuid.UUID) *Blog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a Blog.
func (c *BlogClient) QueryTags(b *Blog) *CateQuery {
	query := &CateQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(blog.Table, blog.FieldID, id),
			sqlgraph.To(cate.Table, cate.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, blog.TagsTable, blog.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Blog.
func (c *BlogClient) QueryAuthor(b *Blog) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(blog.Table, blog.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, blog.AuthorTable, blog.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Blog.
func (c *BlogClient) QueryComments(b *Blog) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(blog.Table, blog.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, blog.CommentsTable, blog.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BlogClient) Hooks() []Hook {
	return c.hooks.Blog
}

// CateClient is a client for the Cate schema.
type CateClient struct {
	config
}

// NewCateClient returns a client for the Cate from the given config.
func NewCateClient(c config) *CateClient {
	return &CateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cate.Hooks(f(g(h())))`.
func (c *CateClient) Use(hooks ...Hook) {
	c.hooks.Cate = append(c.hooks.Cate, hooks...)
}

// Create returns a builder for creating a Cate entity.
func (c *CateClient) Create() *CateCreate {
	mutation := newCateMutation(c.config, OpCreate)
	return &CateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cate entities.
func (c *CateClient) CreateBulk(builders ...*CateCreate) *CateCreateBulk {
	return &CateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cate.
func (c *CateClient) Update() *CateUpdate {
	mutation := newCateMutation(c.config, OpUpdate)
	return &CateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CateClient) UpdateOne(ca *Cate) *CateUpdateOne {
	mutation := newCateMutation(c.config, OpUpdateOne, withCate(ca))
	return &CateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CateClient) UpdateOneID(id uuid.UUID) *CateUpdateOne {
	mutation := newCateMutation(c.config, OpUpdateOne, withCateID(id))
	return &CateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cate.
func (c *CateClient) Delete() *CateDelete {
	mutation := newCateMutation(c.config, OpDelete)
	return &CateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CateClient) DeleteOne(ca *Cate) *CateDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CateClient) DeleteOneID(id uuid.UUID) *CateDeleteOne {
	builder := c.Delete().Where(cate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CateDeleteOne{builder}
}

// Query returns a query builder for Cate.
func (c *CateClient) Query() *CateQuery {
	return &CateQuery{
		config: c.config,
	}
}

// Get returns a Cate entity by its id.
func (c *CateClient) Get(ctx context.Context, id uuid.UUID) (*Cate, error) {
	return c.Query().Where(cate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CateClient) GetX(ctx context.Context, id uuid.UUID) *Cate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlogs queries the blogs edge of a Cate.
func (c *CateClient) QueryBlogs(ca *Cate) *BlogQuery {
	query := &BlogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cate.Table, cate.FieldID, id),
			sqlgraph.To(blog.Table, blog.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, cate.BlogsTable, cate.BlogsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CateClient) Hooks() []Hook {
	return c.hooks.Cate
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a builder for creating a Comment entity.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id uuid.UUID) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CommentClient) DeleteOneID(id uuid.UUID) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id uuid.UUID) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id uuid.UUID) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWriter queries the writer edge of a Comment.
func (c *CommentClient) QueryWriter(co *Comment) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.WriterTable, comment.WriterColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReplyTo queries the reply_to edge of a Comment.
func (c *CommentClient) QueryReplyTo(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.ReplyToTable, comment.ReplyToColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReplies queries the replies edge of a Comment.
func (c *CommentClient) QueryReplies(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comment.RepliesTable, comment.RepliesColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBelongto queries the belongto edge of a Comment.
func (c *CommentClient) QueryBelongto(co *Comment) *BlogQuery {
	query := &BlogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(blog.Table, blog.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.BelongtoTable, comment.BelongtoColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlogs queries the blogs edge of a User.
func (c *UserClient) QueryBlogs(u *User) *BlogQuery {
	query := &BlogQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(blog.Table, blog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BlogsTable, user.BlogsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a User.
func (c *UserClient) QueryComments(u *User) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CommentsTable, user.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
