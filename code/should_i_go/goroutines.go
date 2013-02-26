package main

import (
	"fmt"
	"time"
)

func main() {
	go say("I'm last", 3)
	go say("I'm middle", 2)
	go say("I'm first", 1)
	time.Sleep(4 * time.Second)
}

func say(text string, secs int) {
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Println(text)
}
