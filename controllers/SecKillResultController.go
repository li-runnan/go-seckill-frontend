package controllers

import (
	"fmt"
	"myseckill/models"
)

type SecKillResultController struct {
	BaseController
}

func (s *SecKillResultController) Get(){
	var result []models.SecKillResult
	result = models.GetSecKillResult(s.GetSession("studentNo").(string))
	fmt.Println(result)
	s.Data["result"] = result
	s.TplName = "result.html"
}