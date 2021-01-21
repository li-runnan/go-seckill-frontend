package models

import "github.com/astaxie/beego/orm"

type StudentTicket struct { //默认情况对应数据库的表名为：student_ticket
	Id int
	Studentid string
	Ticketid int
	Flag int //是否抢到了票 1为抢到了 0为没抢到 -1为还没参与秒杀
}

type SecKillResult struct {
	TicketName string //门票名称
	TicketNum int //门票数量
	Flag int //是否抢到了
}

//查询某一学生的抢票结果
func GetSecKillResult(studentId string) []SecKillResult {
	var res []SecKillResult
	var studentTicket []StudentTicket
	o := orm.NewOrm()
	o.QueryTable("student_ticket").Filter("studentid",studentId).All(&studentTicket)
	for i := 0; i<len(studentTicket);i++ {
		ticketId := studentTicket[i].Ticketid
		ticketInfo := GetTicketsInfoById(ticketId)
		ticketName := ticketInfo.Name
		ticketNum := ticketInfo.Num
		flag := studentTicket[i].Flag
		result := SecKillResult{TicketName: ticketName,TicketNum: ticketNum,Flag: flag}
		res = append(res,result)
	}
	return res
}

//将抢票结果写入数据库
func InsertSecKillResult(studentId string,ticketId int,flag int) {
	result := StudentTicket{Studentid: studentId,Ticketid: ticketId,Flag: flag}
	o := orm.NewOrm()
	o.Insert(&result)
}
