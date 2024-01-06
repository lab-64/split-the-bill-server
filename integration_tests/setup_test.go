package integration_tests

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	dbConn *sql.DB
	db     *gorm.DB
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() {
	//dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("USERNAME_TEST")
	password := os.Getenv("PASSWORD_TEST")
	host := os.Getenv("HOST_TEST")
	dbName := os.Getenv("DATABASE_TEST")
	port := os.Getenv("PORT_TEST")

	//dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbName)
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, _ = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	//dbConn, _ = sql.Open(dbdriver, DBURL)
}
