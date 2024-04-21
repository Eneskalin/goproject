package items

type Item struct {
	Name      string
	Price     float64
	Discount  float64
	LastPrice float64
}

func ItemList() []Item {
	product1 := Item{
		Name:     "water",
		Price:    1.20,
		Discount: 0,
	}
	product2 := Item{
		Name:     "coffee",
		Price:    8.00,
		Discount: 0.10,
	}
	product3 := Item{
		Name:     "tea",
		Price:    2.00,
		Discount: 0.10,
	}
	product4 := Item{
		Name:     "cookie",
		Price:    12.00,
		Discount: 0.30,
	}
	productList := []Item{product1, product2, product3, product4}

	return productList

}
