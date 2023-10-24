package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go producer(ch,&wg)
	// go consumer(ch, &wg) // 单消费者
	// 多消费者
	go consumerId(ch,1, &wg)
	go consumerId(ch,2, &wg)
	wg.Wait()
}

// 生产者
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	} 
	close(ch)
}

// 消费者
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range ch {
		fmt.Printf("consumer data:%d\n", d)
	}
}

func consumerId(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range ch {
		fmt.Printf("consumer id:%d, data:%d\n",id, d)
	}
}