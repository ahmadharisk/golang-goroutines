package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("ayam")
	pool.Put("bebek")
	pool.Put("cicak")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("complete")
}
