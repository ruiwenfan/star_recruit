package star

import (
	"fmt"
	"sync"
)

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

func Consumer(workShop chan int, id int) {
	for {
		fmt.Println("Consumer is ", id, "work is ", <-workShop)
	}
}
