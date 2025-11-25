package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 10,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 10,
		},
	}
}
