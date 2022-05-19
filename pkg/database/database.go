package database

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/NubeIO/rubix-assist/model"

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
	dbName := viper.GetString("database.name")
	driver := viper.GetString("database.driver")
	logMode := viper.GetBool("database.logmode")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	//home, err := homedir.Dir()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//dbName := fmt.Sprintf("%s/%s/rubix_updater.db", home, "/.updater")

	newDBLogger := logger.New(
		log.New(getWriter(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  loglevel,    // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	if driver == "" {
		driver = "sqlite"
	}

	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{Logger: newDBLogger})
	default:
		return errors.New("unsupported database driver")
	}

	if err != nil {
		DBErr = err
		return err
	}

	// Auto migrate project models
	err = db.AutoMigrate(
		&model.Host{},
		&model.Token{},
		&model.User{},
		&model.Team{},
		&model.Alert{},
		&model.Message{})
	if err != nil {
		return err
	}
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
	file, err := os.OpenFile("rubix.db.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}
