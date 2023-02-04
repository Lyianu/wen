package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/pkg/setting"
	"github.com/Lyianu/wen/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("Name must not be null")
	valid.MaxSize(name, 100, "name").Message("Name must not exceed 100 characters")
	valid.Required(createdBy, "created_by").Message("CreatedBy must not be null")
	valid.MaxSize(createdBy, 100, "created_by").Message("CreatedBy must not exceed 100 characters")
	valid.Range(state, 0, 1, "state").Message("Status must be 0 or 1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}