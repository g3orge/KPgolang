package main

import (
	"fmt"
	"time"
)

func counter(c chan int) {
	for i := 0; ; i++ {
		c <- i
	}
}
func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan int = make(chan int)

	go counter(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
