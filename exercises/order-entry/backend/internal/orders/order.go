package orders

type Order struct {
	ID          int     `json:"id"`
	OrderNumber string  `json:"orderNumber"`
	Description string  `json:"description"`
	Total       float64 `json:"total"`
}
