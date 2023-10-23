package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"fmt"
	"image"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AdminLogin(c *gin.Context) {
	var adminDetail models.AdminDetail
	if err := c.ShouldBindJSON(&adminDetail); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := usecase.AdminLogin(adminDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}
func DashBoard(c *gin.Context) {

	adminDashBoard, err := usecase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed fine", adminDashBoard, nil)
	c.JSON(http.StatusOK, successRes)

}

func GetUsers(c *gin.Context) {

	pageStr := c.Param("page")

	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	countStr := c.Query("count")

	if countStr == "" {
		countStr = "0"
	}

	pageSize, err := strconv.Atoi(countStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := usecase.GetUsers(page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}

func AddNewUsers(c *gin.Context) {
	var newUser models.SignupDetail
	if err := c.ShouldBindJSON(&newUser); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(newUser)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	userCreated, err := usecase.UserSignup(newUser)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not create new user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succRes := response.ClientResponse(http.StatusOK, "new user created", userCreated, nil)
	c.JSON(http.StatusOK, succRes)
}

func BlockUser(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		fmt.Println(err.Error(), "😁")

		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	err = usecase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be blocked", nil, err.Error())
		fmt.Println(err.Error(), "😊")

		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func UnBlockUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	err = usecase.UnBlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldn't blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func ApproveOrder(c *gin.Context) {

	orderId := c.Param("order_id")

	err := usecase.ApproveOrder(orderId)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not approve the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order approved successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func CancelOrderFromAdminSide(c *gin.Context) {

	orderID := c.Param("order_id")

	err := usecase.CancelOrderFromAdminSide(orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func FilteredSalesReport(c *gin.Context) {

	timePeriod := c.Param("period")
	salesReport, err := usecase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, successRes)

}
func ImageCropping(c *gin.Context) {
	inputImage, err := imaging.Open("static/image-crop.jpg")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}
	cropRect := image.Rect(200, 200, 450, 450)
	croppedImage := imaging.Crop(inputImage, cropRect)
	err = imaging.Save(croppedImage, "static/output.jpg")
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save the image"})
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "image resized succesfully", inputImage, nil)
	c.JSON(http.StatusOK, successRes)

}
