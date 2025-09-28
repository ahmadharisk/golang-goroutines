package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	waktu := <-timer.C
	fmt.Println(waktu)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	waktu := <-channel
	fmt.Println(waktu)
}

func TestAfterFunc(t *testing.T) {
	fmt.Println(time.Now())
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("execute after 5 sec")
		fmt.Println(time.Now())
		group.Done()
	})

	group.Wait()
}
