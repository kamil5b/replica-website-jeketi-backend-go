package database

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB
var Redis fiber.Storage
