package app

type Product struct{ // структура товара
	Id int64 `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
	Description string `json:"description"`
	Author string `json:"author"`
	Photo string `json:"photo"`
}
