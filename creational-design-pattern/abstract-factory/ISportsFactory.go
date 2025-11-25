package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "Adidas" {
		return &Adidas{}, nil
	}

	if brand == "Nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("invalid brand")
}
