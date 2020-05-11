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
	if angka == 0 {
		f.Println("Betul yang dimaksud", angka)
	} else {
		f.Println("Salah")
	}
	// f.Println("Kalimat yang dimasukan : " + scanner.Text() + "")
}
