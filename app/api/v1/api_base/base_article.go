package api_base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/model/model_base"
	"wrblog-api-go/app/service/service_base"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
)

// @Tags  Base - 文章管理
// @Summary  获取文章列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticle query model_base.SelectBaseArticle true "baseArticle"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/page [get]
func ApiArticlePageList(c *gin.Context) {
	baseArticle := &model_base.SelectBaseArticle{}
	if err := c.ShouldBindQuery(&baseArticle); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticlePageList(baseArticle)))
}

// @Tags  Base - 文章管理
// @Summary  获取当前用户的文章列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticle query model_base.SelectBaseArticle true "baseArticle"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/myPage [get]
func ApiMyArticlePageList(c *gin.Context) {
	baseArticle := &model_base.SelectBaseArticle{}
	if err := c.ShouldBindQuery(&baseArticle); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticle.UserId = token.GetUserId(c)
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticlePageList(baseArticle)))
}

// @Tags  Base - 文章管理
// @Summary  获取文章详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param articleId path int true "articleId"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/{articleId} [get]
func ApiArticleInfo(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_base.GetArticleById(articleId)))
}

// @Tags  Base - 文章管理
// @Summary  添加文章
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticle body model_base.BaseArticle true "baseArticle"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/add [post]
func ApiAddArticle(c *gin.Context) {
	baseArticle := &model_base.BaseArticlePo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticle); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	loginUser, _ := token.GetLoginUser(c)
	baseArticle.UserId = loginUser.UserId
	baseArticle.CreateBy = loginUser.Uk
	c.JSON(http.StatusOK, result.Ok(service_base.AddArticle(baseArticle)))
}

// @Tags  Base - 文章管理
// @Summary  修改文章
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticle body model_base.BaseArticle true "baseArticle"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/edit [put]
func ApiEditArticle(c *gin.Context) {
	baseArticle := &model_base.BaseArticlePo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticle); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticle.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_base.EditArticle(baseArticle)))
}

// @Tags  Base - 文章管理
// @Summary  删除文章（逻辑删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param articleIds query []int true "articleIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/removes [delete]
func ApiRemoveArticle(c *gin.Context) {
	articleIds := c.QueryArray("articleIds")
	c.JSON(http.StatusOK, result.Ok(service_base.RemoveArticle(articleIds, token.GetUk(c))))
}

// @Tags  Base - 文章管理
// @Summary  删除文章（实体删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param articleIds query string true "articleIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticle/deletes [delete]
func ApiDeleteArticle(c *gin.Context) {
	articleIds := c.QueryArray("articleIds")
	c.JSON(http.StatusOK, result.Ok(service_base.DeleteArticle(articleIds)))
}
