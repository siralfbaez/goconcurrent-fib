package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count3("sheep", c)

	for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func count3(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
