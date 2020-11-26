package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	totalNum := make(chan int, 300)
	prineNum := make(chan int, 200)
	quit := make(chan bool)
	goruoutineCount := 4

	go func() {
		for i := 0; i < 8000; i++ {
			totalNum <- i
		}
		close(totalNum)
	}()

	for j := 0; j < goruoutineCount; j++ {
		go dealNume2(totalNum, prineNum, quit)
	}

	go func() {
		for i := 0; i < goruoutineCount; i++ {
			fmt.Println(<-quit)
		}
		close(quit)
		close(prineNum)
	}()

	for prime := range prineNum {
		fmt.Println(prime)
	}

	fmt.Println("使用时间：", time.Now().Sub(startTime))
}

func dealNume2(totalNum chan int, prineNum chan int, quit chan bool) {
	for num := range totalNum {
		var isPrime bool
		isPrime = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
			}
		}
		if isPrime {
			prineNum <- num
		}
	}
	quit <- true
}
