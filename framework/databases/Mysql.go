package databases

import (
	"fmt"
	"os"
	"path/filepath"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func loadEnv() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		// Handle the error accordingly
		return
	}

	envPath := filepath.Join(rootDir, "..", ".env")
	fmt.Print(envPath)

	// Memuat variabel lingkungan dari file .env di direktori root
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		// Handle the error accordingly
	}
}

func InitMysql() (*gorm.DB) {
	loadEnv()
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can't connect to database")
	}

	return db
}