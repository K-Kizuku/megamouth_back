package db

import (
	"fmt"
	"megamouth/api/entity/models"
	"megamouth/api/utils/codes"
	"megamouth/api/utils/config"
	"megamouth/api/utils/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (conn *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost,
		config.DbUser,
		config.DbPass,
		config.DbName,
		config.DbPort,
	)
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(codes.CodeDatabase, "postgres connection error")
	}

	if err = conn.AutoMigrate(&models.User{}, &models.Image{}, &models.Post{}, &models.Reaction{}, &models.Authentication{}); err != nil {
		return nil, errors.New(codes.CodeDatabase, "faild migrate")
	}

	return conn, nil
}
