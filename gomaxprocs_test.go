package belajar_golang_gouroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goruoutine:", totalGoroutine)
	// biasanya dapat nilai 2 karena untuk menjalankan function ini dan garbage collection
}
