package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		// time.Sleep(1 * time.Second)
		<-c
	}()
	c <- true
}
