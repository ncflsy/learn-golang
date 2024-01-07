package main

import "fmt"

func main() {
	fmt.Println("halo")

	var firstName = "Nico"
	var lastName = "Flsy"

	fmt.Println(firstName, lastName)

	const date, month = "22", "Oktober"

	fmt.Println(date, month)

	fmt.Printf("halo nama saya %s %s. tanggal lahir saya adalah %s %s", firstName, lastName, date, month)
}