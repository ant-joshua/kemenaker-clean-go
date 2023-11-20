package domains

type GetAllProductFilter struct {
	Page  int     `json:"page" validate:"required"`
	Limit int     `json:"limit" validate:"required"`
	Name  *string `json:"name"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Categories  []int64 `json:"categories"`
}

type UpdateProductRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Categories  []int64 `json:"categories"`
}

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name"`
	Price       int    `json:"price" gorm:"column:price"`
	Description string `json:"description" gorm:"column:description"`
	//Categories  []Category `json:"categories"`
}
