package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"sync"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://1:1@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"QueueTCM", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context, q amqp.Queue, ch *amqp.Channel) {
			defer wg.Done()
			send(i, ctx, q, ch)
		}(i, ctx, q, ch)
	}
	wg.Wait()

}
func send(i int, ctx context.Context, q amqp.Queue, ch *amqp.Channel) {

	body := "Hello World! " + strconv.Itoa(i)
	dateUser := make(map[string]interface{})
	dateUser["email"] = "gh56786@rambler.ru"
	dateUser["theme"] = "Fadn"
	//time.Sleep(5 * time.Second)
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers:     dateUser,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
