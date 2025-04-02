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

// @Tags  Base - 文章库管理
// @Summary  获取库列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLibrary query model_base.SelectBaseArticleLibrary true "baseArticleLibrary"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/page [get]
func ApiArticleLibraryPageList(c *gin.Context) {
	baseArticleLibrary := &model_base.SelectBaseArticleLibrary{}
	if err := c.ShouldBindQuery(&baseArticleLibrary); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleLibraryPageList(baseArticleLibrary)))
}

// @Tags  Base - 文章库管理
// @Summary  获取当前用户的库列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLibrary query model_base.SelectBaseArticleLibrary true "baseArticleLibrary"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/myPage [get]
func ApiMyArticleLibraryPageList(c *gin.Context) {
	baseArticleLibrary := &model_base.SelectBaseArticleLibrary{}
	if err := c.ShouldBindQuery(&baseArticleLibrary); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleLibrary.UserId = token.GetUserId(c)
	c.JSON(http.StatusOK, result.Suc(service_base.GetArticleLibraryPageList(baseArticleLibrary)))
}

// @Tags  Base - 文章库管理
// @Summary  获取库详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param libraryId path int true "libraryId"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/{libraryId} [get]
func ApiArticleLibraryInfo(c *gin.Context) {
	libraryId, err := strconv.Atoi(c.Param("libraryId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_base.GetArticleLibraryInfo(libraryId)))
}

// @Tags  Base - 文章库管理
// @Summary  添加库
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLibrary body model_base.BaseArticleLibrary true "baseArticleLibrary"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/add [post]
func ApiAddArticleLibrary(c *gin.Context) {
	baseArticleLibrary := &model_base.BaseArticleLibraryPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleLibrary); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	loginUser, _ := token.GetLoginUser(c)
	baseArticleLibrary.UserId = loginUser.UserId
	baseArticleLibrary.CreateBy = loginUser.Uk
	c.JSON(http.StatusOK, result.Ok(service_base.AddArticleLibrary(baseArticleLibrary)))
}

// @Tags  Base - 文章库管理
// @Summary  修改库
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param baseArticleLibrary body model_base.BaseArticleLibrary true "baseArticleLibrary"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/edit [put]
func ApiEditArticleLibrary(c *gin.Context) {
	baseArticleLibrary := &model_base.BaseArticleLibraryPo{}
	if err := c.ShouldBindBodyWithJSON(&baseArticleLibrary); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	baseArticleLibrary.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_base.EditArticleLibrary(baseArticleLibrary)))
}

// @Tags  Base - 文章库管理
// @Summary  删除库（逻辑删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param libraryIds query string true "libraryIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/removes [delete]
func ApiRemoveArticleLibrary(c *gin.Context) {
	libraryIds := c.QueryArray("libraryIds")
	c.JSON(http.StatusOK, result.Ok(service_base.RemoveArticleLibrary(libraryIds, token.GetUk(c))))
}

// @Tags  Base - 文章库管理
// @Summary  删除库（实体删除）
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param libraryIds query string true "libraryIds"
// @Success 200 {object} result.Result "OK"
// @Router /base/baseArticleLibrary/deletes [delete]
func ApiDeleteArticleLibrary(c *gin.Context) {
	libraryIds := c.QueryArray("libraryIds")
	c.JSON(http.StatusOK, result.Ok(service_base.DeletedArticleLibrary(libraryIds)))
}
