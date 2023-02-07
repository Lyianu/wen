package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetPages(c *gin.Context) {

}

func GetPage(c *gin.Context) {

}

func AddPage(c *gin.Context) {
	var p models.Page
	err := c.BindJSON(&p)
	code := e.INVALID_PARAMS
	if err != nil {
		util.BadRequest(c, code)
		return
	}
	valid := validation.Validation{}
	valid.MaxSize(p.Title, 100, "title").Message("Title must be shorter than 100")
	valid.MaxSize(p.Content, 65535, "content").Message("Content max length exceed")
	valid.Required(p.Title, "title").Message("Title must not be null")
	valid.Required(p.CreatedBy, "created_by").Message("Created_by must not be null")
	valid.Range(p.State, 0, 1, "state").Message("State must be 0 or 1")

	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["title"] = p.Title
		data["desc"] = p.Desc
		data["content"] = p.Content
		data["created_by"] = p.CreatedBy
		data["modified_by"] = p.CreatedBy
		data["state"] = p.State

		models.AddPage(data)
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditPage(c *gin.Context) {

}

func DeletePage(c *gin.Context) {

}
