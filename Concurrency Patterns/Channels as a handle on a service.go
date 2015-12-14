package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-joe)
		fmt.Printf("%s\n", <-ann)
	}
	fmt.Println("You're both boring, I'm leaving")
}

func boring(msg string) <- chan string {
	c := make(chan string)
	go func() {
		for i := 0;; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
