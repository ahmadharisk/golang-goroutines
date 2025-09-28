package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	// 1) membuat kondisi lock
	cond.L.Lock()

	// 2) kemudiam membuat
	cond.Wait()

	fmt.Println("selesai", value)

	// 3) di unlock
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Signal()
	//	}
	//}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
		}
	}()

	group.Wait()
}
