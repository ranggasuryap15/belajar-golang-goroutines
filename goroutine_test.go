package belajar_golang_gouroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // tidak cocok apabila menggunakan return value
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
