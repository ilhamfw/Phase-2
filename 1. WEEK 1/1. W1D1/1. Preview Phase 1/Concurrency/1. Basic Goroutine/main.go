package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for c := 'a'; c <= 'j'; c++ {
		fmt.Println(string(c))
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Menunggu sebentar untuk memastikan kedua gorutin selesai
	time.Sleep(2 * time.Second)
}
