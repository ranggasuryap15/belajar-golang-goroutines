package belajar_golang_gouroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // ketika menggunakan wait group maka wajib melakukan proses done.

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // gantinya time.Sleep karena time.Sleep itu hanya mengira-ngira waktunya saja
	fmt.Println("Complete")
}
