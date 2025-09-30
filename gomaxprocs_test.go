package golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("total goroutines:", totalGoroutines)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu:", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("total goroutines:", totalGoroutines)

	group.Wait()
}
