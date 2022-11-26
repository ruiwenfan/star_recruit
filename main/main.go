package main

// 这个文件下测试的都是不需要http服务的

func main() {
	// test regexp

	/* 	var n int
	   	fmt.Scanln(&n)
	   	students := make([]string, 10)
	   	for i := 0; i < n; i++ {
	   		fmt.Scanln(&students[i])
	   	}
	   	for i := 0; i < n; i++ {
	   		fmt.Println(students[i])
	   	}
	   	fmt.Println("-----answer-----")
	   	for i := 0; i < n; i++ {
	   		if star.IsValid(students[i]) {
	   			star.GetAnswer(students[i])
	   		} else {
	   			fmt.Println("False")
	   		}
	   	} */

	// test bubbleSort
	/* 	test := []int{2, 3, 1, 4, 2, 0, 1, 10, -1, -3, -10, 0, 239}
	   	star.BubbleSort(test)
	   	fmt.Println(test) */
	// test Singleton
	/* 	var wg sync.WaitGroup
	   	wg.Add(10)
	   	for i := 0; i < 10; i++ {
	   		go func() {
	   			log.Printf("%p", star.GetSingleton2())
	   			wg.Done()
	   		}()
	   	}
	   	wg.Wait() */

	/*
		 	ch := make(chan int, 5)

			var wg sync.WaitGroup
			wg.Add(2)
			n := 0
			for i := 0; i < 3; i++ {
				go func() {
					n++
					star.Consumer(ch, n)
				}()
			}
			wg.Done()
			go func() {
				star.Producer(ch)
				wg.Done()
			}()
			wg.Wait()
	*/

	/*
		 	if err := star.CreatePost(star.Post{Id: 124, Content: "i am not happy!", UserId: 1}); err != nil {
				fmt.Println("star.CreatePost failed err is ", err)
			}
	*/
	/* 	if err := star.UpdatePost(123, "be happy everyday!"); err != nil {
		log.Println("UpdatePost failed err is ", err)
	} */

	/*
		 	if err := star.DeletePost(123); err != nil {
				log.Println("DeletePost failed err is", err)
			}
	*/

	/*
		 	if err := star.CreatePost(star.Post{Id: 555, Content: "today is 2022.11.26", UserId: 1}); err != nil {
				log.Println("CreatePost failed err is", err)
			}
	*/

	/*
		 	postAndUser, err := star.SearchPostAndUser(555)
			if err != nil {
				fmt.Println("SearchPostAndUser failed err is", err)
			}
			fmt.Println(*postAndUser)
	*/
}
