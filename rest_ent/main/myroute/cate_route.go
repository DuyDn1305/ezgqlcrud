package myroute

import (
	"context"
	"restent/ent"

	"github.com/labstack/echo/v4"
)

type cateRoute _Route

var Cate = cateRoute{}

func (route *cateRoute) getAll(c echo.Context) error {
	cate, err := route.c.Cate.Query().All(route.ctx)
	if err != nil {
		return echo.NewHTTPError(500)
	}
	return c.JSON(200, cate)
}


func (route *cateRoute) Init(router *echo.Group, c *ent.Client, ctx context.Context) {
	route.c = c
	route.ctx = ctx
	router.GET("/all", route.getAll)
}

