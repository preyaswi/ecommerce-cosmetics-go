package handlers

import (
	"errors"
	errorss "firstpro/error"
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary  OTP login
// @Description Send OTP to Authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.OTPData true "phone number details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/send-otp [post]
func SendOTP(c *gin.Context) {
	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := usecase.SendOTP(phone.PhoneNumber)

	if err != nil {

		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		fmt.Println()
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Verify OTP
// @Description Verify OTP by passing the OTP in order to authenticate user
// @Tags User OTP Login
// @Accept json
// @Produce json
// @Param phone body models.VerifyData true "Verify OTP Details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/verify-otp [post]
func VerifyOTP(c *gin.Context) {

	var code models.VerifyData

	if err := c.ShouldBindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	users, err := usecase.VerifyOTP(code)

	if err != nil {
		if errors.Is(err, errorss.ErrFailedTovalidateOtp) {
			errorRes := response.ClientResponse(http.StatusForbidden, "failed to verify OTP", nil, err.Error())
			c.JSON(http.StatusForbidden, errorRes)
			return
		}
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}
