package entity

type Product struct {
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       int8    `json:"price"`
	ProductType string  `json:"product_type"`
	Size        float32 `json:"size"`
}
