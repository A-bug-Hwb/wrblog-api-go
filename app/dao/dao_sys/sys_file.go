package dao_sys

import (
	"fmt"
	"time"
	utils2 "wrblog-api-go/app/common/utils"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/pkg/mylog"
	"wrblog-api-go/pkg/mysql"
	"wrblog-api-go/pkg/utils"
)

func AddFile(sysFile *model_sys.SysFile) *model_sys.SysFile {
	db := mysql.Db("wr-file").Model(&model_sys.SysFile{})
	sysFile.FileId = utils2.GetUUIDString()
	sysFile.CreateTime = utils.Time(time.Now())
	err := db.Create(&sysFile).Error
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return sysFile
}

func GetFileById(fileId string) (sysFile *model_sys.SysFile) {
	db := mysql.Db("wr-file").Model(&model_sys.SysFile{})
	err := db.Where("file_id = ?", fileId).First(&sysFile).Error
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据库操作失败：%s", err))
	}
	return
}
