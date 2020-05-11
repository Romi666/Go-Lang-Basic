package main

import f "fmt"

func tebasIsi(tebakan int) bool {
	switch tebakan {
	case 0, 4:
		f.Println("0")
		return true
	default:
		return false
	}
}
