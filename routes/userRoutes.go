package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.Signup)
	r.POST("/login-with-password", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	products := r.Group("/products")
	{
		products.GET("", handlers.ShowAllProducts)
		products.GET("/page/:page", handlers.ShowAllProducts) //TO ARRANGE PAGE WITH COUNT
		products.GET("/:id", handlers.ShowIndividualProducts)
		products.POST("/filter", handlers.FilterCategory)

	}
	r.Use(middleware.AuthMiddleware())
	{
		address := r.Group("/address")
		{
			address.GET("", handlers.GetAllAddress)
			address.POST("", handlers.AddAddress)
			address.PUT("/:id", handlers.UpdateAddress)

		}
		users := r.Group("/users")
		{

			users.GET("", handlers.UserDetails)
			users.PUT("", handlers.UpdateUserDetails)
			users.PUT("/update-password", handlers.UpdatePassword)
		}

		//wishlist
		wishlist := r.Group("/wishlist")
		{

			wishlist.POST("/:id", handlers.AddWishList)
			wishlist.GET("", handlers.GetWishList)
			wishlist.DELETE("/:id", handlers.RemoveFromWishlist)
		}

		//cart
		cart := r.Group("/cart")
		{
			cart.POST("/:id", handlers.AddToCart)
			cart.DELETE("/:id", handlers.RemoveFromCart)
			cart.GET("", handlers.DisplayCart)
			cart.DELETE("", handlers.EmptyCart)
		}

		//order
		order := r.Group("/order")
		{

			order.POST("", handlers.OrderItemsFromCart)
			order.GET("", handlers.GetOrderDetails)
			order.GET("/:page", handlers.GetOrderDetails)
			order.PUT("/:id", handlers.CancelOrder)
		}
		r.GET("/checkout", handlers.CheckOut)
		r.GET("/place-order/:address_id/:payment", handlers.PlaceOrder)

		//coupon
		r.POST("/coupon/apply", handlers.ApplyCoupon)

		//refferal
		r.GET("/referral/apply", handlers.ApplyReferral)
	}

	//payment
	r.GET("/payment", handlers.MakePaymentRazorPay)
	r.GET("/payment-success", handlers.VerifyPayment)
	return r

}
