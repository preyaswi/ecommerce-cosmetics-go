package database

import (
	"firstpro/config"
	"firstpro/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	pasqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dberr := gorm.Open(postgres.Open(pasqlInfo), &gorm.Config{})
	if dberr != nil {
		return nil, fmt.Errorf("failed  to connect to database:%w", dberr)
	}

	DB = db
	DB.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.RazerPay{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.ProductImages{})
	db.AutoMigrate(&domain.Coupons{})
	db.AutoMigrate(&domain.UsedCoupon{})
	db.AutoMigrate(&domain.OrderCoupon{})
	db.AutoMigrate(&domain.ProductOffer{})
	db.AutoMigrate(&domain.CategoryOffer{})
	db.AutoMigrate(&domain.Referral{})
	db.AutoMigrate(&domain.ProductOfferUsed{})
	db.AutoMigrate(&domain.CategoryOfferUsed{})
	db.AutoMigrate(&domain.WishList{})
	return DB, dberr

}
