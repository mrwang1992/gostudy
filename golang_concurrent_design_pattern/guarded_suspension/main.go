package main

import (
	"fmt"
	"time"
)

type Requests struct {
	name string
}

func clientRun(queue chan *Requests, id int) {

	count := 0
	for {
		count += 1
		queue <- &Requests{
			name:fmt.Sprintf("requests %d-%d", id, count),
		}
		fmt.Printf("client send request => %d-%d \n", id, count)
	}
}

func serverRun(queue chan *Requests, id int) {
	for req := range queue{
		fmt.Printf("server %d handle %v \n", id, req)
	}
}

func main() {
	queue := make(chan *Requests, 3)

	go clientRun(queue, 1)
	go clientRun(queue, 2)
	go clientRun(queue, 3)

	go serverRun(queue, 1)
	go serverRun(queue, 2)
	go serverRun(queue, 3)

	time.Sleep(1000 * time.Second)
}
