package v1

import (
	"log"
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/Lyianu/wen/pkg/setting"
	"github.com/Lyianu/wen/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/unknwon/com"
)

// GetArticle returns desired article(in markdown) specified by id
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID must be positive")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticleHTML returns desired article(in HTML) specified by id
func GetArticleHTML(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID must be positive")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			a := models.GetArticle(id)
			a.Content = string(markdown.ToHTML([]byte(a.Content), nil, nil))
			data = a

			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle adds an article to the db using user-provided data
// it first parse the JSON object in the request, valid it and call
// models.AddArticle to finish
func AddArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	article.CreatedBy = c.GetString("username")

	valid := validation.Validation{}
	// might change to tag names in the future
	for _, v := range article.TagID {
		valid.Min(v, 1, "ID must be positive")
	}
	valid.Required(article.Title, "title").Message("Title must not be null")
	//valid.Required(article.Desc, "desc").Message("Description must be not null")
	//valid.Required(article.Content, "content").Message("Content must not be null")
	valid.Required(article.CreatedBy, "created_by").Message("Created_by must not be null")
	valid.Range(article.State, 0, 1, "state").Message("State must be 0 or 1")

	//TODO: assign default tag when user don't specify
	article.Tags = models.FindTags(article.TagID...)

	code := e.INVALID_PARAMS
	if err == nil && !valid.HasErrors() {
		data := make(map[string]interface{})
		data["tag_id"] = article.TagID
		data["title"] = article.Title
		data["desc"] = article.Desc
		data["content"] = article.Content
		data["created_by"] = article.CreatedBy
		data["state"] = article.State

		models.AddArticle(data)
		code = e.SUCCESS

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}
	data := make(map[string]interface{})
	data["id"] = models.GetArticleTotal(map[string]interface{}{})

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles return the requested articles to user, using
// user-provided constraints and get articles with models.GetArticles
func GetArticles(c *gin.Context) {
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

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("Tag_id must be positive")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		util.LogValidationErrors(valid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	var data struct {
		TagID      []int  `json:"tag_id"`
		Title      string `json:"title"`
		Desc       string `json:"desc"`
		Content    string `json:"content"`
		ModifiedBy string `json:"modified_by"`
		State      string `json:"state"`
	}
	c.BindJSON(&data)
	if data.ModifiedBy == "" {
		data.ModifiedBy = c.GetString("username")
	}
	id := com.StrTo(c.Param("id")).MustInt()

	var state int = 0
	if data.State != "" {
		state = com.StrTo(data.State).MustInt()
	}
	valid.Range(state, 0, 1, "state").Message("State must be 0 or 1")
	valid.Min(id, 1, "id").Message("ID must be positive")
	valid.MaxSize(data.Title, 100, "title").Message("Title length must not exceed 100")
	valid.MaxSize(data.Desc, 255, "desc").Message("Description length must not exceed 255")
	valid.MaxSize(data.Content, 65535, "content").Message("Content length must not exceed 65535")
	valid.Required(data.ModifiedBy, "modified_by").Message("Modified_by must not be null")
	valid.MaxSize(data.ModifiedBy, 100, "modified_by").Message("Modified_by length must not exceed 100")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagsByID(data.TagID...) {
				editData := make(map[string]interface{})
				if len(data.TagID) > 0 {
					editData["tag_id"] = data.TagID
				}
				if data.Title != "" {
					editData["title"] = data.Title
				}
				if data.Desc != "" {
					editData["desc"] = data.Desc
				}
				if data.Content != "" {
					editData["content"] = data.Content
				}

				editData["modified_by"] = data.ModifiedBy

				models.EditArticle(id, editData)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
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

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID must be positive")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			code = e.SUCCESS
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
