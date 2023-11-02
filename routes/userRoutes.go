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
			address.POST("/post", handlers.AddAddress)
			address.PUT("put/:id", handlers.UpdateAddress)

		}
		users := r.Group("/users")
		{

			users.GET("/show-user-details", handlers.UserDetails)
			users.PATCH("/edit-user-profile", middleware.AuthMiddleware(), handlers.UpdateUserDetails)
			users.POST("/update-password", middleware.AuthMiddleware(), handlers.UpdatePassword)
		}
		
		//wishlist
		wishlist := r.Group("/wishlist")
		{

			wishlist.POST("/post/:id", middleware.AuthMiddleware(), handlers.AddWishList)
			wishlist.GET("", middleware.AuthMiddleware(), handlers.GetWishList)
			wishlist.DELETE("/delete/:id", middleware.AuthMiddleware(), handlers.RemoveFromWishlist)
		}

		//cart
		cart := r.Group("/cart")
		{
			cart.POST("/post/:id", middleware.AuthMiddleware(), handlers.AddToCart)
			cart.DELETE("/remove-from-cart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)
			cart.GET("", middleware.AuthMiddleware(), handlers.DisplayCart)
			cart.DELETE("", middleware.AuthMiddleware(), handlers.EmptyCart)
		}

		//order
		order := r.Group("/order")
		{

			order.POST("/post", middleware.AuthMiddleware(), handlers.OrderItemsFromCart)
			order.GET("/:page", middleware.AuthMiddleware(), handlers.GetOrderDetails)
			order.PUT("/put/:id", middleware.AuthMiddleware(), handlers.CancelOrder)
		}
		r.GET("/checkout", middleware.AuthMiddleware(), handlers.CheckOut)
		r.GET("/place-order/:address_id/:payment", middleware.AuthMiddleware(), handlers.PlaceOrder)

		//coupon
		r.POST("/coupon/apply", middleware.AuthMiddleware(), handlers.ApplyCoupon)

		//refferal
		r.GET("/referral/apply", middleware.AuthMiddleware(), handlers.ApplyReferral)
	}

	//payment
	r.GET("/payment", handlers.MakePaymentRazorPay)
	r.GET("/payment-success", handlers.VerifyPayment)
	return r

}
