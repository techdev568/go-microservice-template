package database

import (
	"fmt"

	"github.com/techdev568/go-microservice-template/internal/config"
	"github.com/techdev568/go-microservice-template/internal/models"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config, log *zap.SugaredLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate tables
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	log.Info("connected to MySQL database")
	return db, nil
}
