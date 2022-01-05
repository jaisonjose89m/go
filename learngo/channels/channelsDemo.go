package channels

import (
	"fmt"
	"sync"
)

func UnbufferedChannelDemo() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) { // bidirectional channel
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) { // bidirectional channel
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}

func BufferedChannelDemo() {
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { // receive only channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // send only channel
		ch <- 50
		ch <- 55
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
