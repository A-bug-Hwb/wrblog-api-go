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

// @Tags  Base - 文章标签管理
// @Summary  获取标签列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLabel query model_base.SelectBaseArticleLabel true "baseArticleLabel"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/page [get]
func ApiArticleLabelPageList(c *gin.Context) {
	baseArticleLabel := &model_base.SelectBaseArticleLabel{}
	if err := c.ShouldBindQuery(&baseArticleLabel); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleLabelPageList(baseArticleLabel)))
}

// @Tags  Base - 文章标签管理
// @Summary  获取当前用户的标签列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLabel query model_base.SelectBaseArticleLabel true "baseArticleLabel"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/myPage [get]
func ApiMyArticleLabelPageList(c *gin.Context) {
	baseArticleLabel := &model_base.SelectBaseArticleLabel{}
	if err := c.ShouldBindQuery(&baseArticleLabel); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleLabel.UserId = token.GetUserId(c)
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleLabelPageList(baseArticleLabel)))
}

// @Tags  Base - 文章标签管理
// @Summary  获取标签详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param labelId path int true "labelId"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/{labelId} [get]
func ApiArticleLabelInfo(c *gin.Context) {
	labelId, err := strconv.Atoi(c.Param("labelId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_base.GetArticleLabelInfo(labelId)))
}

// @Tags  Base - 文章标签管理
// @Summary  添加标签
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLabel body model_base.BaseArticleLabel true "baseArticleLabel"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/add [post]
func ApiAddArticleLabel(c *gin.Context) {
	baseArticleLabel := &model_base.BaseArticleLabelPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleLabel); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	loginUser, _ := token.GetLoginUser(c)
	baseArticleLabel.UserId = loginUser.UserId
	baseArticleLabel.CreateBy = loginUser.Uk
	c.JSON(http.StatusOK, result.Ok(service_base.AddArticleLabel(baseArticleLabel)))
}

// @Tags  Base - 文章标签管理
// @Summary  修改标签
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLabel body model_base.BaseArticleLabel true "baseArticleLabel"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/edit [put]
func ApiEditArticleLabel(c *gin.Context) {
	baseArticleLabel := &model_base.BaseArticleLabelPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleLabel); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleLabel.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_base.EditArticleLabel(baseArticleLabel)))
}

// @Tags  Base - 文章标签管理
// @Summary  删除标签（逻辑删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param labelIds query string true "labelIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/removes [delete]
func ApiRemoveArticleLabel(c *gin.Context) {
	labelIds := c.QueryArray("labelIds")
	c.JSON(http.StatusOK, result.Ok(service_base.RemoveArticleLabel(labelIds, token.GetUk(c))))
}

// @Tags  Base - 文章标签管理
// @Summary  删除标签（实体删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param labelIds query string true "labelIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLabel/deletes [delete]
func ApiDeleteArticleLabel(c *gin.Context) {
	labelIds := c.QueryArray("labelIds")
	c.JSON(http.StatusOK, result.Ok(service_base.DeletedArticleLabel(labelIds)))
}
