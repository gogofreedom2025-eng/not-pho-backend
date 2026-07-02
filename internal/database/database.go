package database

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/gogo1/not-pho-backend/internal/models"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	path := os.Getenv("DATABASE_PATH")
	if path == "" {
		path = "dev.db"
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.AutoMigrate(&models.Beer{}); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return db, nil
}

func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Beer{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count beers: %w", err)
	}
	if count > 0 {
		return nil
	}

	beers := []models.Beer{
		{Name: "Natty Light", Brewery: "Anheuser-Busch", Style: "American Light Lager", ABV: 4.2, Description: "The classic college fridge staple"},
		{Name: "Bud Select", Brewery: "Anheuser-Busch", Style: "American Light Lager", ABV: 4.3, Description: "Smooth, easy-drinking light beer"},
		{Name: "Guinness Draught", Brewery: "Guinness", Style: "Irish Stout", ABV: 4.2, Description: "Nitrogen-poured dry stout"},
		{Name: "Blue Moon", Brewery: "Blue Moon", Style: "Belgian White", ABV: 5.4, Description: "Wheat beer with orange peel and coriander"},
	}

	if err := db.Create(&beers).Error; err != nil {
		return fmt.Errorf("seed beers: %w", err)
	}

	log.Printf("seeded %d beers", len(beers))
	return nil
}
