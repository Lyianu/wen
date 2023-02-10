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
// shows a 404 page
func AddSite(c *gin.Context) {

}
