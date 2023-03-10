package belajar_golang_gouroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Rangga Surya Prayoga"
		fmt.Println("Selesai mengirim data ke channel:", channel)
	}()

	data := <-channel
	fmt.Println(data)
}
