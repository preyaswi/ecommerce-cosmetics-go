package handlers

import (
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

func AddCoupon(c *gin.Context) {

	var coupon models.AddCoupon

	if err := c.ShouldBindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not bind the coupon details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	message, err := usecase.AddCoupon(coupon)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add coupon", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Coupon Added", message, nil)
	c.JSON(http.StatusCreated, successRes)

}
func GetCoupon(c *gin.Context) {

	coupons, err := usecase.GetCoupon()

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not get coupon details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Coupon Retrieved successfully", coupons, nil)
	c.JSON(http.StatusOK, successRes)

}
func ExpireCoupon(c *gin.Context) {

	id := c.Param("id")
	couponID, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coupon id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = usecase.ExpireCoupon(couponID)
	if err != nil {
		///
		if errors.Is(err,errorss.ErrCouponAlreadyexist){
			errorRes := response.ClientResponse(http.StatusForbidden, "could not expire coupon", nil, err.Error())
		c.JSON(http.StatusForbidden, errorRes)
		return
		}
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not expire coupon", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Coupon expired successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func AddProdcutOffer(c *gin.Context) {

	var productOffer models.ProductOfferReceiver

	if err := c.ShouldBindJSON(&productOffer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := validator.New().Struct(productOffer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = usecase.AddProductOffer(productOffer)

	if err != nil {
		if errors.Is(err, errorss.ErrOfferAlreadyexist) {
			errRes := response.ClientResponse(http.StatusForbidden, "could not add offers", nil, err.Error())
			c.JSON(http.StatusForbidden, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add offer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added offer", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}
func AddCategoryOffer(c *gin.Context) {

	var categoryOffer models.CategoryOfferReceiver

	if err := c.ShouldBindJSON(&categoryOffer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := validator.New().Struct(categoryOffer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = usecase.AddCategoryOffer(categoryOffer)

	if err != nil {
		if errors.Is(err, errorss.ErrOfferAlreadyexist) {
			errRes := response.ClientResponse(http.StatusForbidden, "could not add offers", nil, err.Error())
			c.JSON(http.StatusForbidden, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not add offer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added offer", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}
