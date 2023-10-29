package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// func AddOrder(c *gin.Context) {
// 	id := c.Param("id")
// 	product_id, err := strconv.Atoi(id)
// 	if err != nil {
// 		errResponse := response.ClientResponse(http.StatusBadGateway, "Prodcut id is given in the wrong format", nil, err.Error())
// 		c.JSON(http.StatusBadGateway, errResponse)
// 		return
// 	}
// 	addressid := c.Param("address_id")
// 	address_id, err := strconv.Atoi(addressid)
// 	if err != nil {
// 		errResponse := response.ClientResponse(http.StatusBadGateway, "address id is given in the wrong format", nil, err.Error())
// 		c.JSON(http.StatusBadGateway, errResponse)
// 		return
// 	}
// 	user_ID,_:=c.Get("user_id")
// 	OrderResponse,err:=usecase.AddOrder(product_id,address_id,user_ID.(int))
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadGateway, "could not order the product", nil, err.Error())
// 		c.JSON(http.StatusBadGateway, errRes)
// 		return
// 	}
// 	successRes := response.ClientResponse(200, "product ordered Successfully", OrderResponse, nil)
// 	c.JSON(200, successRes)

// }

func OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get("user_id")
	userID := id.(int)

	var orderFromCart models.OrderFromCart
	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderSuccessResponse, err := usecase.OrderItemsFromCart(orderFromCart, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)

}
func GetOrderDetails(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, _ := c.Get("user_id")
	userID := id.(int)
	// id:= c.Query("user_id")
	// userID, _ := strconv.Atoi(id)
	fullOrderDetails, err := usecase.GetOrderDetails(userID, page, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	fmt.Println("full order details is ", fullOrderDetails)

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}


func CancelOrder(c *gin.Context) {

	orderID := c.Param("id")

	id, _ := c.Get("user_id")
	userID := id.(int)
	// id:= c.Query("user_id")
	// userID, _ := strconv.Atoi(id)

	err := usecase.CancelOrders(orderID, userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func PlaceOrder(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userId := userID.(int)
	straddress := c.Param("address_id")
	paymentMethod := c.Param("payment")
	addressId, err := strconv.Atoi(straddress)
	fmt.Println("payment is ", paymentMethod, "address is ", addressId)
	if err != nil {

		errorRes := response.ClientResponse(http.StatusInternalServerError, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	if paymentMethod == "1" {

		Invoice, err := usecase.ExecutePurchaseCOD(userId, addressId)
		if err != nil {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "error in making cod ", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}
		successRes := response.ClientResponse(http.StatusOK, "Placed Order with cash on delivery", Invoice, nil)
		c.JSON(http.StatusOK, successRes)
	}
}
