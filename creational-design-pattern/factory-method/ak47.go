package factory_method

// Ak47 Concrete Product
type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}
}
