package repository

import (
	"errors"
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
	"fmt"
	"strconv"

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
func CheckCategory(current string) (bool, error) {
	var i int
	err := database.DB.Raw("select count(*) from categories where category_name =?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	}
	return true, err

}
func UpdateCategory(current, new string) (domain.Category, error) {
	if database.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}
	if err := database.DB.Exec("update categories set category_name =$1 where category_name =$2", new, current).Error; err != nil {
		return domain.Category{}, err
	}
	var newcat domain.Category
	if err := database.DB.First(&newcat, "category_name=?", new).Error; err != nil {
		return domain.Category{}, err
	}
	return newcat, nil

}
func DeleteCategory(categoryID string) error {
	id, err := strconv.Atoi(categoryID)
	fmt.Println(id)
	if err != nil {
		return errors.New("couldn't convert")
	}
	result := database.DB.Exec("delete from categories where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
func GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := database.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}
func GetPriceOfProductFromID(prodcut_id int) (float64, error) {
	var productPrice float64

	if err := database.DB.Raw("select price from products where id = ?", prodcut_id).Scan(&productPrice).Error; err != nil {
		return 0.0, err
	}
	return productPrice, nil
}

func CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := database.DB.Raw("select count(*) from categories where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}

		if count < 1 {
			return errors.New("genre does not exist")
		}
	}
	return nil
}
func GetProductFromCategory(id int) ([]models.ProductBrief, error) {

	var product []models.ProductBrief
	err := database.DB.Raw(`
		SELECT *
		FROM products
		JOIN categories ON products.category_id = categories.id
		 where categories.id = ?
	`, id).Scan(&product).Error

	if err != nil {
		return []models.ProductBrief{}, err
	}
	return product, nil
}
