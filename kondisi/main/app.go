package main

import "log"

func main() {

	//IF

	//IF ELSE

	//ELSE IF

	//NESTED IF
	math := "A"
	eng := "C"

	if math == "A" {
		if eng == "A" || eng == "b" {
			log.Printf(`nilai math %s dan nilai %s, selamat anda lulus`, math, eng)
		} else {
			log.Printf(`nilai math %s dan nilai %s, maaf anda tidak lulus`, math, eng)
		}
	} else {
		log.Printf(`nilai math %s, maaf anda tidak lulus`, math)
	}

	//SWITCH CASE
	math = "C"

	switch math {
	case "A":
		log.Print("Nilai A")
	case "B":
		log.Print("Nilai B")
	case "C":
		log.Print("Nilai C")
	case "D":
		log.Print("Nilai D")
	}
}
