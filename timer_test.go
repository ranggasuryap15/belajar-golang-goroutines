package belajar_golang_gouroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	waktu := <-channel
	fmt.Println(waktu)
}

func TestAfterFunc(t *testing.T) {
	// untuk men-delay pekerjaan saat mengeksekusi
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println(time.Now())
		fmt.Println("Execute after 1 second")
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
