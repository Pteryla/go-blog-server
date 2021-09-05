package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog-server/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()
	collection := v1.NewCollection()
	user := v1.NewUser()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/collections", collection.Create)
		apiv1.DELETE("/collections/:id", collection.Delete)
		apiv1.PUT("/collections/:id", collection.Update)
		apiv1.PATCH("/collections/:id/state", collection.Update)
		apiv1.GET("/collections", collection.List)

		apiv1.POST("/users", user.Create)
		apiv1.DELETE("/users/:id", user.Delete)
		apiv1.PUT("/users/:id", user.Update)
		apiv1.PATCH("/users/:id/state", user.Update)
		apiv1.GET("/users", user.List)
		apiv1.GET("/users/:id", user.Get)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}
	return r
}
