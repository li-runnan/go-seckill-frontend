package util

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"myseckill/models"
)


//Ret ...
type Ret struct {
	Code int    `json:"code,int"`
	Data string `json:"data"`
}

func failOnError(err error,msg string) {
	if err != nil {
		log.Fatalf("%s:%s",msg,err)
	}
}

func SendMessage(jsons []byte) {
	conn, err:= amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"failed to declare a queue")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: jsons,
		})
	log.Printf("[x] Sent %s",string(jsons))
	failOnError(err,"failed to publish a message")
}

func ReceiveMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			deal(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func deal(jsons []byte) {
	SecKillRequest := models.Request{}
	err := json.Unmarshal(jsons,&SecKillRequest)
	if err != nil{
		panic(err)
	}
	studentId := SecKillRequest.StudentId
	ticketId := SecKillRequest.TicketId
	flag := SecKillRequest.Flag
	models.InsertSecKillResult(studentId,ticketId,flag)
}