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

	//ADMIN LOGIN
	r.POST("/admin-login", handlers.AdminLogin)
	r.Use(middleware.AuthMIddleware())
	{

		r.GET("/dashboard", handlers.DashBoard)
		r.GET("/get-users", handlers.GetUsers)
		r.GET("get-users/:page", handlers.GetUsers)
		r.POST("get-users/add-users", handlers.AddNewUsers)
		r.GET("/get-users/block-users/:id", handlers.BlockUser)
		r.GET("/get-users/un-block-users/:id", handlers.UnBlockUser)
		// r.GET("/products", handlers.ShowAllProducts)
		// r.POST("/products/add-product", handlers.AddProduct)
		r.POST("/category/add",handlers.AddCategory)
	}

	return r

}
