package belajar_golang_gouroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x = 0
	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	// race condition itu balapan menjalankan goroutine secara bersamaan dengan waktu yang presisi

	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
}
