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

// @Summary Add  a new coupon by Admin
// @Description Add A new Coupon which can be used by the users from the checkout section
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.AddCoupon true "Add new Coupon"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/coupons [post]
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

// @Summary Get coupon details
// @Description Get Available coupon details for admin side
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer [get]
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

// @Summary Expire Coupon
// @Description Expire Coupon by admin which are already present by passing coupon id
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Coupon id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/coupons/expire/{id} [patch]
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
		if errors.Is(err, errorss.ErrCouponAlreadyexist) {
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

// @Summary Add  Product Offer
// @Description Add a new Offer for a product by specifying a limit
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.ProductOfferReceiver true "Add new Product Offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/product-offer [post]
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

// @Summary Add  Category Offer
// @Description Add a new Offer for a Category by specifying a limit
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.CategoryOfferReceiver true "Add new Category Offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/offer/category-offer [post]
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
