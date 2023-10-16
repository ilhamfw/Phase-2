package main

import (
	"fmt"
	"sync"
)

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

    wg.Add(2) // Jumlah gorutin yang akan ditunggu

    go printNumbers(&wg)
    go printLetters(&wg)

    wg.Wait() // Tunggu kedua gorutin selesai
}
