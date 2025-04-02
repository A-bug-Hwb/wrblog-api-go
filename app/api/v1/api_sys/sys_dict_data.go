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
// @Summary  获取字典数据列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictData query model_sys.SelectSysDictData true "sysDictData"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/page [get]
func ApiDictDataPageList(c *gin.Context) {
	sysDictData := &model_sys.SelectSysDictData{}
	if err := c.ShouldBindQuery(&sysDictData); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Suc(service_sys.GetDictDataPageList(sysDictData)))
}

// @Tags  System - 字典管理
// @Summary  获取指定字典数据列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param dictType path string true "dictType"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/list/{dictType} [get]
func ApiDictDataByType(c *gin.Context) {
	dictType := c.Param("dictType")
	if dictType == "" {
		mylog.MyLog.Panic("参数缺失：{dictType}")
	}
	c.JSON(http.StatusOK, result.Ok(service_sys.ApiDictDataByType(dictType)))
}

// @Tags  System - 字典管理
// @Summary  获取字典数据详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param dictCode path int true "dictCode"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/{dictCode} [get]
func ApiDictDataByCode(c *gin.Context) {
	dictCode, err := strconv.Atoi(c.Param("dictCode"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_sys.GetDictDataById(dictCode)))
}

// @Tags  System - 字典管理
// @Summary  添加菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictData body model_sys.SysDictData true "sysDictData"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/add [post]
func ApiAddDictData(c *gin.Context) {
	sysDictData := &model_sys.SysDictDataPo{}
	if err := c.ShouldBindBodyWithJSON(&sysDictData); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	sysDictData.CreateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.AddDictData(sysDictData)))
}

// @Tags  System - 字典管理
// @Summary  修改菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysDictData body model_sys.SysDictData true "sysDictData"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/edit [put]
func ApiEditDictData(c *gin.Context) {
	sysDictData := &model_sys.SysDictDataPo{}
	if err := c.ShouldBindBodyWithJSON(&sysDictData); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	sysDictData.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.EditDictData(sysDictData)))
}

// @Tags  System - 管理
// @Summary  删除菜单[逻辑删除]
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param dictIds query string true "dictIds"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysDict/data/removes [delete]
func ApiRemoveDictData(c *gin.Context) {
	dictCodes := c.QueryArray("dictCodes")
	c.JSON(http.StatusOK, result.Ok(service_sys.RemoveDictData(dictCodes, token.GetUk(c))))
}
