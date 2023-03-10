package belajar_golang_gouroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	// mutex itu untuk mengatasi race condition, ini agak sedikit lambat, tapi tidak terlalu lambat banget karena cuma beda beberapa nano second

	time.Sleep(7 * time.Second)
	fmt.Println("Counter :", x)
}
