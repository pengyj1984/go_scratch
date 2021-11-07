package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Count(ch chan int, i int) {
	time.Sleep(3 * 1e9)
	fmt.Printf("Counting %d, time: %d\n", i, time.Now().Nanosecond())
	ch <- i
}

func Read(ch chan int) {
	for i := 0; i < 10; i++ {
		v := rand.Int() & 15
		for j := 0; j < v; j++ {
			ch <- j
		}
		fmt.Printf("Read %d numbers at time %d\n", v, time.Now().Nanosecond())
	}
	close(ch)
}

func main() {
	fmt.Printf("Start at %d\n", time.Now().Nanosecond())
	//chs := make([]chan int , 10)
	//for i := 0; i < 10; i++{
	//	chs[i] = make(chan int)
	//	go Count(chs[i], i)
	//}
	//
	//for _, ch := range chs{
	//	fmt.Printf("finish %d, time: %d\n", <-ch, time.Now().Nanosecond())
	//}

	fmt.Printf("test 2\n")
	//ch := make(chan int, 1)
	//i := 0
	//for{
	//	select{
	//		case ch <- 0:
	//		case ch <- 1:
	//		case ch <- 2:
	//		case ch <- 3:
	//	}
	//	fmt.Println("value = ", <-ch)
	//	i++
	//	if i >= 100 {
	//		break
	//	}
	//}
	ch2 := make(chan int, 16)
	go Read(ch2)

	for v := range ch2 {
		fmt.Printf("Get %d elems from channel at %d\n", v, time.Now().Nanosecond())
	}
}
