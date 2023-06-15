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
