package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
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

}
