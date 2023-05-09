package test

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 发送次数等于接收次数下的channel
// 此种情况下，goroutine在发送和结束时各自结束，通道会由于没有任何代码使用而被GC回收
func TestChanSendEqReceive(t *testing.T) {
	var (
		ch = make(chan int)
		wg = sync.WaitGroup{}
	)
	wg.Add(2)

	// send
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	// receive
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch)
		}
	}()

	wg.Wait()
}

// 多个发送者，一个接收者，
func TestManySendAndOneReceive(t *testing.T) {
	var (
		maxNum  = 3
		sendNum = 5
		wg      = sync.WaitGroup{}
		numCh   = make(chan int, 10)
		stopCh  = make(chan struct{})
	)

	wg.Add(1)

	// send
	for i := 0; i < sendNum; i++ {
		go func() {
			for {
				value := rand.Intn(maxNum)
				select {
				case <-stopCh:
					fmt.Println("收到退出信号")
					return
				case numCh <- value:
					//fmt.Println("发送成功", value)
				}
			}
		}()
	}

	// receive
	go func() {
		defer wg.Done()
		for value := range numCh {
			// 如果随机到0，则通知退出
			fmt.Println("接收成功", value)
			if value == 0 {
				close(stopCh)
				return
			}
		}
	}()

	wg.Wait()
}

// 多个发送者，多个接收者
func TestManySendAndManyReceive(t *testing.T) {
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// 提前关闭goroutine
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				// 提前关闭goroutine
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

// channel多重赋值，可以用来判断channel是否关闭
func TestChanMultipleAssignment(t *testing.T) {
	var ch = make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
	}()

	go func() {
		for i := 0; i < 3; i++ {
			v, ok := <-ch
			// 第一次v能获取到通道中的值，ok为true，后续全部为false
			// 注：当通道关闭时，继续获取通道的值，会获取该通道类型对应的零值
			fmt.Println(i, ok, v)
		}
	}()

	time.Sleep(time.Second)
}

// channel使用for range，可以用来判断channel是否关闭
func TestChanForRange(t *testing.T) {
	var ch = make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
	}()

	go func() {
		for v := range ch {
			fmt.Println("循环了一次")
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)
}
