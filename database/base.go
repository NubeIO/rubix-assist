package base

import (
	"errors"
	"github.com/NubeIO/rubix-assist/pkg/logger"
	"github.com/NubeIO/rubix-assist/pkg/model"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func (d *DB) GetHostByName(name string, isUUID bool) (*model.Host, error) {
	m := new(model.Host)
	switch isUUID {
	case true:
		if err := d.DB.Where("uuid = ? ", name).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	case false:
		if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
			logger.Errorf("GetHost error: %v", err)
			return nil, err
		}
		return m, nil
	default:
		return nil, errors.New("ERROR no valid uuid or name was provided in the request")
	}
}
