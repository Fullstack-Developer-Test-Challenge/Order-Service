package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewRabbitMQPublisher(url string) *RabbitMQPublisher {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"order_exchange", // exchange name
		"direct",         // exchange type
		true,             // durable
		false,            // auto-delete
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare exchange")

	q, err := ch.QueueDeclare(
		"orders_queue", // queue name
		true,           // durable
		false,          // auto delete
		false,          // exclusive
		false,          // no wait
		nil,            // args
	)
	failOnError(err, "Failed to declare queue")

	err = ch.QueueBind(
		q.Name,           // queue name
		"order.created",  // routing key
		"order_exchange", // exchange name
		false,            // no wait
		nil,              // args
	)
	failOnError(err, "Failed to declare queue")

	log.Println("RabbitMQ publisher successfully connected and exchange declared.")

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
	}
}

func (p *RabbitMQPublisher) PublishOrderCreated(productID, qty int) {
	body, _ := json.Marshal(map[string]interface{}{
		"pattern": "order.created",
		"data": map[string]interface{}{
			"productId": productID,
			"quantity":  qty,
		},
	})

	err := p.channel.Publish(
		"order_exchange",
		"order.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish message")

	log.Printf("Published event 'order.created': {productId: %d, quantity: %d}", productID, qty)
}

func (p *RabbitMQPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
	log.Println("RabbitMQ connection closed.")
}
