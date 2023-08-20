package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"sync"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints: []string{"192.168.10.43:12379", "192.168.10.43:22379", "192.168.10.43:32379"},
		// TODO 怎么检测到错误
		Endpoints:   []string{"192.168.10.43:12379, 192.168.10.43:22379, 192.168.10.43:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)

	_, err = concurrency.NewSession(cli, concurrency.WithContext(ctx))
	if err != nil {
		log.Println(err)
		return
	}
	return

	var (
		wg    = &sync.WaitGroup{}
		num   = 1000
		count = 0
	)

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			sess, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
			if err != nil {
				log.Println(err)
				return
			}
			defer sess.Close()
			mu := concurrency.NewMutex(sess, "lock")
			ctx := context.Background()
			ctx, _ = context.WithTimeout(ctx, 10*time.Second)
			err = mu.Lock(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			count++
			err = mu.Unlock(ctx)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	time.Sleep(10 * time.Second)
	//wg.Wait()
	fmt.Println(count)
}
