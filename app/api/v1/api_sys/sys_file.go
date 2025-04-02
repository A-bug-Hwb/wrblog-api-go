package api_sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/result"
	"wrblog-api-go/pkg/utils"
)

// @Tags  Common - 公共接口
// @Summary  上传文件
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} result.Result "OK"
// @Router /file/upload [post]
func ApiUpload(c *gin.Context) {
	var res *result.Result
	file, err := c.FormFile("file")
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("文件上传失败：%s", err))
	} else {
		oldFileName := file.Filename
		suffix := filepath.Ext(oldFileName)
		newFileName := fmt.Sprintf("%s-%d%s", oldFileName[:len(oldFileName)-len(suffix)], time.Now().Unix(), suffix)
		newFilePath := fmt.Sprintf("/%s/%s", utils.GetNowDay(), newFileName)
		err = c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", config.Conf.ConfigInfo.Profile, newFilePath))
		if err != nil {
			mylog.MyLog.Panic("文件上传失败：%s", err)
			res = result.Fail(fmt.Sprintf("文件上传失败：%s", err))
		} else {
			var sysFile *model_sys.SysFile
			sysFile = dao_sys.AddFile(&model_sys.SysFile{
				FileType:     suffix,
				OldName:      oldFileName,
				NewName:      newFileName,
				FileSuffix:   suffix,
				RelativePath: config.Conf.ConfigInfo.FilePrefix + newFilePath,
				AbsolutePath: config.Conf.ConfigInfo.Profile + newFilePath,
			})
			if err != nil {
				mylog.MyLog.Panic("文件上传失败：%s", err)
				res = result.Fail(fmt.Sprintf("文件上传失败：%s", err))
			}
			res = result.Ok(sysFile)
		}
	}
	c.JSON(http.StatusOK, res)
}
