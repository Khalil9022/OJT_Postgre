package api

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/khalil9022/OJT_Postgre/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Open DB Root only for creating the intended DB
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database : %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database %w", err)
	}

	if err := db.AutoMigrate(&models.Staging_Customers{}, &models.Branch_Tabs{}, &models.Customer_Data_Tabs{}, &models.Loan_Data_Tabs{}, &models.Mst_Company_Tabs{}, &models.Skala_Rental_Tabs{}, &models.Staging_Errors{}, &models.Vehicle_Data_Tabs{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database : %w", err)
	}

	return db, err
}
