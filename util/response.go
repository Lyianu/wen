package util

import (
	"log"
	"net/http"

	"github.com/Lyianu/wen/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BadRequest sends a response with StatusBadRequest
// with Body containing the programmer-defined error code
func BadRequest(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func LogValidationErrors(valid validation.Validation) {
	for _, err := range valid.Errors {
		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	}
}
