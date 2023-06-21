package kafka

import (
	"testing"
)

func TestClient(t *testing.T) {
	client, err := newClient()
	if err != nil {
		panic(err)
	}
	topics, err := client.Topics()
	if err != nil {
		panic(err)
	}
	t.Log(topics)
}

func TestSendMsg(t *testing.T) {
	err := sendMsg("mytopic", "msg")
	if err != nil {
		t.Fatal(err)
	}
}
