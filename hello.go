package main

import (
	f "fmt"
	"time"
)

func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// f.Print(sum)
	detik := 0
	for {
		time.Sleep(1 * time.Second)
		detik++
		f.Println(detik)
		if detik == 10 {
			break
		}
	}
	f.Println("Selesai")
}
