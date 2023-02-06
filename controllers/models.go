package controllers

// Order contains the order information
type Order struct {
	Id           string  `json:"id"`
	Status       string  `json:"status"`
	Items        []Items `json:"items"`
	Total        float32 `json:"total"`
	CurrencyUnit string  `json:"currencyunit"`
}

// Items contains the item information
type Items struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
