package main

import (
	"fmt"
	"sync"
)

func main() {
	// Tạo một WaitGroup
	var wg sync.WaitGroup

	// Tạo một channel để gửi số nguyên
	numbers := make(chan int)

	// Tăng WaitGroup lên 1 để chờ goroutine hoàn thành
	wg.Add(1)

	// Tạo một goroutine để tính tổng các số nguyên trong channel
	go func() {
		// Khi goroutine hoàn thành, giảm WaitGroup xuống 1
		defer wg.Done()

		var sum int
		for n := range numbers {
			sum += n
		}

		fmt.Println("Tổng các số là", sum)
	}()

	// Gửi các số nguyên vào channel
	numbers <- 1
	numbers <- 2
	numbers <- 3

	// Đóng channel để cho goroutine biết rằng không có số nào khác để xử lý
	close(numbers)

	// Chờ WaitGroup đến 0 để hoàn thành
	wg.Wait()
}
