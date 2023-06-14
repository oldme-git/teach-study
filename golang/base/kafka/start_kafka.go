package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
)

const (
	topic = "kafka one"
	group = "kafka one_1"
	host  = "192.168.10.44:9092"
)

func getAddr() []string {
	return strings.Split(host, ",")
}

func getConf() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	return conf
}

type ExConsumerGroup struct {
}

func producer() {
	producer, err := sarama.NewSyncProducer(getAddr(), getConf())
	if err != nil {
		panic(err)
	}
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("发送信息"),
	})
	if err != nil {
		panic(err)
	}
}

func consumer() {
	//consumer, err := sarama.NewConsumer()
}
