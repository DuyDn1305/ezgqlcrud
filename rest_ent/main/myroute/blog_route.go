package myroute

import (
	"context"
	"fmt"
	"restent/ent"
	"restent/ent/blog"
	"restent/ent/cate"
	"restent/ent/comment"
	"restent/ent/user"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type blogRoute _Route

var Blog = blogRoute{}

func (route *blogRoute) getAll(c echo.Context) error {
	blogs, err := route.c.Blog.Query().
		WithTags(func (q *ent.CateQuery) {
			q.Select(cate.FieldName)
		}).
		WithAuthor(func (q * ent.UserQuery) {
			q.Select(user.FieldEmail).Select(user.FieldCreatedAt)
		}).
		All(route.ctx)
	if err != nil {
		return echo.NewHTTPError(500)
	}
	return c.JSON(200, blogs)
}

func (route *blogRoute) getBlog(c echo.Context) error {

	qUser := func(q * ent.UserQuery) {
		q.Select(user.FieldName)
	}

	if id, err := uuid.Parse(c.Param("id")); err == nil {
		b, err := route.c.Blog.Query().
			Where(blog.ID(id)).
			WithTags(func (q *ent.CateQuery) {
				q.Select(cate.FieldName)
			}).
			WithAuthor(qUser).
			WithComments(func (q * ent.CommentQuery) {
				q.WithReplies(func (q *ent.CommentQuery) {
					q.WithWriter(qUser)
				}).WithWriter(qUser).Where(comment.Not(comment.HasReplyTo())).All(route.ctx)
			}).Only(route.ctx)
		if err == nil {
			return c.JSON(200, b)
		}
		return badRequest(c, "Blog not found")
	}
	return badRequest(c, "UUID invalid")
}

func (route *blogRoute) createBlog(c echo.Context) error {
	claims := c.Get("claims").(Dict)
	user_id, _ := uuid.Parse(claims["user_id"].(string))

	type BlogInput struct {
		ent.Blog
		List_cate []uuid.UUID `json:"list_cate"`
	};

	data := BlogInput{}
	if c.Bind(&data) == nil  {
		b := &data.Blog
		route.c.Blog.
			Create().
			SetTitle(b.Title).
			SetContent(b.Content).
			SetThumbnail(b.Thumbnail).
			SetAuthorID(user_id).
			AddTagIDs(data.List_cate...).
			Save(route.ctx)
		return c.JSON(200, Dict{
			"message": "Blog created",
		})
	}
	return badRequest(c, "Cannot created")
}

func (route *blogRoute) updateBlog(c echo.Context) error {

	var id uuid.UUID
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return badRequest(c, "UUID invalid")
	}

	type BlogInput struct {
		ent.Blog
		List_cate []uuid.UUID `json:"list_cate"`
	};

	data := BlogInput{}
	if c.Bind(&data) == nil  {
		b := &data.Blog
		route.c.Blog.
			UpdateOneID(id).
			SetTitle(b.Title).
			SetContent(b.Content).
			SetThumbnail(b.Thumbnail).
			ClearTags().AddTagIDs(data.List_cate...).
			Save(route.ctx)
		return c.JSON(200, Dict{
			"message": "Blog updated",
		})
	}
	return badRequest(c, "Cannot updated")
}

func (route *blogRoute) deleteBlog(c echo.Context) error {
	if id, err := uuid.Parse(c.Param("id")); err == nil {
		if err := route.c.Blog.DeleteOneID(id).Exec(route.ctx); err == nil {
			return c.JSON(200, Dict{
				"message": "Delete "+id.String(),
			})
		}
		return badRequest(c, "Blog not found")
	}
	return badRequest(c, "UUID invalid")
}


func (route *blogRoute) createCmt(c echo.Context) error {
	
	var id uuid.UUID // blog id, comment id
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return badRequest(c, "id invalid")
	}

	claims := c.Get("claims").(Dict)
	user_id, _ := uuid.Parse(claims["user_id"].(string))

	var cmt ent.Comment

	if c.Bind(&cmt) == nil {
		
		builder := route.c.Comment.Create().SetContent(cmt.Content).SetBelongtoID(id).SetWriterID(user_id)
		replyToParam := c.Param("replyTo")
		if replyToParam != "" {
			replyTo, err := uuid.Parse(c.Param("replyTo"))
			if err != nil {
				return badRequest(c, "Reply to invalid id")
			}
			// check replyTo hasReplyTo
			reply_target, err := route.c.Comment.Query().Where(comment.ID(replyTo)).QueryReplyTo().OnlyID(route.ctx)
			fmt.Println(reply_target, err)
			if err != nil { //replyTo is root 
				builder.SetReplyToID(replyTo)
			} else {
				builder.SetReplyToID(reply_target)
			}
		}
		if _, err := builder.Save(route.ctx); err != nil {
			return severError(c, err.Error())
		}
	}
	return c.JSON(200, Dict{"message": "Comment created"})

}


func (route *blogRoute) Init(router *echo.Group, c *ent.Client, ctx context.Context) {
	route.c = c
	route.ctx = ctx
	// router.Use(authMiddleware)

	router.GET("/", route.getAll)
	router.GET("/:id", route.getBlog)
	router.POST("/", route.createBlog, authMiddleware)
	router.POST("/:id/comment", route.createCmt, authMiddleware)
	router.POST("/:id/comment/reply/:replyTo", route.createCmt, authMiddleware)
	router.PUT("/:id", route.updateBlog, authMiddleware)
	router.DELETE("/:id", route.deleteBlog, authMiddleware)
	// router.POST("/create", route.createBlog)

}