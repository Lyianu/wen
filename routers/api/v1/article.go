package v1

import (
	"log"
	"net/http"

	"github.com/Lyianu/wen/models"
	"github.com/Lyianu/wen/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

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

func AddArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)

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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func GetArticles(c *gin.Context) {

}

func EditArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
