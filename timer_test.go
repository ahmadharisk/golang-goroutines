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

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		time.Sleep(8 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(2 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}
