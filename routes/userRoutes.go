package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.POST("/signup", handlers.Signup)
	r.POST("/login-with-password", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	r.GET("/products", handlers.ShowAllProducts)
	r.GET("/products/page/:page", handlers.ShowAllProducts) //TO ARRANGE PAGE WITH COUNT
	r.GET("/products/:id", handlers.ShowIndividualProducts)
	r.POST("/filter", handlers.FilterCategory)

	r.GET("/show-address", middleware.AuthMiddleware(), handlers.GetAllAddress)
	r.POST("/add-address", middleware.AuthMiddleware(), handlers.AddAddress)
	r.GET("/show-user-details", middleware.AuthMiddleware(), handlers.UserDetails)

	r.PATCH("/edit-user-profile", middleware.AuthMiddleware(), handlers.UpdateUserDetails)
	r.POST("/update-password", middleware.AuthMiddleware(), handlers.UpdatePassword)

	//wishlist
	r.POST("/wish-list/add/:id", middleware.AuthMiddleware(), handlers.AddWishList)
	r.GET("/wish-list", middleware.AuthMiddleware(), handlers.GetWishList)
	r.DELETE("/wish-list/delete/:id", middleware.AuthMiddleware(), handlers.RemoveFromWishlist)

	//cart
	r.POST("/add-to-cart/:id", middleware.AuthMiddleware(), handlers.AddToCart)
	r.DELETE("/remove-from-cart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)
	r.GET("/display-cart", middleware.AuthMiddleware(), handlers.DisplayCart)
	r.DELETE("/empty-cart", middleware.AuthMiddleware(), handlers.EmptyCart)

	//order
	// r.POST("/orders/add/:id",middleware.AuthMiddleware(),handlers.AddOrder)
	r.POST("/order-from-cart", middleware.AuthMiddleware(), handlers.OrderItemsFromCart)
	r.GET("/orders/:page", middleware.AuthMiddleware(), handlers.GetOrderDetails)
	r.PUT("/cancel-orders/:id", middleware.AuthMiddleware(), handlers.CancelOrder)

	r.GET("/checkout", middleware.AuthMiddleware(), handlers.CheckOut)
	r.GET("/place-order/:address_id/:payment", middleware.AuthMiddleware(), handlers.PlaceOrder)

	//coupon
	r.POST("/coupon/apply", middleware.AuthMiddleware(), handlers.ApplyCoupon)

	//refferal
	r.GET("/referral/apply", middleware.AuthMiddleware(), handlers.ApplyReferral)

	//payment
	r.GET("/payment", handlers.MakePaymentRazorPay)
	r.GET("/payment-success", handlers.VerifyPayment)
	return r

}
