package usecase

import (
	"errors"
	"firstpro/domain"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(adminDetail models.AdminDetail) (domain.TokenAdmin, error) {

	adminCompareDetails, err := repository.AdminLogin(adminDetail)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetail.Password))
	fmt.Println(adminCompareDetails.Password, "😂", adminDetail.Password, "😁")
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	tokenString, err := helper.GenerateTokenAdmin(adminDetailResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}
	return domain.TokenAdmin{
		Admin: adminDetailResponse,
		Token: tokenString,
	}, nil

}
func DashBoard() (models.CompleteAdminDashboard, error) {

	userDetails, err := repository.DashboardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := repository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	return models.CompleteAdminDashboard{

		DashboardUser:    userDetails,
		DashBoardProduct: productDetails,
		DashboardOrder:   orderDetails,
	}, nil

}
func GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := repository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func BlockUser(id int) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}
func UnBlockUser(id int) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if !user.Blocked {
		return errors.New("user is already unblocked")
	} else {
		user.Blocked = true
	}
	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func ApproveOrder(orderID string) error {
	ok, err := repository.CheckOrderID(orderID)
	if !ok {
		return err
	}

	shipmentStatus, err := repository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "cancelled" {

		return errors.New("the order is cancelled, cannot approve it")
	}

	if shipmentStatus == "pending" {

		return errors.New("the order is pending, cannot approve it")
	}
	if shipmentStatus == "processing" {
		fmt.Println("reached here")
		err := repository.ApproveOrder(orderID)

		if err != nil {
			return err
		}

		return nil
	}

	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil

}

func CancelOrderFromAdminSide(orderID string) error {

	orderProducts, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}

	err = repository.CancelOrders(orderID)
	if err != nil {
		return err
	}

	// update the quantity to products since the order is cancelled
	err = repository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return err
	}

	return nil

}
