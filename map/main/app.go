package main

import "fmt"

func main() {
	// Literal
	gajiKaryawan := map[string]int{}
	fmt.Println(gajiKaryawan)

	gajiKaryawan = map[string]int{
		"angga": 15000000,
		"adji":  8000000,
	}

	// Menambah Key Value
	gajiKaryawan["surya"] = 7000000
	fmt.Println(gajiKaryawan)

	//Memanggil Value dengan Key
	fmt.Println(gajiKaryawan["angga"])

	// Make
	gajiKaryawan2 := make(map[string]int)
	gajiKaryawan2["angga"] = 21000000
	fmt.Println(gajiKaryawan2)

	// Mengganti Key Value
	gajiKaryawan2["angga"] = 22000000
	fmt.Println(gajiKaryawan2)

	// Menghapus Key Value
	delete(gajiKaryawan2, "angga")
	fmt.Println(gajiKaryawan2["angga"])

	// Checking
	gajiKaryawan2["adji"] = 22000000
	val, exist := gajiKaryawan2["adji"]
	if exist {
		fmt.Println(val)
		delete(gajiKaryawan2, "adji")
	}
	fmt.Println(gajiKaryawan2["adji"])

}
