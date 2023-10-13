package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.POST("/signup", handlers.Signup)
	r.POST("/login-with-password", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	r.GET("/products", handlers.ShowAllProducts)
	r.GET("/products/page/:page", handlers.ShowAllProducts) //TO ARRANGE PAGE WITH COUNT
	r.GET("/products/:id", handlers.ShowIndividualProducts)
	

	
	r.POST("/filter", handlers.FilterCategory)
	r.GET("/showaddress", middleware.AuthMiddleware(), handlers.GetAllAddress)
	r.POST("/add-address", middleware.AuthMiddleware(), handlers.AddAddress)
	r.GET("/show-user-details", middleware.AuthMiddleware(), handlers.UserDetails)
	r.POST("/edit-user-profile", middleware.AuthMiddleware(), handlers.UpdateUserDetails)
	r.POST("/update-password", middleware.AuthMiddleware(), handlers.UpdatePassword)


	r.POST("/addtocart/:id", middleware.AuthMiddleware(), handlers.AddToCart)
	r.DELETE("/removefromcart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)
	r.GET("/displaycart", middleware.AuthMiddleware(), handlers.DisplayCart)
	r.DELETE("/emptycart", middleware.AuthMiddleware(), handlers.EmptyCart)

	r.GET("/orders/:page", middleware.AuthMiddleware(), handlers.GetOrderDetails)
	r.PUT("/cancel-orders/:id", middleware.AuthMiddleware(), handlers.CancelOrder)

	r.GET("/checkout", middleware.AuthMiddleware(), handlers.CheckOut)
	


	//ADMIN LOGIN
	r.POST("/admin-login", handlers.AdminLogin)
	r.Use(middleware.AuthorizationMiddleware())
	{

		r.GET("/dashboard", handlers.DashBoard)
		r.GET("/get-users", handlers.GetUsers)
		r.GET("get-users/:page", handlers.GetUsers)
		r.POST("get-users/add-users", handlers.AddNewUsers)
		r.GET("/get-users/block-users/:id", handlers.BlockUser)
		r.GET("/get-users/un-block-users/:id", handlers.UnBlockUser)
		// r.GET("/products", handlers.ShowAllProducts)
		// r.POST("/products/add-product", handlers.AddProduct)
		r.POST("/category/add", handlers.AddCategory)
		r.PUT("/category/update", handlers.UpdateCategory)
		r.DELETE("/category/delete", handlers.DeleteCategory)
		r.GET("/approve-order/:order_id", handlers.ApproveOrder)
		r.GET("/cancel-order/:order_id",  handlers.CancelOrderFromAdminSide)

	}

	return r

}
