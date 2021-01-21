package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"myseckill/models"
	"myseckill/util"
	"strconv"
	"time"
)

type SecKillController struct {
	BaseController
}

//Ret ...
type Ret struct {
	Code int    `json:"code,int"`
	Data string `json:"data"`
}

func (s *SecKillController) Get() {
	TicketId := 2
	StudentId := s.GetSession("studentNo").(string)
	StudentTicket := models.StudentTicket{Studentid: StudentId,Ticketid: TicketId}
	o := orm.NewOrm()
	err := o.Read(&StudentTicket,"Studentid","Ticketid")
	if err == orm.ErrNoRows {
		flag := -1
		s.Data["flag"] = flag
	} else {
		flag := StudentTicket.Flag
		s.Data["flag"] = flag
	}
	ticket := models.GetTicketsInfoById(TicketId)
	startTime := ticket.Starttime
	s.Data["startTime"] = startTime
	nowTime := time.Now()
	s.Data["nowTime"] = nowTime
	s.TplName = "seckill.html"
}

//点击抢购
//获取参数发送至消息队列

func (s *SecKillController) Post() {
	//获取门票id
	TicketId := 2
	//从redis中取出门票数量,即库存
	TicketNum := util.GetTicketNumFromRedis(TicketId)
	stock, _ := strconv.ParseInt(TicketNum,10,64)
	if stock <= 0 {
		//库存不够 秒杀失败
		studentId := s.GetSession("studentNo").(string)
		ticketId := TicketId
		flag := 0 //没抢到
		//封装订单请求
		request := models.Request{StudentId: studentId,TicketId: ticketId,Flag: flag}
		jsons, err := json.MarshalIndent(request,"","\t")
		if err != nil{
			panic(err)
		}
		ret := new(Ret)
		ret.Code = 200
		ret.Data = "抱歉～票已抢光,秒杀失败!"
		s.Ctx.Request.Header.Set("Content-Type","application/json; charset=utf-8")
		retJSON, _ := json.Marshal(ret)
		s.Ctx.ResponseWriter.Write(retJSON)

		//将订单请求发送至消息队列
		util.SendMessage(jsons)

		//s.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('抱歉～票已抢光，秒杀失败');window.location.href='/seckill';</script>"))
	} else {
		//预减库存
		util.MyDecr(TicketId)
		//获取订单请求参数
		studentId := s.GetSession("studentNo").(string)
		ticketId := TicketId
		flag := 1
		//封装订单请求信息
		request := models.Request{StudentId: studentId,TicketId: ticketId,Flag: flag}
		jsons, err := json.MarshalIndent(request,"","\t")
		if err != nil{
			panic(err)
		}
		ret := new(Ret)
		ret.Code = 200
		ret.Data = "恭喜您!秒杀成功!"
		s.Ctx.Request.Header.Set("Content-Type","application/json; charset=utf-8")
		retJSON, _ := json.Marshal(ret)
		s.Ctx.ResponseWriter.Write(retJSON)
		//将订单请求发送至消息队列
		util.SendMessage(jsons)
	}
}
