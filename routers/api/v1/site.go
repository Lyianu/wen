package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/util"
	"github.com/gin-gonic/gin"
)

func GetSite(c *gin.Context) {
	c.JSON(http.StatusOK, models.Wen)
}

// AddSite setup the website for the first time, if it has been setup before,
// shows a Bad Request message
func AddSite(c *gin.Context) {
	var site struct {
		models.Site
		Username string `json:"user"`
		Password string `json:"pass"`
	}
	if models.Wen.Name != "" || c.BindJSON(&site) != nil {
		c.String(http.StatusBadRequest, "Bad request")
	}

	models.AddSite(site.Name, site.SiteImageURL, site.BgTitle, site.Desc)
	models.AddAuth(site.Username, site.Password)
}

func EditSite(c *gin.Context) {
	var site models.Site
	err := c.BindJSON(&site)
	if err != nil {
		util.BadRequest(c, http.StatusBadRequest)
	}
	if len(site.Name) == 0 {
		util.BadRequest(c, http.StatusBadRequest)
	}

	models.EditSite(map[string]interface{}{
		"name":     site.Name,
		"bg_title": site.BgTitle,
		"desc":     site.Desc,
		"img_url":  site.SiteImageURL,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": make(map[string]interface{}),
	})
}
