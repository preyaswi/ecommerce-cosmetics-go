package usecase

import (
	"errors"
	"firstpro/domain"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"

	"github.com/jinzhu/copier"
)

// func AddOrder(product_id int, address_id int, user_id int) (models.OrderResponse, error) {
// 	ok, _, err := repository.CheckProduct(product_id)
// 	//here the second return is category and we will use this later when we need to add the offer details later
// 	if err != nil {

// 		return models.OrderResponse{}, err
// 	}

// 	if !ok {
// 		return models.OrderResponse{}, errors.New("product Does not exist")
// 	}
// 	ok, err = repository.CheckAddress(address_id, user_id)
// 	if err != nil {
// 		return models.OrderResponse{}, err
// 	}
// 	if !ok {
// 		return models.OrderResponse{}, errors.New("address is not given in right format")
// 	}
// 	quantityOfProduct, err := repository.GetQuantityFromProductID(product_id)
// 	if err != nil {

//			return models.OrderResponse{}, err
//		}
//		if quantityOfProduct == 0 {
//			return models.OrderResponse{}, errors.New("out of stock")
//		}
//	}
func OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error) {
	var orderBody models.OrderIncoming
	err := copier.Copy(&orderBody, &orderFromCart)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderBody.UserID = uint(userID)
	cartExist, err := repository.DoesCartExist(userID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !cartExist {
		return domain.OrderSuccessResponse{}, errors.New("cart empty can't order")
	}

	addressExist, err := repository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}

	// get all items a slice of carts
	cartItems, err := repository.GetAllItemsFromCart(int(orderBody.UserID))
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	var orderDetails domain.Order
	var orderItemDetails domain.OrderItem
	// add general order details - that is to be added to orders table
	orderDetails = helper.CopyOrderDetails(orderDetails, orderBody)

	// get grand total iterating through each products in carts
	for _, c := range cartItems {
		orderDetails.GrandTotal += c.TotalPrice
	}
	discount_price, err := repository.GetCouponDiscountPrice(int(orderBody.UserID), orderDetails.GrandTotal)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	err = repository.UpdateCouponDetails(discount_price, orderDetails.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderDetails.FinalPrice = orderDetails.GrandTotal - discount_price
	if orderBody.PaymentID == 2 {
		orderDetails.PaymentStatus = "not paid"
		orderDetails.ShipmentStatus = "pending"
	}

	err = repository.CreateOrder(orderDetails)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	for _, c := range cartItems {
		// for each order save details of products and associated details and use order_id as foreign key ( for each order multiple product will be there)
		orderItemDetails.OrderID = orderDetails.OrderId
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = int(c.Quantity)
		orderItemDetails.TotalPrice = c.TotalPrice

		err := repository.AddOrderItems(orderItemDetails, orderDetails.UserID, c.ProductID, c.Quantity)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}

	}

	err = repository.UpdateUsedOfferDetails(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	orderSuccessResponse, err := repository.GetBriefOrderDetails(orderDetails.OrderId)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	return orderSuccessResponse, nil

}
func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil

}

func CancelOrders(orderId string, userId int) error {
	userTest, err := repository.UserOrderRelationship(orderId, userId)
	if err != nil {
		return err
	}
	if userTest != userId {
		return errors.New("the order is not done by this user")
	}
	orderProductDetails, err := repository.GetProductDetailsFromOrders(orderId)
	if err != nil {
		return err
	}
	shipmentStatus, err := repository.GetShipmentStatus(orderId)
	if err != nil {
		return err
	}
	if shipmentStatus == "delivered" {
		return errors.New("item already delivered, cannot cancel")
	}

	if shipmentStatus == "pending" || shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)
		return errors.New("the order is in" + message + ", so no point in cancelling")
	}

	if shipmentStatus == "cancelled" {
		return errors.New("the order is already cancelled, so no point in cancelling")
	}
	err = repository.CancelOrders(orderId)
	if err != nil {
		return err
	}
	err = repository.UpdateQuantityOfProduct(orderProductDetails)
	if err != nil {
		return err
	}
	return nil

}

func ExecutePurchaseCOD(userID int, addressID int) (models.Invoice, error) {
	ok, err := repository.CartExist(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	if !ok {
		return models.Invoice{}, errors.New("cart doesnt exist")
	}
	cartDetails, err := repository.DisplayCart(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	addresses, err := repository.GetAllAddress(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	Invoice := models.Invoice{
		Cart:        cartDetails,
		AddressInfo: addresses,
	}
	return Invoice, nil

}
