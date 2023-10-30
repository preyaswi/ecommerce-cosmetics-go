package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {

	//ADMIN LOGIN
	r.POST("/admin-login", handlers.AdminLogin)
	r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBoard)
	r.GET("/sales-report/:period", middleware.AuthorizationMiddleware(), handlers.FilteredSalesReport)

	r.GET("/get-users", middleware.AuthorizationMiddleware(), handlers.GetUsers)
	r.GET("/get-users/:page", middleware.AuthorizationMiddleware(), handlers.GetUsers)
	r.POST("/get-users/add-users", middleware.AuthorizationMiddleware(), handlers.AddNewUsers)
	r.GET("/get-users/block-users/:id", middleware.AuthorizationMiddleware(), handlers.BlockUser)
	r.GET("/get-users/un-block-users/:id", middleware.AuthorizationMiddleware(), handlers.UnBlockUser)
	//products management
	r.POST("/products/add-product", middleware.AuthorizationMiddleware(), handlers.AddProduct)
	r.PUT("/products/update-product", middleware.AuthorizationMiddleware(), handlers.UpdateProduct) //update the product quantity
	r.DELETE("/products/delete-product", middleware.AuthorizationMiddleware(), handlers.DeleteProduct)
	//category management
	r.POST("/category/add", middleware.AuthorizationMiddleware(), handlers.AddCategory)
	r.PUT("/category/update", middleware.AuthorizationMiddleware(), handlers.UpdateCategory)
	r.DELETE("/category/delete", middleware.AuthorizationMiddleware(), handlers.DeleteCategory)
	//order
	r.GET("/approve-order/:order_id", middleware.AuthorizationMiddleware(), handlers.ApproveOrder)
	r.GET("/cancel-order/:order_id", middleware.AuthorizationMiddleware(), handlers.CancelOrderFromAdminSide)
	//image cropping
	r.POST("/image-crop", middleware.AuthorizationMiddleware(), handlers.CropImage)

	//coupon
	r.POST("/offer/coupons/add-coupons", middleware.AuthorizationMiddleware(), handlers.AddCoupon)
	r.GET("/offer/coupons", middleware.AuthorizationMiddleware(), handlers.GetCoupon)
	r.PATCH("/offer/coupons/expire/:id", middleware.AuthorizationMiddleware(), handlers.ExpireCoupon)

	//product and category offer
	r.POST("/offer/product-offer", middleware.AuthorizationMiddleware(), handlers.AddProdcutOffer)
	r.POST("/offer/category-offer",middleware.AuthorizationMiddleware(),handlers.AddCategoryOffer)
	return r

}
