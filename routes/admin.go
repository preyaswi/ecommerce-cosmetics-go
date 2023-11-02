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
			users.POST("/post", handlers.AddNewUsers)
			users.GET("/block-users/:id", handlers.BlockUser)
			users.GET("/un-block-users/:id", handlers.UnBlockUser)
		}

		//products management
		products := r.Group("/products")
		{
			products.POST("/post", handlers.AddProduct)
			products.PUT("/put", handlers.UpdateProduct) //update the product quantity
			products.DELETE("/delete", handlers.DeleteProduct)

		}
		//category management
		category := r.Group("/category")
		{
			category.POST("/post", handlers.AddCategory)
			category.PUT("/update", handlers.UpdateCategory)
			category.DELETE("/delete", handlers.DeleteCategory)

		}

		//order
		order := r.Group("/order")
		{
			order.GET("/approve/:order_id", handlers.ApproveOrder)
			order.GET("/cancel/:order_id", handlers.CancelOrderFromAdminSide)
		}

		//image cropping
		r.POST("/image-crop", middleware.AuthorizationMiddleware(), handlers.CropImage)

		offer := r.Group("/offer")
		{
			//coupon
			coupons := offer.Group("/coupons")
			{
				coupons.POST("/post", handlers.AddCoupon)
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
