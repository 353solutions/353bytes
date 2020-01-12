package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fanIn(chans ...<-chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan string) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func producer(name string, count int) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			msg := fmt.Sprintf("%s - %d", name, i)
			ch <- msg
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().Unix())
	ch1, ch2 := producer("p1", 8), producer("p2", 5)
	for val := range fanIn(ch1, ch2) {
		fmt.Println(val)
	}
}
