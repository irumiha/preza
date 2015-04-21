package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sortSize := 1000
	out := make(chan int32)
	start := make(chan struct{})

	for i := 1; i <= sortSize; i++ {
		newInt := rand.Int31n(int32(i))
		go sleepsort(newInt, start, out)
	}
	close(start)

	collectResults(sortSize, out)
}

func sleepsort(number int32, start chan struct{}, out chan int32) {
	<-start
	time.Sleep(time.Duration(number) * time.Millisecond)
	out <- number
}

func collectResults(sortSize int, rc chan int32) {
	for {
		select {
		case outInt := <-rc:
			fmt.Println(outInt)
			sortSize--
		default:
		}
		if sortSize == 0 {
			break
		}
	}
}
