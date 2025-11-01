package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	b := 1
	for b < 10 {
		fmt.Println(b)
		b++
	}

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	var started bool
	var stopped atomic.Bool
	for {
		if !started {
			started = true
			fmt.Println("Started")
			go func() {
				for {
					select {
					case <-ctx.Done():
						fmt.Println("Goroutine finished")
						stopped.Store(true)
						return
					}
				}
			}()

		}
		fmt.Println("main")
		if stopped.Load() {
			break
		}

	}

	//遍历数组
	var arr [10]string
	arr[0] = "Hello"
	for i := range arr {
		fmt.Println("current index:", i)
	}

	for i, v := range arr {
		fmt.Println("current index:", i, "value:", v)
	}

	//遍历切片
	s := make([]string, 10)
	s[0] = "World"
	for i := range s {
		fmt.Println("current index:", i)
	}

	for i, v := range s {
		fmt.Println("current index:", i, "value:", v)
	}

	//遍历map
	m := make(map[string]string)
	m["1"] = "One"
	m["2"] = "Two"
	for k := range m {
		fmt.Println("current key:", k)
	}

	for k, v := range m {
		fmt.Println("current key:", k, "value:", v)
	}

	for i := 0; i < 5; i++ {
	outter:
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outter
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}

}
