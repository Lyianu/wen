package routers

import (
	"github.com/Lyianu/wen/middleware/debug"
	"github.com/Lyianu/wen/middleware/jwt"
	"github.com/Lyianu/wen/pkg/setting"
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

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/site", v1.GetSite)
		apiv1.POST("/site", v1.AddSite)

		apiv1.GET("/tags", v1.GetTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticleHTML)

		apiv1.GET("/pages/:id", v1.GetPage)
		apiv1.GET("/pages", v1.GetPages)

		apiv1.POST("/user", v1.GetAuth)
		apiv1.POST("/user/register", v1.AddAuth)
	}
	apiv1private := apiv1.Group("/")
	apiv1private.Use(jwt.JWT())
	{
		apiv1private.POST("/tags", v1.AddTag)
		apiv1private.PUT("/tags/:id", v1.EditTag)
		apiv1private.DELETE("/tags/:id", v1.DeleteTag)

		apiv1private.POST("/articles", v1.AddArticle)
		apiv1private.PUT("/articles/:id", v1.EditArticle)
		apiv1private.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1private.POST("/pages", v1.AddPage)
		apiv1private.PUT("/pages", v1.EditPage)
		apiv1private.DELETE("/pages/:id", v1.DeletePage)

		apiv1private.PUT("/site", v1.EditSite)
	}

	return r
}
