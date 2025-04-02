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

// @Tags  Base - 文章组管理
// @Summary  获取组列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleGroup query model_base.SelectBaseArticleGroup true "baseArticleGroup"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/{page} [get]
func ApiArticleGroupPageList(c *gin.Context) {
	baseArticleGroup := &model_base.SelectBaseArticleGroup{}
	if err := c.ShouldBindQuery(&baseArticleGroup); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleGroupPageList(baseArticleGroup)))
}

// @Tags  Base - 文章组管理
// @Summary  获取当前用户的组列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleGroup query model_base.SelectBaseArticleGroup true "baseArticleGroup"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/myPage [get]
func ApiMyArticleGroupPageList(c *gin.Context) {
	baseArticleGroup := &model_base.SelectBaseArticleGroup{}
	if err := c.ShouldBindQuery(&baseArticleGroup); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleGroup.UserId = token.GetUserId(c)
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleGroupPageList(baseArticleGroup)))
}

// @Tags  Base - 文章组管理
// @Summary  获取组详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param groupId path int true "groupId"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/:groupId [get]
func ApiArticleGroupInfo(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_base.GetArticleGroupInfo(groupId)))
}

// @Tags  Base - 文章组管理
// @Summary  添加组
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleGroup body model_base.BaseArticleGroup true "baseArticleGroup"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/add [post]
func ApiAddArticleGroup(c *gin.Context) {
	baseArticleGroup := &model_base.BaseArticleGroupPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleGroup); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	loginUser, _ := token.GetLoginUser(c)
	baseArticleGroup.UserId = loginUser.UserId
	baseArticleGroup.CreateBy = loginUser.Uk
	c.JSON(http.StatusOK, result.Ok(service_base.AddArticleGroup(baseArticleGroup)))
}

// @Tags  Base - 文章组管理
// @Summary  修改组
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleGroup body model_base.BaseArticleGroup true "baseArticleGroup"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/edit [put]
func ApiEditArticleGroup(c *gin.Context) {
	baseArticleGroup := &model_base.BaseArticleGroupPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleGroup); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleGroup.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_base.EditArticleGroup(baseArticleGroup)))
}

// @Tags  Base - 文章组管理
// @Summary  删除组（逻辑删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param groupIds query string true "groupIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/removes [delete]
func ApiRemoveArticleGroup(c *gin.Context) {
	groupIds := c.QueryArray("groupIds")
	c.JSON(http.StatusOK, result.Ok(service_base.RemoveArticleGroup(groupIds, token.GetUk(c))))
}

// @Tags  Base - 文章组管理
// @Summary  删除组（实体删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param groupIds query string true "groupIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleGroup/deletes [delete]
func ApiDeleteArticleGroup(c *gin.Context) {
	groupIds := c.QueryArray("groupIds")
	c.JSON(http.StatusOK, result.Ok(service_base.DeletedArticleGroup(groupIds)))
}
