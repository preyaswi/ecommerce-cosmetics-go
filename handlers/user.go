package handlers

import (
	"context"
	"errors"
	errorss "firstpro/error"
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary SignUp functionality for user
// @Description SignUp functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.SignupDetail true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/signup [post]
func Signup(c *gin.Context) {
	var userSignup models.SignupDetail
	if err := c.ShouldBindJSON(&userSignup); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// checking whether the data sent by the user has all the correct constraints specified by Users struct
	err := validator.New().Struct(userSignup)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//creating a newuser signup with the given deatil passing into the bussiness logic layer

	userCreated, err := usecase.UserSignup(userSignup)
	if err != nil {
		if errors.Is(err, errorss.ErrEmailAlreadyExist) {
			errRes := response.ClientResponse(http.StatusForbidden, "email already exist", nil, err.Error())
			c.JSON(http.StatusForbidden, errRes)
			return
		}
		if errors.Is(err, errorss.ErrPhoneAlreadyExist) {
			errRes := response.ClientResponse(http.StatusForbidden, "phonenumber already exist", nil, err.Error())
			c.JSON(http.StatusForbidden, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong formaaaaat", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)
}

// @Summary LogIn functionality for user
// @Description LogIn functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.LoginDetail true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/login-with-password [post]
func UserLoginWithPassword(c *gin.Context) {
	var userLoginDetail models.LoginDetail
	if err := c.ShouldBindJSON(&userLoginDetail); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := validator.New().Struct(userLoginDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userLoggedInWithPassword, err := usecase.UserLoginWithPassword(userLoginDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully Logged In With password", userLoggedInWithPassword, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Get all address for the user
// @Description Display all the added user addresses
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/address [get]
func GetAllAddress(c *gin.Context) {

	userID, _ := c.Get("user_id")

	// userId := c.Param("user_id")
	// userID, err := strconv.Atoi(userId)

	addressInfo, err := usecase.GetAllAddress(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User Address", addressInfo, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary AddAddress functionality for user
// @Description AddAddress functionality at the user side
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/address [post]
func AddAddress(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var address models.AddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(address)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints does not match", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := usecase.AddAddress(userID.(int), address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in adding address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	SuccessRes := response.ClientResponse(200, "address added successfully", nil, nil)

	c.JSON(200, SuccessRes)

}

// @Summary Update User Address
// @Description Update User address by sending in address id
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "address id"
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/address/{id} [put]
func UpdateAddress(c *gin.Context) {

	id := c.Param("id")
	addressId, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "address id not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")
	user_id := userID.(int)

	var address models.AddressInfo

	if err := c.ShouldBindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	updatedAddress, err := usecase.UpdateAddress(address, addressId, user_id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed updating address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "address updated successfully", updatedAddress, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary User Details
// @Description User Details from User Profile
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users [get]
func UserDetails(c *gin.Context) {

	userID, _ := c.Get("user_id")

	userDetails, err := usecase.UserDetails(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user Details", userDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Update User details
// @Description Update User details
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body models.UsersProfileDetails true "User detail update"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users/update-password [put]
func UpdateUserDetails(c *gin.Context) {

	user_id, _ := c.Get("user_id")

	var user models.UsersProfileDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	updatedDetails, err := usecase.UpdateUserDetails(user, user_id.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed update user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Updated User Details", updatedDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Update User Password
// @Description Update User Password
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body models.UpdatePassword true "User Password update"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users/update-password [put]
func UpdatePassword(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", user_id.(int))
	var body models.UpdatePassword

	if err := c.ShouldBindJSON(&body); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := usecase.UpdatePassword(ctx, body)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed updating password", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Password updated successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)
}

// @Summary Checkout Order
// @Description Checkout at the user side
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/checkout [get]
func CheckOut(c *gin.Context) {

	userID, _ := c.Get("user_id")
	checkoutDetails, err := usecase.Checkout(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Checkout Page loaded successfully", checkoutDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/wishlist/{id} [get]
func AddWishList(c *gin.Context) {
	userID, _ := c.Get("user_id")
	productId := c.Param("id")
	product_id, err := strconv.Atoi(productId)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = usecase.AddToWishlist(product_id, userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to add product on wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully product added to wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Display Wishlist
// @Description Display wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/wishlist [get]
func GetWishList(c *gin.Context) {

	userID, _ := c.Get("user_id")
	wishList, err := usecase.GetWishList(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve wishlist detailss", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "SuccessFully retrieved wishlist", wishList, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/wishlist/{id} [delete]
func RemoveFromWishlist(c *gin.Context) {
	userId, _ := c.Get("user_id")
	productId := c.Param("id")
	product_id, err := strconv.Atoi(productId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = usecase.RemoveFromWishlist(product_id, userId.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product cannot remove from the wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "SuccessFully removed product from the wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Apply referrals
// @Description Apply referrals amount to order
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/referral/apply [get]
func ApplyReferral(c *gin.Context) {

	userID, _ := c.Get("user_id")
	message, err := usecase.ApplyReferral(userID.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add referral amount", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	if message != "" {
		errRes := response.ClientResponse(http.StatusOK, message, nil, nil)
		c.JSON(http.StatusOK, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added referral amount", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
