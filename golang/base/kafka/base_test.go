package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestAbc(t *testing.T) {
	t.Log(getAddr())
}

func TestBase(t *testing.T) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"192.168.10.44:9092"}, conf)
	if err != nil {
		t.Fatal(err)
	}

	message := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.StringEncoder("Hello, Kafka2!"),
	}
	_, _, err = producer.SendMessage(message)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBase2(t *testing.T) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	consumer, err := sarama.NewConsumer([]string{"192.168.10.44:9092"}, conf)
	if err != nil {
		t.Fatal(err)
	}

	partitionConsumer, err := consumer.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		t.Fatal(err)
	}

	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message: Topic = %s, Partition = %d, Offset = %d, Value = %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Value))
	}
}
