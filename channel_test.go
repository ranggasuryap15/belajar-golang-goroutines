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

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rangga Surya Prayoga"
	fmt.Println("Selesai mengirim data ke channel:", channel)
}

func OnlyIn(channel chan<- string) { // untuk mengirim data
	time.Sleep(2 * time.Second)
	channel <- "Rangga Surya Prayoga"
}

func OnlyOut(channel <-chan string) { // untuk menerima data
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}
