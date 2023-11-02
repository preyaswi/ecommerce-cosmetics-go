package repository

import (
	"errors"
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/helper"
	"firstpro/utils/models"

	"time"


	"gorm.io/gorm"
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
func TotalRevenue() (models.DashboardRevenue, error) {

	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	startTime, endTime = helper.GetTimeFromPeriod("month")
	err = database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	startTime, endTime = helper.GetTimeFromPeriod("year")
	err = database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	return revenueDetails, nil
}
func AmountDetails() (models.DashboardAmount, error) {

	var amountDetails models.DashboardAmount
	err := database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true ").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	err = database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'not paid' and shipment_status = 'processing' or shipment_status = 'pending' or shipment_status = 'order placed' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	return amountDetails, nil

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

	if err := database.DB.Raw("select id,firstname,lastname,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {

		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}
func GetUserByID(id int) (*domain.User, error) {

	var user domain.User
	result := database.DB.Where(&domain.User{ID: uint(id)}).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	if err := database.DB.Raw("select * from users where id=?", id).Scan(&user).Error; err != nil {
		return nil, nil
	}
	return &user, nil
}
func UpdateBlockUserByID(user *domain.User) error {
	err := database.DB.Exec("update users set blocked=? where id=?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func DashBoardOrder() (models.DashboardOrder, error) {

	var orderDetails models.DashboardOrder
	err := database.DB.Raw("select count(*) from orders where payment_status = 'paid' and approval = true ").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders where shipment_status = 'pending' or shipment_status = 'processing'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders where shipment_status = 'cancelled'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select sum(quantity) from order_items").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}
	return orderDetails, nil
}
func FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error) {

	var salesReport models.SalesReport

	result := database.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = database.DB.Raw("select count(*) from orders").Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = database.DB.Raw("select count(*) from orders where payment_status = 'paid' and approval = true and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = database.DB.Raw("select count(*) from orders where shipment_status = 'processing' and approval = false and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	var productID int
	result = database.DB.Raw("select product_id from order_items group by product_id order by sum(quantity) desc limit 1").Scan(&productID)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = database.DB.Raw("select name from products where id = ?", productID).Scan(&salesReport.TrendingProduct)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	return salesReport, nil
}

