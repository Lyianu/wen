package routers

import (
	"github.com/Lyianu/wen/middleware/debug"
	"github.com/Lyianu/wen/middleware/jwt"
	"github.com/Lyianu/wen/pkg/setting"
	"github.com/Lyianu/wen/routers/api"
	v1 "github.com/Lyianu/wen/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	if setting.RunMode == "debug" {
		r.Use(debug.CorsMiddleware())
	}

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/site", v1.GetSite)

		apiv1.GET("/tags", v1.GetTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.GET("/pages/:id", v1.GetPage)
		apiv1.GET("/pages", v1.GetPages)
	}
	apiv1private := apiv1.Group("/")
	apiv1private.Use(jwt.JWT())
	{
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1.POST("/pages", v1.AddPage)
		apiv1.PUT("/pages", v1.EditPage)
		apiv1.DELETE("/pages/:id", v1.DeletePage)
	}

	return r
}
