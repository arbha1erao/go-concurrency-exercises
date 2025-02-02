//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, ch chan<- *Tweet) {
	defer close(ch)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		ch <- tweet
	}
}

func consumer(ch <-chan *Tweet) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	ch := make(chan *Tweet)
	var wg sync.WaitGroup
	wg.Add(2)

	// Producer
	go func() {
		defer wg.Done()
		producer(stream, ch)
	}()

	// Consumer
	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
