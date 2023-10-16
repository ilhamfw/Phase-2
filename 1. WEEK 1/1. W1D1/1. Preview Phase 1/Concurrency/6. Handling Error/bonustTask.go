package main

import (
	"fmt"
)

func sendNumbers(evenCh, oddCh chan int, errorCh chan error) {
	for i := 1; i <= 22; i++ {
		if i > 20 {
			errorCh <- fmt.Errorf("error %d", i) // Mengirim error ke channel errorCh
		} else if i%2 == 0 {
			evenCh <- i // Mengirim angka genap ke channel evenCh
		} else {
			oddCh <- i // Mengirim angka ganjil ke channel oddCh
		}
	}
	close(evenCh)
	close(oddCh)
	close(errorCh)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	errorCh := make(chan error)

	// Menjalankan gorutin untuk mengisi channel dengan angka genap, ganjil, dan error
	go sendNumbers(evenCh, oddCh, errorCh)

	for {
		select {
		case even, ok := <-evenCh:
			if !ok {
				evenCh = nil // Menonaktifkan channel
			} else {
				fmt.Printf("Received even: %d\n", even)
			}
		case odd, ok := <-oddCh:
			if !ok {
				oddCh = nil // Menonaktifkan channel
			} else {
				fmt.Printf("Received odd: %d\n", odd)
			}
		case err, ok := <-errorCh:
			if !ok {
				errorCh = nil // Menonaktifkan channel
			} else {
				fmt.Printf("Received error: %v\n", err)
			}
		}

		// Keluar dari loop jika semua channel telah dinonaktifkan
		if evenCh == nil && oddCh == nil && errorCh == nil {
			break
		}
	}
}
