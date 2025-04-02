package api_sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/app/service/service_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
)

// @Tags  System - 字典管理
// @Summary  获取字典类型列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictType query model_sys.SelectSysDictType true "sysDictType"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/type/page [get]
func ApiDictTypePageList(c *gin.Context) {
	sysDictType := &model_sys.SelectSysDictType{}
	if err := c.ShouldBindQuery(&sysDictType); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_sys.GetDictTypePageList(sysDictType)))
}

// @Tags  System - 字典管理
// @Summary  获取字典类型详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param dictId path int true "dictId"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/type/{dictId} [get]
func ApiDictTypeById(c *gin.Context) {
	dictId, err := strconv.Atoi(c.Param("dictId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_sys.GetDictTypeById(dictId)))
}

// @Tags  System - 字典管理
// @Summary  添加菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictType body model_sys.SysDictType true "sysDictType"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/type/add [post]
func ApiAddDictType(c *gin.Context) {
	sysDictType := &model_sys.SysDictTypePo{}
	if err := c.ShouldBindBodyWithJSON(&sysDictType); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	sysDictType.CreateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.AddDictType(sysDictType)))
}

// @Tags  System - 字典管理
// @Summary  修改菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictType body model_sys.SysDictType true "sysDictType"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/type/edit [put]
func ApiEditDictType(c *gin.Context) {
	sysDictType := &model_sys.SysDictTypePo{}
	if err := c.ShouldBindBodyWithJSON(&sysDictType); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	sysDictType.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.EditDictType(sysDictType)))
}

// @Tags  System - 管理
// @Summary  删除菜单[逻辑删除]
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param dictIds query string true "dictIds"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/type/removes [delete]
func ApiRemoveDictType(c *gin.Context) {
	dictIds := c.QueryArray("dictIds")
	c.JSON(http.StatusOK, result.Ok(service_sys.RemoveDictType(dictIds, token.GetUk(c))))
}

// @Tags  System - 管理
// @Summary  删除缓存
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/removeCache [delete]
func ApiRemoveCacheAll(c *gin.Context) {
	c.JSON(http.StatusOK, result.Ok(service_sys.RemoveCacheAll()))
}
