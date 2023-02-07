package api

import (
	"log"
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(c *gin.Context) {
	a := auth{}
	data := make(map[string]interface{})
	err := c.BindJSON(&a)
	code := e.INVALID_PARAMS

	if err != nil {
		util.BadRequest(c, code)
		return
	}

	valid := validation.Validation{}
	if ok, _ := valid.Valid(&a); ok {
		isExist := models.CheckAuth(a.Username, a.Password)
		if isExist {
			token, err := util.GenerateToken(a.Username, a.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
