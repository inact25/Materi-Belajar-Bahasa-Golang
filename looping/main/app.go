package main

import "fmt"

func main() {
	//BASIC FOR
	for i := 0; i <= 5; i++ {
		fmt.Printf("angka ke - %v \n", i)
	}

	//FOR A WHILE
	i := 0
	for i <= 5 {
		i++
		fmt.Printf("jumlah sekarang - %v \n", i)

	}
	fmt.Printf("Total - %v \n", i)

	//FOR EVER
	for {
		fmt.Println("This command will be executed again and again forever")
	}
}
