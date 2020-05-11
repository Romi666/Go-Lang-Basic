package main

import (
	io "bufio"
	f "fmt"
	"os"
	"strconv"
)

func main() {
	var angka int
	// f.Print("Masukan angka : ")
	// f.Scan(&angka)
	// f.Println("Angka yang dimasukkan : ", angka)

	scanner := io.NewScanner(os.Stdin)
	f.Print("Tebak angka : ")
	scanner.Scan()
	angka, _ = strconv.Atoi(scanner.Text())
	switch angka {
	case 0:
		f.Println("0")
		fallthrough
	case 1:
		f.Println("1")
		fallthrough
	case 2:
		f.Println("2")
		fallthrough
	case 3:
		f.Println("3")
		fallthrough
	default:
		f.Println("Salahhhh")
	}
}
