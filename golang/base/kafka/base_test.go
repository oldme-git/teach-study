package kafka

import (
	"context"
	"fmt"
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
		return
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

	time.Sleep(3 * time.Second)
}
