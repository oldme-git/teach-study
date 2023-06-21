package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

const (
	topic = "kafka one"
	group = "kafka one_1"
	host  = "192.168.10.43:9092,192.168.10.43:9093,192.168.10.43:9094"
)

func getAddr() []string {
	return strings.Split(host, ",")
}

func getConf() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Net.DialTimeout = time.Second
	return conf
}

// 创建一个client
func newClient() (sarama.Client, error) {
	return sarama.NewClient(getAddr(), getConf())
}

// 创建一个生产者连接
func newProducer() (sarama.SyncProducer, error) {
	return sarama.NewSyncProducer(getAddr(), getConf())
}

// 创建一个消费者连接
func newConsumer() (sarama.Consumer, error) {
	return sarama.NewConsumer(getAddr(), getConf())
}

// 发送一条消息
func sendMsg(topic string, value string) error {
	producer, err := newProducer()
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
