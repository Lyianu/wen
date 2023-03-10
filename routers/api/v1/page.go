package v1

import (
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/unknwon/com"
)

func GetPages(c *gin.Context) {
	// data stores the final data given to user
	data := make(map[string]interface{})
	// maps is the constraints given by user
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("State must be 0 or 1")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetPages(maps)
	} else {
		util.LogValidationErrors(valid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetPage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID must be positive")
	code := e.INVALID_PARAMS
	var page models.Page
	if !valid.HasErrors() {
		if models.ExistPageByID(id) {
			page = models.GetPage(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_PAGE
		}
	} else {
		util.LogValidationErrors(valid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": page,
	})
}

func GetPageHTML(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID must be positive")
	code := e.INVALID_PARAMS
	var page models.Page
	if !valid.HasErrors() {
		if models.ExistPageByID(id) {
			page = models.GetPage(id)
			page.Content = string(markdown.ToHTML([]byte(page.Content), nil, nil))
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_PAGE
		}
	} else {
		util.LogValidationErrors(valid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": page,
	})
}

func AddPage(c *gin.Context) {
	var p models.Page
	err := c.BindJSON(&p)
	code := e.INVALID_PARAMS
	if err != nil {
		util.BadRequest(c, code)
		return
	}
	p.State = 0
	p.CreatedBy = c.GetString("username")
	valid := validation.Validation{}
	valid.MaxSize(p.Title, 100, "title").Message("Title must be shorter than 100")
	valid.MaxSize(p.Content, 65535, "content").Message("Content max length exceed")
	valid.Required(p.Title, "title").Message("Title must not be null")
	valid.Required(p.CreatedBy, "created_by").Message("Created_by must not be null")
	//valid.Range(state, 0, 1, "state").Message("State must be 0 or 1")

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
	} else {
		util.BadRequest(c, http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditPage(c *gin.Context) {
	var data struct {
		Title      string `json:"title"`
		Desc       string `json:"desc"`
		Content    string `json:"content"`
		ModifiedBy string `json:"modified_by"`
		State      string `json:"state"`
	}
	c.BindJSON(&data)
	var state = 0
	if data.State != "" {
		state = com.StrTo(data.State).MustInt()
	}

	if data.ModifiedBy == "" {
		data.ModifiedBy = c.GetString("username")
	}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Range(state, 0, 1, "state").Message("State must be 0 or 1")
	valid.Min(id, 1, "id").Message("ID must be positive")
	valid.MaxSize(data.Title, 100, "title").Message("Title length must not exceed 100")
	valid.MaxSize(data.Desc, 255, "desc").Message("Description length must not exceed 255")
	valid.MaxSize(data.Content, 65535, "content").Message("Content length must not exceed 65535")
	valid.Required(data.ModifiedBy, "modified_by").Message("Modified_by must not be null")
	valid.MaxSize(data.ModifiedBy, 100, "modified_by").Message("Modified_by length must not exceed 100")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistPageByID(id) {
			d := make(map[string]interface{})
			if data.Title != "" {
				d["title"] = data.Title
			}
			if data.Desc != "" {
				d["desc"] = data.Desc
			}
			if data.Content != "" {
				d["content"] = data.Content
			}
			d["modified_by"] = data.ModifiedBy

			models.EditPage(id, d)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_PAGE
		}
	} else {
		util.LogValidationErrors(valid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func DeletePage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID must be positive")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistPageByID(id) {
			code = e.SUCCESS
			models.DeletePage(id)
		} else {
			code = e.ERROR_NOT_EXIST_PAGE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
