package main

import (
    "fmt"
)

func sendNumbers(evenCh, oddCh chan int) {
    for i := 1; i <= 20; i++ {
        if i%2 == 0 {
            evenCh <- i // Mengirim angka genap ke channel evenCh
        } else {
            oddCh <- i // Mengirim angka ganjil ke channel oddCh
        }
    }
    close(evenCh)
    close(oddCh)
}

func main() {
    evenCh := make(chan int)
    oddCh := make(chan int)

    // Menjalankan gorutin untuk mengisi channel dengan angka genap dan ganjil
    go sendNumbers(evenCh, oddCh)

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
        }

        // Keluar dari loop jika kedua channel telah dinonaktifkan
        if evenCh == nil && oddCh == nil {
            break
        }
    }
}
