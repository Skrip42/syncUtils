package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	syncutils "github.com/Skrip42/syncUtils"
)

func main() {
	cond := syncutils.NewCond()
	wg := sync.WaitGroup{}

	wg.Add(4)

	ctx, cancel := context.WithCancel(context.Background())

	// standard sync.Cond behavior
	go func() {
		<-cond.Wait()
		fmt.Println("the first goroutine finished waiting")
		wg.Done()
	}()
	go func() {
		<-cond.Wait()
		fmt.Println("the second goroutine finished waiting")
		wg.Done()
	}()
	go func() {
		<-cond.Wait()
		fmt.Println("the third goroutine finished waiting")
		wg.Done()
	}()

	time.Sleep(time.Second)
	cond.Signal() // the third goroutine finished waiting
	time.Sleep(time.Second)
	cond.Broadcast() // the second and first goroutine finished waiting

	// Signal and Bloadcast is not blocked
	time.Sleep(time.Second)
	cond.Signal()    // nothing happened
	cond.Broadcast() // nothing happened

	// canceled waiting by context
	go func() {
		select {
		case <-cond.Wait():
			fmt.Println("the fourth goroutine finished waiting")
		case <-ctx.Done():
			fmt.Println("fourth goroutine finished by context")
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	cancel() // fourth goroutine finished by context

	wg.Wait()
}
