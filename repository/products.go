package repository

import (
	"errors"
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
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

func AddProduct(product domain.Products) (domain.Products, error) {
	var p models.ProductReceiver
	err := database.DB.Raw("inption,brand_id,sert into products (name,sku,category_id,design_descriquantity,price,product_status) values (?,?,?,?,?,?,?,?) returning name,sku,category_id,design_description,brand_id,quantity,price,product_status", product.Name, product.SKU, product.CategoryID, product.DesignDescription, product.BrandID, product.Quantity, product.Price, product.ProductStatus).Scan(&p).Error
	if err != nil {
		return domain.Products{}, err
	}
	var productResponse domain.Products
	err = database.DB.Raw("select * from products where products.name=?", p.Name).Scan(&productResponse).Error
	if err != nil {
		return domain.Products{}, err
	}
	return productResponse, nil

}

func CheckProductExist(pid int) (bool, error) {
	var k int
	err := database.DB.Raw("SELECT COUNT(*) FROM products WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}
func UpdateProduct(pid int, quantity int) (models.ProductUpdateReciever, error) {

	// Check the database connection
	if database.DB == nil {
		return models.ProductUpdateReciever{}, errors.New("database connection is nil")
	}

	// Update the
	if err := database.DB.Exec("UPDATE products SET quantity = quantity + $1 WHERE id= $2", quantity, pid).Error; err != nil {
		return models.ProductUpdateReciever{}, err
	}

	// Retrieve the update
	var newdetails models.ProductUpdateReciever
	var newQuantity int
	if err := database.DB.Raw("SELECT quantity FROM products WHERE id=?", pid).Scan(&newQuantity).Error; err != nil {
		return models.ProductUpdateReciever{}, err
	}
	newdetails.ProductID = pid
	newdetails.Quantity = newQuantity

	return newdetails, nil
}
func DeleteProduct(productID string) error {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return errors.New("couldn't convert")
	}
	result := database.DB.Exec("delete from products where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
func AddCategory(category domain.Category) (domain.Category, error) {
	var b string
	err := database.DB.Raw("insert into categories (category_name) values (?) returning category_name", category.CategoryName).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoryResponse domain.Category
	err = database.DB.Raw("SELECT id ,category_name FROM categories WHERE category_name = ?", b).Scan(&categoryResponse).Error
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
