package model_sys

import (
	"wrblog-api-go/app/model"
)

type SysFile struct {
	FileId       string `json:"fileId" form:"fileId" gorm:"primaryKey;autoIncrement"`
	FileType     string `json:"fileType" form:"fileType"`
	UploadType   string `json:"uploadType" form:"uploadType"`
	OldName      string `json:"oldName" form:"oldName"`
	NewName      string `json:"newName" form:"newName"`
	FileSuffix   string `json:"fileSuffix" form:"fileSuffix"`
	RelativePath string `json:"relativePath" form:"relativePath"`
	AbsolutePath string `json:"absolutePath" form:"absolutePath"`
	model.BaseEntityVo
}

func (SysFile) TableName() string {
	return "sys_file"
}
