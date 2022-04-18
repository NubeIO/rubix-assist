package database

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/NubeIO/rubix-assist/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB       *gorm.DB
	err      error
	DBErr    error
	mkdirAll = os.MkdirAll
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup() error {
	conf := config.GetConfig()
	connection := path.Join(path.Join(conf.GetAbsDataDir(), fmt.Sprintf("%s.db", conf.Database.Dbname)))
	createDirectoryIfSqlite(conf.Database.Driver, connection)
	var db = DB

	loglevel := logger.Silent
	if conf.Database.LogMode {
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

	switch conf.Database.Driver {
	case "sqlite":
		_connection := fmt.Sprintf("%s?_foreign_keys=on", connection)
		db, err = gorm.Open(sqlite.Open(_connection), &gorm.Config{
			Logger: newDBLogger,
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		})
	default:
		return errors.New("unsupported database driver")
	}

	if err != nil {
		DBErr = err
		return err
	}

	// Auto migrate project models
	if err := db.AutoMigrate(&model.Host{},
		&model.Token{},
		&model.User{},
		&model.Team{},
		&model.Alert{},
		&model.Message{}); err != nil {
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
	conf := config.GetConfig()
	file, err := os.OpenFile(fmt.Sprintf("%s/rubix.db.log", conf.GetAbsDataDir()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}

func createDirectoryIfSqlite(driver, connection string) {
	if driver == "sqlite" {
		if _, err := os.Stat(filepath.Dir(connection)); os.IsNotExist(err) {
			if err := mkdirAll(filepath.Dir(connection), 0777); err != nil {
				panic(err)
			}
		}
	}
}
