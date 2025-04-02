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

// @Tags  System - 菜单管理
// @Summary  获取前端需要的路由
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /sys/getRouters [get]
func ApiRouters(c *gin.Context) {
	c.JSON(http.StatusOK, result.Ok(service_sys.GetRouters(token.GetUserId(c))))
}

// @Tags  System - 菜单管理
// @Summary  获取菜单管理树表
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysMenu query model_sys.SelectSysMenu true "sysMenu"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/getTree [get]
func ApiSysMenuTree(c *gin.Context) {
	sysMenu := &model_sys.SelectSysMenu{}
	if err := c.ShouldBindQuery(&sysMenu); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_sys.GetMenuTree(sysMenu)))
}

// @Tags  System - 菜单管理
// @Summary  获取菜单详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param menuId path int true "menuId"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/{menuId} [get]
func ApiSysMenuById(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	c.JSON(http.StatusOK, result.Ok(service_sys.GetMenuById(menuId)))
}

// @Tags  System - 菜单管理
// @Summary  添加菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysMenu body model_sys.SysMenu true "sysMenu"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/add [post]
func ApiAddSysMenu(c *gin.Context) {
	sysMenu := &model_sys.SysMenuPo{}
	if err := c.ShouldBindBodyWithJSON(&sysMenu); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}
	sysMenu.CreateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.AddSysMenu(sysMenu)))
}

// @Tags  System - 菜单管理
// @Summary  修改菜单
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param sysMenu body model_sys.SysMenu true "sysMenu"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/edit [put]
func ApiEditSysMenu(c *gin.Context) {
	sysMenu := &model_sys.SysMenuPo{}
	if err := c.ShouldBindBodyWithJSON(&sysMenu); err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("参数读取失败:%s", err))
	}

	sysMenu.UpdateBy = token.GetUk(c)
	c.JSON(http.StatusOK, result.Ok(service_sys.EditSysMenu(sysMenu)))
}

// @Tags  System - 菜单管理
// @Summary  删除菜单[逻辑删除]
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param menuId query string true "menuId"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/remove [delete]
func ApiRemoveSysMenu(c *gin.Context) {
	menuId := c.Query("menuId")
	row, err := service_sys.RemoveSysMenu(menuId, token.GetUk(c))
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Ok(row))
}

// @Tags  System - 菜单管理
// @Summary  删除菜单[物理删除]
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param menuId query string true "menuId"
// @Success 200 {object} result.Result "OK"
// @Router /sys/sysMenu/delete [delete]
func ApiDeleteSysMenu(c *gin.Context) {
	menuId := c.Query("menuId")
	bo, err := service_sys.DeletedSysMenu(menuId)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.Ok(bo))
}
