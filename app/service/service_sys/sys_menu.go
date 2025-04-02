package service_sys

import (
	"errors"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/app/dao/dao_sys"
	"wrblog-api-go/app/model/model_sys"
	"wrblog-api-go/app/model/model_vo"
)

func GetRouters(userId int) (routers []*model_vo.RouterVo) {
	selectSysMenu := &model_sys.SelectSysMenu{}
	if !token.IsAdmin(userId) {
		selectSysMenu.UserId = userId
	}
	sysMenus := GetMenuTree(selectSysMenu)
	routers = buildMenus(sysMenus)
	return
}

func GetMenuTree(selectSysMenu *model_sys.SelectSysMenu) (sysMenus []*model_sys.SysMenuVo) {
	sysMenuAll := dao_sys.GetMenuList(selectSysMenu)
	for _, sysMenu := range sysMenuAll {
		if sysMenu.ParentId == 0 {
			sysMenus = append(sysMenus, sysMenu)
		}
	}
	listToTree(sysMenus, sysMenuAll)
	if len(sysMenus) == 0 {
		sysMenus = sysMenuAll
	}
	return
}

func GetMenuById(menuId int) *model_sys.SysMenuVo {
	return dao_sys.GetMenuById(menuId)
}

func AddSysMenu(sysMenu *model_sys.SysMenuPo) int {
	_, menuId := dao_sys.SaveSysMenu(sysMenu)
	return menuId
}

func EditSysMenu(sysMenu *model_sys.SysMenuPo) int64 {
	row, _ := dao_sys.SaveSysMenu(sysMenu)
	return row
}

// RemoveSysMenu 逻辑删除
func RemoveSysMenu(menuId string, uk string) (row int64, err error) {
	if dao_sys.IsChildren(menuId) {
		err = errors.New("存在子菜单，不允许删除！")
	} else {
		row = dao_sys.RemoveSysMenu(menuId, uk)
	}
	return
}

// DeletedSysMenu 删除
func DeletedSysMenu(menuId string) (row int64, err error) {
	if !dao_sys.IsChildren(menuId) {
		err = errors.New("存在子菜单，不允许删除！")
	} else {
		row = dao_sys.DeletedSysMenu(menuId)
	}
	return
}

func listToTree(sysMenus []*model_sys.SysMenuVo, sysMenuAll []*model_sys.SysMenuVo) {
	for _, sysMenu := range sysMenus {
		var childrens []*model_sys.SysMenuVo
		for _, sysMenuIn := range sysMenuAll {
			if sysMenuIn.ParentId == sysMenu.MenuId {
				childrens = append(childrens, sysMenuIn)
			}
		}
		listToTree(childrens, sysMenuAll)
		sysMenu.Children = childrens
	}
}

func buildMenus(sysMenuList []*model_sys.SysMenuVo) (routers []*model_vo.RouterVo) {
	for _, sysMenu := range sysMenuList {
		router := &model_vo.RouterVo{
			Hidden:    sysMenu.Visible == "1",
			Name:      getRouterName(sysMenu),
			Component: getComponent(sysMenu),
			Path:      getRouterPath(sysMenu),
			Query:     sysMenu.Query,
			Meta: &model_vo.MetaVo{
				Title: sysMenu.MenuName,
				Icon:  sysMenu.Icon,
				Link:  sysMenu.Path,
			},
		}
		cMenus := sysMenu.Children
		if cMenus != nil && len(cMenus) > 0 && sysMenu.MenuType == "M" {
			router.AlwaysShow = true
			router.Redirect = "noRedirect"
			router.Children = buildMenus(sysMenu.Children)
		} else if isMenuFrame(sysMenu) {
			router.Meta = nil
			var childrens []*model_vo.RouterVo
			children := &model_vo.RouterVo{
				Path:      sysMenu.Path,
				Component: sysMenu.Component,
				Name:      getRouteName(sysMenu.MenuName, sysMenu.Path),
				Meta: &model_vo.MetaVo{
					Title: sysMenu.MenuName,
					Icon:  sysMenu.Icon,
					Link:  sysMenu.Path,
				},
			}
			childrens = append(childrens, children)
			router.Children = childrens
		} else if sysMenu.ParentId == 0 && isInnerLink(sysMenu) {
			router.Meta = &model_vo.MetaVo{
				Title: sysMenu.MenuName,
				Icon:  sysMenu.Icon,
			}
			router.Path = "/"
			var childrens []*model_vo.RouterVo
			children := &model_vo.RouterVo{
				Path:      sysMenu.Path,
				Component: "InnerLink",
				Name:      getRouteName(sysMenu.MenuName, sysMenu.Path),
			}
			childrens = append(childrens, children)
			router.Children = childrens
		}
		routers = append(routers, router)
	}
	return
}

func getRouterName(sysMenu *model_sys.SysMenuVo) string {
	if isMenuFrame(sysMenu) {
		return ""
	}
	return getRouteName(sysMenu.MenuName, sysMenu.Path)
}

func getRouteName(name string, path string) string {
	if name == "" {
		return path
	}
	return name
}

func getRouterPath(sysMenu *model_sys.SysMenuVo) (routerPath string) {
	routerPath = sysMenu.Path
	// 非外链并且是一级目录（类型为目录）
	if sysMenu.ParentId == 0 && sysMenu.MenuType == "M" && sysMenu.IsFrame == "1" {
		routerPath = "/" + sysMenu.Path
	} else if isMenuFrame(sysMenu) {
		routerPath = "/"
	}
	return
}

func getComponent(sysMenu *model_sys.SysMenuVo) (component string) {
	component = "Layout"
	if sysMenu.Component != "" && !isMenuFrame(sysMenu) {
		component = sysMenu.Component
	} else if sysMenu.Component == "" && sysMenu.ParentId == 0 && isInnerLink(sysMenu) {
		component = "InnerLink"
	} else if sysMenu.Component == "" && isParentView(sysMenu) {
		component = "ParentView"
	}
	return
}

func isMenuFrame(sysMenu *model_sys.SysMenuVo) bool {
	return sysMenu.ParentId == 0 && sysMenu.MenuType == "C" && sysMenu.IsFrame == "1"
}

func isInnerLink(sysMenu *model_sys.SysMenuVo) bool {
	return sysMenu.IsFrame == "1" && len(sysMenu.Path) > 5 && (sysMenu.Path[0:4] == "http" || sysMenu.Path[0:5] == "https")
}

func isParentView(sysMenu *model_sys.SysMenuVo) bool {
	return sysMenu.ParentId != 0 && sysMenu.MenuType == "M"
}
