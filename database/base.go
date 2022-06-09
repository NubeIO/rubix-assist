package base

import (
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}
