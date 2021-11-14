package database

import (
	"errors"
	"github.com/NubeIO/rubix-updater/model"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup() error {
	var db = DB

	driver := viper.GetString("database.driver")
	logmode := viper.GetBool("database.logmode")

	loglevel := logger.Silent
	if logmode {
		loglevel = logger.Info
	}

	newDBLogger := logger.New(
		log.New(getWriter(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loglevel,    // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open("rubix_updater.db"), &gorm.Config{Logger: newDBLogger})
	default:
		return errors.New("unsupported database driver")
	}

	if err != nil {
		DBErr = err
		return err
	}

	// Auto migrate project models
	db.AutoMigrate(&model.Host{}, &model.Token{})
	DB = db

	return nil
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}

func getWriter() io.Writer {
	file, err := os.OpenFile("ugin.db.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}
