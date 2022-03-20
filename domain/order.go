package domain

type Order struct {
	ID        string      `json:"id"`
	Customer  string      `json:"customer"`
	Items     []ItemOrder `json:"items"`
	Total     float64     `json:"total"`
	CreatedAt string      `json:"created_at"`
}

type OrderUsecase interface {
	DoOrder(requestItems []RequestItem) (Order, error)
}
