package datasource

import (
	"fmt"
	"log"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("Successfully Connect to Database")
	Db := &Database{db: db}
	return Db, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func (d *Database) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		log.Fatalf("Failed to close database: %v\n", err)
	}
	sqlDB.Close()
	log.Println("Database connection closed")
}

func (d *Database) Migrate(args ...interface{}) {
	log.Println("Migrate schema...")
	d.db.AutoMigrate(args...)
	log.Println("Migrate Done.")
}
