package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) string

func fakeSearch(kind string) Search{
	return func(query string) string{
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%s result for %q\n", kind, query)
	}
}

func Google1(query string)(results []string){
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func Google2(query string)(results []string){
	c := make(chan string)
	go func(){c <- Web(query)}()
	go func(){c <- Image(query)}()
	go func(){c <- Video(query)}()

	for i := 0; i < 3; i++{
		result := <-c
		results = append(results, result)
	}
	return
}

func Google3(query string)(results []string){
	c := make(chan string)
	go func(){c <- Web(query)}()
	go func(){c <- Image(query)}()
	go func(){c <- Video(query)}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++{
		select{
		case result := <-c:
			results = append(results, result)
		case <- timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) string{
	c := make(chan string)
	searchReplica := func(i int){c <- replicas[i](query)}
	for i := range replicas{
		go searchReplica(i)
	}
	return <- c
}

func Google4(query string)(result string){
	result = First(query, fakeSearch("replica 1"), fakeSearch("replica 2"), fakeSearch("replica 3"))
	return
}

func main(){
	fmt.Println("NumCPU = " + strconv.Itoa(runtime.NumCPU()))
	fmt.Println("GOROOT = " + runtime.GOROOT())

	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google4("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}