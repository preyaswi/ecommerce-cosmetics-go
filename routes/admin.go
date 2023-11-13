package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	//ADMIN LOGIN

	r.POST("/admin-login", handlers.AdminLogin)

	r.Use(middleware.AuthorizationMiddleware())
	{

		r.GET("/dashboard", handlers.DashBoard)
		r.GET("/sales-report/:period", handlers.FilteredSalesReport)
		users := r.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.GET("/:page", handlers.GetUsers)
			users.POST("", handlers.AddNewUsers)
			users.GET("/block-users/:id", handlers.BlockUser)
			users.GET("/un-block-users/:id", handlers.UnBlockUser)
		}

		//products management
		products := r.Group("/products")
		{
			products.POST("", handlers.AddProduct)
			// products.POST("/upload-product-image",handlers.UploadImage)
			products.PUT("", handlers.UpdateProduct) //update the product quantity
			products.DELETE("", handlers.DeleteProduct)

		}
		//category management
		category := r.Group("/category")
		{
			category.POST("", handlers.AddCategory)
			category.PUT("", handlers.UpdateCategory)
			category.DELETE("/:id", handlers.DeleteCategory)

		}

		//order
		order := r.Group("/order")
		{
			order.GET("/approve/:order_id", handlers.ApproveOrder)
			order.GET("/cancel/:order_id", handlers.CancelOrderFromAdminSide)
		}

		//image cropping
		r.POST("/image-crop", handlers.CropImage)

		offer := r.Group("/offer")
		{
			//coupon
			coupons := offer.Group("/coupons")
			{
				coupons.POST("", handlers.AddCoupon)
				coupons.GET("", handlers.GetCoupon)
				coupons.PATCH("/expire/:id", handlers.ExpireCoupon)

			}
			//product and category offer
			offer.POST("/product-offer", handlers.AddProdcutOffer)
			offer.POST("/category-offer", handlers.AddCategoryOffer)
		}
	}
	return r

}
