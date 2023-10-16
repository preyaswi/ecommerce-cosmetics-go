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

	r.POST("/add-to-cart/:id", middleware.AuthMiddleware(), handlers.AddToCart)
	r.DELETE("/remove-from-cart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)
	r.GET("/display-cart", middleware.AuthMiddleware(), handlers.DisplayCart)
	r.DELETE("/empty-cart", middleware.AuthMiddleware(), handlers.EmptyCart)

	r.GET("/orders/:page", middleware.AuthMiddleware(), handlers.GetOrderDetails)
	r.PUT("/cancel-orders/:id", middleware.AuthMiddleware(), handlers.CancelOrder)
	r.GET("/checkout", middleware.AuthMiddleware(), handlers.CheckOut)
	r.GET("/place-order/:address_id/:payment", middleware.AuthMiddleware(), handlers.PlaceOrder)

	//ADMIN LOGIN
	r.POST("/admin-login", handlers.AdminLogin)
	r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBoard)

	r.GET("/get-users", middleware.AuthorizationMiddleware(), handlers.GetUsers)
	r.GET("/get-users/:page", middleware.AuthorizationMiddleware(), handlers.GetUsers)
	r.POST("/get-users/add-users", middleware.AuthorizationMiddleware(), handlers.AddNewUsers)
	r.GET("/get-users/block-users/:id", middleware.AuthorizationMiddleware(), handlers.BlockUser)
	r.GET("/get-users/un-block-users/:id", middleware.AuthorizationMiddleware(), handlers.UnBlockUser)
	// r.GET("/products", handlers.ShowAllProducts)
	// r.POST("/products/add-product", handlers.AddProduct)
	r.POST("/category/add", middleware.AuthorizationMiddleware(), handlers.AddCategory)
	r.PUT("/category/update", middleware.AuthorizationMiddleware(), handlers.UpdateCategory)
	r.DELETE("/category/delete", middleware.AuthorizationMiddleware(), handlers.DeleteCategory)
	r.GET("/approve-order/:order_id", middleware.AuthorizationMiddleware(), handlers.ApproveOrder)
	r.GET("/cancel-order/:order_id", middleware.AuthorizationMiddleware(), handlers.CancelOrderFromAdminSide)

	return r

}
