package domain

type Products struct {
	Id    uint   `json:"id" gorm:"column:id; primaryKey"`
	Title string `json:"title" gorm:"column:title"`
	Desc  string `json:"desc" gorm:"desc"`
}

type ProductForm struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (Products) TableName() string {
	return "products"
}
