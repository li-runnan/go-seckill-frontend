package models

import "time"

type Student struct { //默认情况对应数据库的表名为：student
	Id int `orm:"pk"` //默认情况对应数据库里student表字段为id
	Studentid string //默认情况对于数据库里student表字段为studentid
	Name string //默认情况对于数据库里student表字段为name
	Password string //默认情况对于数据库里student表字段为password
	Email string //默认情况对于数据库里student表字段为email
	Ipaddress string //默认情况对于数据库里student表字段为ipaddress
	Phone string //默认情况对于数据库里student表字段为phone
	Lastlogintime time.Time
}



