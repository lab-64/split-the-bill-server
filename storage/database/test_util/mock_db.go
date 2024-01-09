package test_util

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMockDB initializes a mock database and returns the sql.DB, gorm.DB and sqlmock.Sqlmock instances.
func InitMockDB() (*sql.DB, *gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, dbMock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return sqlDB, gormDB, dbMock, nil
}
