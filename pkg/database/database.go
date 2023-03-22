package database

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"os"
	"path"
)

const (
	username = "admin"
	password = "N00BWires"
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

	if driver == "" {
		driver = "sqlite"
	}
	connection := fmt.Sprintf("%s?_foreign_keys=on", path.Join(config.Config.GetAbsDataDir(), dbName))
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
		&amodel.Location{},
		&amodel.Network{},
		&amodel.Host{},
		&amodel.Token{},
		&amodel.User{},
		&amodel.Team{},
		&amodel.Task{},
		&amodel.Transaction{},
		&amodel.SnapshotLog{},
		&amodel.SnapshotCreateLog{},
		&amodel.SnapshotRestoreLog{})
	if err != nil {
		return err
	}
	DB = db

	user_, _ := user.GetUser()
	if user_ == nil {
		_, _ = user.CreateUser(&user.User{Username: username, Password: password})
	}
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
	fileLocation := fmt.Sprintf("%s/rubix.db.log", config.Config.GetAbsDataDir())
	file, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}
