package handlers

import (
	"firstpro/domain"
	"firstpro/usecase"
	"firstpro/utils/response"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllProducts(c *gin.Context) {
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

	count, err := strconv.Atoi(countStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	fmt.Println(page, count, "👌")
	products, err := usecase.ShowAllProducts(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func ShowIndividualProducts(c *gin.Context) {
	idstr := c.Param("id")
	fmt.Println(idstr, "😂")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	}
	product, err := usecase.ShowIndividualProducts(id)
	fmt.Println(product, "😊")
	if err != nil {

		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
	c.JSON(http.StatusOK, successRes)

}
func AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	categoryResponse, err := usecase.AddCategory(category)
	if err!=nil{
		errRes:=response.ClientResponse(http.StatusBadRequest,"could not add the category",nil, err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	
	successRes := response.ClientResponse(http.StatusOK, "Successfully added Category", categoryResponse, nil)
	c.JSON(http.StatusOK, successRes)


}
