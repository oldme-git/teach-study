package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
	"time"
)

// 测试消费
func TestConsumer(t *testing.T) {
	var partition int32 = 3

	kafka, err := newClient()
	if err != nil {
		t.Fatal(err)
	}
	consumer, err := kafka.newConsumer()
	if err != nil {
		t.Fatal(err)
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, 0)

	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for msg := range partitionConsumer.Messages() {
			fmt.Printf("接收topic=%s, partition=%d, offset=%d, key=%s, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
		}
	}()

	time.Sleep(10 * time.Second)

	// 关闭消费者
	if err := consumer.Close(); err != nil {
		log.Fatal("Failed to close consumer: ", err)
	}
}

// 测试生产
func TestProducer(t *testing.T) {
	kafka, err := newClient()
	if err != nil {
		t.Fatal(err)
	}
	producer, err := kafka.newProducer()

	if err != nil {
		t.Fatal(err)
	}

	defer producer.Close()
	msg := creMsg(topic, 3, "my key2", "I'm value")
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}

	log.Printf("Message sent successfully: topic=%s, partition=%d, offset=%d\n", topic, partition, offset)
}

// 测试消费者组
func TestConsumerGrp(t *testing.T) {
	ctx := context.Background()

	kafka, err := newClient()
	if err != nil {
		t.Fatal(err)
	}
	consumerGroup, err := kafka.newConsumerGroup(group)
	if err != nil {
		return
	}
	defer consumerGroup.Close()
	consumer := &Consumer{}

	go func() {
		for {
			err := consumerGroup.Consume(ctx, []string{topic}, consumer)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	for {

	}
}

// 异步发送消息
func TestAsyncProducer(t *testing.T) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	client, err := sarama.NewClient(getAddr(), conf)
	if err != nil {
		t.Fatal(err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		t.Fatal(err)
	}
	defer producer.AsyncClose()

	go func() {
		for {
			select {
			case msg := <-producer.Successes():
				fmt.Printf("发送成功, topic=%s, partition=%d, offset=%d\n", msg.Topic, msg.Partition, msg.Offset)
			case err = <-producer.Errors():
				fmt.Printf("发送消息失败: %s\n", err.Error())
			}
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("async value"),
	}

	producer.Input() <- msg

	time.Sleep(3 * time.Second)
}
