package main

import (
	"fmt"
	"sync"
	"time"
)

type car struct {
	Color  string
	Weight int
}

var (
	mtx   sync.Mutex
	total int32
)

func main() {
	pool := sync.Pool{}
	pool.New = func() interface{} {
		// 使用锁保证池内对象数量不超过限制
		mtx.Lock()
		defer mtx.Unlock()

		if total >= 5 {
			return nil
		}
		total++
		return new(car)
	}

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			c := pool.Get()
			if c != nil {
				fmt.Println("goroutine", i, ": get car success")
				time.Sleep(1 * time.Second)
				pool.Put(c)
			} else {
				fmt.Println("goroutine", i, ": get car failed")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 输出
// goroutine 9 : get car success
// goroutine 3 : get car success
// goroutine 4 : get car failed
// goroutine 0 : get car success
// goroutine 5 : get car success
// goroutine 6 : get car success
// goroutine 8 : get car failed
// goroutine 1 : get car failed
// goroutine 2 : get car failed
// goroutine 7 : get car failed
