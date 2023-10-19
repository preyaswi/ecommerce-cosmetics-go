package usecase

import (
	"errors"
	"firstpro/domain"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {

	productsBrief, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBrief{}, err
	}
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	return productsBrief, nil
}
func ShowIndividualProducts(id int) (*models.ProductBrief, error) {
	product, err := repository.ShowIndividualProducts(id)
	fmt.Println("ahgfvgf", product)
	if err != nil {
		fmt.Println("ahcfaecf")
		return &models.ProductBrief{}, err
	}
	fmt.Println("321423")
	return product, nil

}
func AddCategory(category domain.Category) (domain.Category, error) {
	productResponse, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return productResponse, nil
}
func UpdateCategory(current string, new string) (domain.Category, error) {
	result, err := repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}
	if !result {
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}
	newCat, err := repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCat, err
}
func DeleteCategory(categoryID string) error {
	err := repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}

func FilterCategory(data map[string]int) ([]models.ProductBrief, error) {

	err := repository.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	var productFromCategory []models.ProductBrief
	for _, id := range data {

		product, err := repository.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductBrief{}, err
		}
		for _, product := range product {

			quantity, err := repository.GetQuantityFromProductID(product.ID)
			if err != nil {
				return []models.ProductBrief{}, err
			}
			if quantity == 0 {
				product.ProductStatus = "out of stock"
			} else {
				product.ProductStatus = "in stock"
			}
			if product.ID != 0 {
				productFromCategory = append(productFromCategory, product)
			}
		}

		// if a product exist for that genre. Then only append it

	}
	return productFromCategory, nil

}

func AddProduct(products domain.Products) (domain.Products, error) {
	productResponse, err := repository.AddProduct(products)
	if err != nil {
		return domain.Products{}, err
	}
	return productResponse, nil
}
func UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error) {

	result, err := repository.CheckProductExist(pid)
	if err != nil {

		return models.ProductUpdateReciever{}, err
	}

	if !result {
		return models.ProductUpdateReciever{}, errors.New("there is no product as you mentioned")
	}

	newcat, err := repository.UpdateProduct(pid, stock)
	if err != nil {
		return models.ProductUpdateReciever{}, err
	}

	return newcat, err
}
func DeleteProduct(productID string) error {
	err := repository.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}
