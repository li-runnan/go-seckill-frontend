package controllers

type HomeController struct {
	BaseController
}

// @Title Home
// @Description 进入主页
// @router / [get]
func (h *HomeController) Get(){
	h.TplName = "index.html"
}
