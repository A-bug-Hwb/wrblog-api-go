package api_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wrblog-api-go/app/service"
	"wrblog-api-go/pkg/result"
)

// @Tags  Auth - 认证授权
// @Summary  获取图片验证码
// @Accept json
// @Produce json
// @Success 200 {object} result.Result "OK"
// @Router /auth/getImgCode [get]
func ApiImgCode(c *gin.Context) {
	data := service.GetImgCode()
	c.JSON(http.StatusOK, result.Ok(data))
}

// @Tags  Auth - 认证授权
// @Summary  获取登录加密公钥
// @Accept json
// @Produce json
// @Success 200 {object} result.Result "OK"
// @Router /auth/getPublicKey [get]
func ApiPublicKey(c *gin.Context) {
	data := service.GetPublicKey()
	c.JSON(http.StatusOK, result.Ok(data))
}
