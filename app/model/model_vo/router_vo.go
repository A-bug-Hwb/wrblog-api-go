package model_vo

type RouterVo struct {
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Hidden     bool        `json:"hidden"`
	Redirect   string      `json:"redirect"`
	Component  string      `json:"component"`
	Query      string      `json:"query"`
	AlwaysShow bool        `json:"alwaysShow"`
	Meta       *MetaVo     `json:"meta"`
	Children   []*RouterVo `json:"children"`
}

type MetaVo struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Link  string `json:"link"`
}
