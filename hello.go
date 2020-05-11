package main

import (
	io "bufio"
	f "fmt"
	"os"
)

func main() {
	// var angka int
	// f.Print("Masukan angka : ")
	// f.Scan(&angka)
	// f.Println("Angka yang dimasukkan : ", angka)

	scanner := io.NewScanner(os.Stdin)
	f.Print("Masukan kalimat : ")
	scanner.Scan()
	f.Println("Kalimat yang dimasukan : " + scanner.Text() + "")
}
