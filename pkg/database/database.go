package database

import (
	"errors"
	"fmt"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/pkg/helpers/homedir"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"os"
	"os/user"
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

	currentUser, err := user.Current()
	if currentUser.Username != "root" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("user home:", home)
		//dbName = fmt.Sprintf("%s/%s/rubix_updater.db", home, "/rubix-assist")
		dbName = fmt.Sprintf("rubix_updater.db")
	}

	if driver == "" {
		driver = "sqlite"
	}
	connection := fmt.Sprintf("%s?_foreign_keys=on", dbName)
	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})
	default:
		return errors.New("unsupported database driver")
	}

	if err != nil {
		DBErr = err
		return err
	}

	// Auto migrate project models
	err = db.AutoMigrate(
		&model.Location{},
		&model.Network{},
		&model.Host{},
		&model.Token{},
		&model.User{},
		&model.Team{},
		&model.Task{},
		&model.Transaction{})
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
