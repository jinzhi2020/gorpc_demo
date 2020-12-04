package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetch(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop!")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Downloading...")
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
	wg.Add(1)
	go fetch(ctx)
	wg.Wait()
}
