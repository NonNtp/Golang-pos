package dto

type ProductRequest struct {
	Name       string  `form:"name" binding:"required"`
	SKU        string  `form:"sku" binding:"required"`
	Desc       string  `form:"desc"`
	Price      float64 `form:"price" binding:"required"`
	Status     uint    `form:"status" binding:"required"`
	CategoryId uint    `form:"categoryId" binding:"required"`
}

type CreateOrUpdateProductResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Desc       string  `json:"desc"`
	Price      float64 `json:"price"`
	Status     uint    `json:"status"`
	CategoryId uint    `json:"categoryId"`
	Image      string  `json:"image"`
}

type ReadProductResponse struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	SKU      string           `json:"sku"`
	Desc     string           `json:"desc"`
	Price    float64          `json:"price"`
	Status   uint             `json:"status"`
	Category CategoryResponse `json:"category"`
	Image    string           `json:"image"`
}
