package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

const (
	topic = "kafka_one"
	group = "kafka_one_1"
	host  = "192.168.10.43:9092,192.168.10.43:9093,192.168.10.43:9094"
)

func getAddr() []string {
	return strings.Split(host, ",")
}

func getConf() *sarama.Config {
	conf := sarama.NewConfig()
	// 生产消息后是否需要通知生产者
	// 同步模式会直接返回
	// 异步模式会返回到Successes和Errors通道中
	conf.Producer.Return.Successes = true
	conf.Net.DialTimeout = time.Second
	return conf
}

type Client struct {
	client sarama.Client
}

func newClient() (*Client, error) {
	c := &Client{}
	client, err := sarama.NewClient(getAddr(), getConf())
	if err != nil {
		return nil, err
	}
	c.client = client
	return c, nil
}

// 获取所有topic
func (c *Client) topics() ([]string, error) {
	return c.client.Topics()
}

// 创建一个生产者连接
func (c *Client) newProducer() (sarama.SyncProducer, error) {
	return sarama.NewSyncProducerFromClient(c.client)
}

// 创建一个消费者连接
func (c *Client) newConsumer() (sarama.Consumer, error) {
	return sarama.NewConsumerFromClient(c.client)
}

// 创建一个消费者组
func (c *Client) newConsumerGroup(gid string) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroupFromClient(gid, c.client)
}

func creMsg(topic string, partition int32, key string, value string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.StringEncoder(value),
		Partition: partition,
	}
}

type Consumer struct {
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费者启动")
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费者关闭")
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("接收topic=%s, partition=%d, offset=%d, key=%s, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
	return nil
}
