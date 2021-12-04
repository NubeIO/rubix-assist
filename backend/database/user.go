package dbase

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetUser(id string) (*model.User, error) {
	m := new(model.User)
	if err := d.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetUsers() ([]model.User, error) {
	var m []model.User
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateUser(User *model.User) (*model.User, error) {
	User.ID, _ = uuid.MakeUUID()
	if err := d.DB.Create(&User).Error; err != nil {
		return nil, err
	} else {
		return User, nil
	}
}

func (d *DB) UpdateUser(id string, User *model.User) (*model.User, error) {
	m := new(model.User)
	query := d.DB.Where("id = ?", id).Find(&m).Updates(User)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return User, query.Error
	}
}

func (d *DB) DeleteUser(id string) (ok bool, err error) {
	m := new(model.User)
	query := d.DB.Where("id = ? ", id).Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}

// DropUsers delete all.
func (d *DB) DropUsers() (bool, error) {
	var m *model.User
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}
