package database

import (
	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv5"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type Database struct {
	Context *gorm.DB
}

func NewDatabase() (*Database, error) {
	d := Database{}
	err := d.Connect()
	return &d, err
}

func (d *Database) Connect() error {

	// Google SQL Connection
	cleanup, err := pgxv5.RegisterDriver(
		"cloudsql-postgres",
		cloudsqlconn.WithLazyRefresh(),
		cloudsqlconn.WithIAMAuthN(),
	)
	if err != nil {
		panic(err)
	}
	// cleanup will stop the driver from retrieving ephemeral certificates
	// Don't call cleanup until you're done with your database connections
	defer cleanup()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PW"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsql-postgres",
		DSN:        dsn,
	}), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// get the underlying *sql.DB type to verify the connection
	sdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	var t time.Time
	if err := sdb.QueryRow("select now()").Scan(&t); err != nil {
		panic(err)
	}

	// set database
	d.Context = db
	return nil
}
