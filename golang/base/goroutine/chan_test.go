package goroutine

import (
	"fmt"
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

// 多个发送者，一个接收者
func TestManySendAndOneReceive(t *testing.T) {
	var (
		sendNum = 3
		wg      = sync.WaitGroup{}
		numCh   = make(chan int)
		stopCh  = make(chan struct{})
		// 10毫秒后通知发送端停止发送数据
		after = time.After(10 * time.Millisecond)
	)
	wg.Add(1)

	// send
	for i := 0; i < sendNum; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					fmt.Println("收到退出信号")
					return
				case numCh <- 1:
					//fmt.Println("发送成功", value)
				}
			}
		}()
	}

	// receive
	go func() {
		for {
			select {
			case v := <-numCh:
				fmt.Println("接收到数据", v)
			case <-after:
				close(stopCh)
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}

// 多个发送者，多个接收者
func TestManySendAndManyReceive(t *testing.T) {
	var (
		maxRandomNumber = 5000
		receiver        = 10
		sender          = 10
		wg              = sync.WaitGroup{}
		numCh           = make(chan int)
		stopCh          = make(chan struct{})
		toStop          = make(chan string, 1)
		stoppedBy       string
	)
	wg.Add(receiver)

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < sender; i++ {
		go func(id string) {
			for {
				value := rand.Intn(maxRandomNumber)
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
				case numCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < receiver; i++ {
		go func(id string) {
			defer wg.Done()
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
				case value := <-numCh:
					if value == maxRandomNumber-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					t.Log(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wg.Wait()
	t.Log("stopped by", stoppedBy)
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

// 会引发channel panic的情况一:发送数据到已经关闭的channel
// panic: send on closed channel
func TestChannelPanic1(t *testing.T) {
	var ch = make(chan int)
	close(ch)
	time.Sleep(10 * time.Millisecond)
	go func() {
		ch <- 1
	}()
	t.Log(<-ch)
}

// 会引发channel panic的情况一的另外一种:发送数据时关闭channel
// panic: send on closed channel
func TestChannelPanic11(t *testing.T) {
	var ch = make(chan int)
	go func() {
		go func() {
			// 没有接收数据的地方,此处会一直阻塞
			ch <- 1
		}()
	}()

	time.Sleep(20 * time.Millisecond)
	close(ch)
}

// 会引发channel panic的情况二:重复关闭channel
// panic: close of closed channel
func TestChannelPanic2(t *testing.T) {
	var ch = make(chan int)
	close(ch)
	close(ch)
}

// 会引发channel panic的情况三:未初始化关闭
// panic: close of nil channel
func TestChannelPanic3(t *testing.T) {
	var ch chan int
	close(ch)
}

func TestChan(t *testing.T) {
	var (
		ch = make(chan int)
		wg = sync.WaitGroup{}
		// 10毫秒后通知发送端停止发送数据
		after = time.After(10 * time.Millisecond)
	)
	wg.Add(2)

	// send
	go func() {
		for {
			select {
			case <-after:
				close(ch)
				wg.Done()
				return
			default:
				ch <- 1
			}
		}
	}()

	// receive
	go func() {
		defer wg.Done()
		for v := range ch {
			t.Log(v)
		}
		return
	}()

	wg.Wait()
}
