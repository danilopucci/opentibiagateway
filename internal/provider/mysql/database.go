package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB creates and returns a GORM DB instance
func NewMySqlDatabase(dsn string) (*gorm.DB, error) {

	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, // Change to Silent/Info/Warn/Error as needed
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Ping test using generic DB interface
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get raw DB: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return db, nil
}

func GenerateDsnFromEnv() string {
	user := os.Getenv("MYSQL_DATABASE_USER")
	password := os.Getenv("MYSQL_DATABASE_PASSWORD")
	host := os.Getenv("MYSQL_DATABASE_HOST")
	port := os.Getenv("MYSQL_DATABASE_PORT")
	databaseName := os.Getenv("MYSQL_DATABASE_NAME")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		databaseName,
	)
}
