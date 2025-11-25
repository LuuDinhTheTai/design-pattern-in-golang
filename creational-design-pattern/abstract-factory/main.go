package main

import "fmt"

func main() {
	adidas, _ := GetSportsFactory("Adidas")
	nike, _ := GetSportsFactory("Nike")

	adidasShoe := adidas.makeShoe()
	adidasShirt := adidas.makeShirt()

	nikeShoe := nike.makeShoe()
	nikeShirt := nike.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
