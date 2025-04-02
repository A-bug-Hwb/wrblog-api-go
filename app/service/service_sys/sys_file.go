package service_sys

import (
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
)

func AddFile(sysFile *model_sys.SysFile) *model_sys.SysFile {
	return dao_sys.AddFile(sysFile)
}

func GetFileById(fileId string) *model_sys.SysFile {
	return dao_sys.GetFileById(fileId)
}
