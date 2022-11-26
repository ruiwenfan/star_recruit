package star

import (
	"fmt"
	"sync"
)

// 生产者，想chan写入数据
func Producer(workShop chan int) {
	var wg sync.WaitGroup
	n := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		workShop <- n
		n++
	}
	wg.Wait()
	close(workShop)
}

// 消费者
func Consumer(workShop chan int, id int) {
	for {
		fmt.Println("Consumer is ", id, "work is ", <-workShop)
	}
}
