package repositories

import (
	"fmt"
	"os"

	"github.com/PontakornDev/ginAuth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// config := utils.ReadConfigs()
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQLDB")), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Database is Connected")
	DB = db
	return err
}

func MigrationDB() {
	// DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Users{})
}
