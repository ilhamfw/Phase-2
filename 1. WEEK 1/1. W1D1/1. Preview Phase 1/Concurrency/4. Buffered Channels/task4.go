package main

import (
    "fmt"
    "sync"
)

func produce(ch chan int, wg *sync.WaitGroup) {
    for i := 1; i <= 10; i++ {
        ch <- i
        fmt.Printf("Producing: %d\n", i)
    }
    close(ch) // Menutup channel setelah selesai mengirim
    wg.Done()  // Sinyal bahwa gorutin ini telah selesai
}

func consume(ch chan int, wg *sync.WaitGroup) {
    for num := range ch {
        fmt.Printf("Consuming: %d\n", num)
    }
    wg.Done() // Sinyal bahwa gorutin ini telah selesai
}

func printNumbers(wg *sync.WaitGroup) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d\n", i)
    }
    wg.Done() // Sinyal bahwa gorutin ini telah selesai
}

func printLetters(wg *sync.WaitGroup) {
    for c := 'a'; c <= 'j'; c++ {
        fmt.Printf("%c\n", c)
    }
    wg.Done() // Sinyal bahwa gorutin ini telah selesai
}

func main() {
    var wg sync.WaitGroup

    // Menggunakan WaitGroup untuk menunggu semua gorutin selesai
    wg.Add(4) // Jumlah gorutin yang akan ditunggu

    // Membuat buffered channel dengan kapasitas 5
    ch := make(chan int, 5)

    // Menjalankan fungsi produce dalam gorutin
    go produce(ch, &wg)

    // Menjalankan fungsi consume dalam gorutin
    go consume(ch, &wg)

    // Menjalankan fungsi printNumbers dalam gorutin
    go printNumbers(&wg)

    // Menjalankan fungsi printLetters dalam gorutin
    go printLetters(&wg)

    // Menunggu hingga semua gorutin selesai
    wg.Wait()
}
