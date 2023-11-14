package main

import (
	"firstpro/config"
	database "firstpro/db"
	"firstpro/docs"
	"firstpro/routes"
	"fmt"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   Zog_festiv eCommerce API
// @version  1.0
// @description API for ecommerce website

// @securityDefinitions.apiKey JWT
// @in       header
// @name      token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host   www.zogfestiv.store
// @BasePath  /

// @schemes http
func main() {
	docs.SwaggerInfo.Title = "Zog_festiv"
	docs.SwaggerInfo.Description = "Yo Yo Yo 148 3 to the 3 to the 6 to the 9 "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"https"}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
	}
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	corss := cors.DefaultConfig()
	corss.AllowOrigins = []string{"*"}
	corss.AllowMethods = []string{"GET", "POST", "PUT", "POST"}
	router.Use(cors.New(corss))

	userGroup := router.Group("/user")
	adminGroup := router.Group("/admin")
	routes.Routes(userGroup, db)
	routes.AdminRoutes(adminGroup, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	listenAddr := fmt.Sprintf("%s:%s", cfg.DBPort, cfg.DBHost)
	fmt.Printf("Starting server on %s...\n", cfg.BASE_URL)
	if err := router.Run(cfg.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s: %v", listenAddr, err)
	}
}
