package models

//发往消息队列的请求结构体
type Request struct {
	StudentId string
	TicketId int
	Flag int
}
