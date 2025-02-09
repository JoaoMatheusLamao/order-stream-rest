package models

// Order is a struct that contains the order information
type Order struct {
	OrderID         string   `json:"order_id"`
	InternalOrderID string   `json:"internal_order_id"`
	Items           []Item   `json:"items"`
	Total           float64  `json:"total"`
	Customer        Customer `json:"customer"`
}

// Item is a struct that contains the item information
type Item struct {
	ItemID   string  `json:"item_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Customer is a struct that contains the customer information
type Customer struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}
