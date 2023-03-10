package belajar_golang_gouroutines

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	// buffered itu digunakan untuk menampung hasil pengiriman data
	channel := make(chan string, 3) // default buffer itu satu
	defer close(channel)

	go func() {
		channel <- "Rangga"
		channel <- "Surya"
		channel <- "Prayoga"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // wajib close channel agar tidak kena deadlock
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
}
