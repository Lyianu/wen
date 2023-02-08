package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/gin-gonic/gin"
)

func GetSite(c *gin.Context) {
	c.JSON(http.StatusOK, models.Wen)
}
