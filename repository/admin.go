package repository

import (
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
	"fmt"
)

func AdminLogin(adminDetail models.AdminDetail) (domain.Admin, error) {
	var adminCompareDetail domain.Admin
	if err := database.DB.Raw("select * from users where email = ? and isadmin=true", adminDetail.Email).Scan(&adminCompareDetail).Error; err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetail, nil
}

func DashboardUserDetails() (models.DashboardUser, error) {

	var userDetails models.DashboardUser
	err := database.DB.Raw("select count(*) from users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	err = database.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	return userDetails, nil
}

func DashBoardProductDetails() (models.DashBoardProduct, error) {

	var productDetails models.DashBoardProduct
	err := database.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	err = database.DB.Raw("select count(*) from products where quantity = 0").Scan(&productDetails.OutOfStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	return productDetails, nil
}

func GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	var userDetails []models.UserDetailsAtAdmin

	if page <= 0 {
		page = 1
	}

	if count <= 0 {
		count = 6
	}
	offset := (page - 1) * count
	fmt.Println("sjfwa", count, "sfji", offset)

	if err := database.DB.Raw("select id,firstname,lastname,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {

		return []models.UserDetailsAtAdmin{}, err
	}
	fmt.Println(userDetails)

	return userDetails, nil

}
