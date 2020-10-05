package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numb string
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&numb)
	fmt.Println("Angka yang dimasukkan : ", numb)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kalimat : ")
	scanner.Scan()
	fmt.Println("Kalimat yang dimasukkan : '" + scanner.Text() + "'")
}
