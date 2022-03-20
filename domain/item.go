package domain

type Item struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ItemOrder struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type RequestItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type Items []Item

func GetItems() Items {
	return items
}

// seed data
var items = Items{
	Item{
		SKU:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 10,
	},
	Item{
		SKU:      "43N23P",
		Name:     "MacBook Pro",
		Price:    5399.99,
		Quantity: 5,
	},
	Item{
		SKU:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 10,
	},
	Item{
		SKU:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.00,
		Quantity: 2,
	},
}
