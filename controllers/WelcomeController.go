package controllers

type WelcomeController struct {
	BaseController
}

func (w *WelcomeController) Get() {
	w.TplName = "welcome.html"
}
