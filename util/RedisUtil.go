package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"myseckill/models"
	"time"
)

var redisdb *redis.Client
var ctx = context.Background()

//初始化连接
func InitClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "123.56.230.53:6379",
		Password: "",
		DB: 0, // use default db
	})

	pong, err := redisdb.Ping(ctx).Result()
	log.Println(pong,err)
	if err != nil {
		return err
	}
	return nil
}

//将距离秒杀时间不足24小时的门票id和门票数量写入redis
func SetTicketsInfoToRedisIn24hours() {
	ticket := models.GetTicketsInfoIn24hours()
	fmt.Println(ticket)
	for i := 0; i< len(ticket);i++{
		key := ticket[i].Id
		num := ticket[i].Num
		err := redisdb.Set(ctx, string(key),num,time.Hour * 24).Err()
		if err != nil {
			panic(err)
		}

		val, err := redisdb.Get(ctx, string(key)).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key:",key,"value:",val)
	}
}

//为便于演示，在系统初始化时将开放秒杀的门票id和门票数量写入redis
func SetTicketsInfoToRedis() {
	ticket := models.GetTicketsInfo()
	fmt.Println(ticket)
	for i := 0; i< len(ticket);i++{
		key := ticket[i].Id
		num := ticket[i].Num
		err := redisdb.Set(ctx, string(key),num,0).Err()
		if err != nil {
			panic(err)
		}

		val, err := redisdb.Get(ctx, string(key)).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key:",key,"value:",val)
	}
}

func GetTicketNumFromRedis(id int) string {
	val, err := redisdb.Get(ctx,string(id)).Result()
	if err != nil{
		panic(err)
	}
	return val
}

func MyDecr(id int) int {
	val, err := redisdb.Decr(ctx,string(id)).Result()
	if err != nil{
		panic(err)
	}
	return int(val)
}

