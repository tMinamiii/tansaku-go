package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/tMinamiii/tansaku-go/graph"
)

func route(e *echo.Echo) {
	gqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
	e.POST("/query", func(c echo.Context) error {
		gqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	playgroundHandler := playground.Handler("GraphQL", "/query")
	e.GET("/", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

}
