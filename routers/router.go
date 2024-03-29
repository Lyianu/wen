package routers

import (
	"github.com/Lyianu/wen/middleware/debug"
	"github.com/Lyianu/wen/middleware/jwt"
	"github.com/Lyianu/wen/middleware/redis"
	"github.com/Lyianu/wen/pkg/setting"
	v1 "github.com/Lyianu/wen/routers/api/v1"
	"github.com/Lyianu/wen/util"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	if setting.RunMode == "debug" {
		r.Use(debug.CorsMiddleware())
	}

	r.Use(static.Serve("/", static.LocalFile("frontend", true)))

	if setting.RunMode == "release" {
		r.NoRoute(func(c *gin.Context) {
			c.File("./frontend/index.html")
		})
	} else if setting.RunMode == "debug" {
		r.NoRoute(util.ProxyReact)
	}

	apiv1 := r.Group("/api/v1")
	if setting.RedisHost != "" {
		apiv1.Use(redis.Redis())
	}
	{
		apiv1.GET("/site", v1.GetSite)
		apiv1.POST("/site", v1.AddSite)

		apiv1.GET("/tags", v1.GetTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticleHTML)

		apiv1.GET("/pages/:id", v1.GetPageHTML)
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

		apiv1private.GET("/articles/:id/md", v1.GetArticle)
		apiv1private.POST("/articles", v1.AddArticle)
		apiv1private.PUT("/articles/:id", v1.EditArticle)
		apiv1private.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1private.GET("/pages/:id/md", v1.GetPage)
		apiv1private.POST("/pages", v1.AddPage)
		apiv1private.PUT("/pages/:id", v1.EditPage)
		apiv1private.DELETE("/pages/:id", v1.DeletePage)

		apiv1private.PUT("/site", v1.EditSite)
	}

	return r
}
