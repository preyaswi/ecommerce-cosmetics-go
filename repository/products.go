package repository

import (
	"errors"
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
	"fmt"

	"gorm.io/gorm"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	if page <= 0 {
		page = 1
	}

	if count <= 0 {
		count = 5
	}

	offset := (page - 1) * count
	var productsBrief []models.ProductBrief

	err := database.DB.Raw(`
		SELECT * FROM products limit ? offset ?
	`, count, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(productsBrief)

	return productsBrief, nil

}
func ShowIndividualProducts(id int) (*models.ProductBrief, error) {
	var product models.ProductBrief
	result := database.DB.Raw(`
SELECT 
       p.id,
	   p.name,
	   p.sku,
	   c.category_name,
	   p.brand_id,
	   p.quantity,
	   p.price,
	   p.product_status
FROM
	   products p
JOIN
	   categories c ON p.category_id=c.id
WHERE
	   p.id=?`, id).Scan(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &product, nil
}
func AddCategory(category domain.Category) (domain.Category, error) {
	var b string
	err := database.DB.Raw("insert into categories (category_name) values (?) returning category_name", category.CategoryName).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoryResponse domain.Category
	err = database.DB.Raw("SELECT C.id ,C.category_name FROM categories c WHERE c.category_name = ?", b).Scan(&categoryResponse).Error
	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil
}
