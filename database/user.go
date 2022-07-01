package base

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetUser(uuid string) (*model.User, error) {
	m := new(model.User)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetUsers() ([]*model.User, error) {
	var m []*model.User
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateUser(User *model.User) (*model.User, error) {
	User.UUID, _ = uuid.MakeUUID()
	if err := d.DB.Create(&User).Error; err != nil {
		return nil, err
	} else {
		return User, nil
	}
}

func (d *DB) UpdateUser(uuid string, User *model.User) (*model.User, error) {
	m := new(model.User)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(User)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return User, query.Error
	}
}

func (d *DB) DeleteUser(uuid string) (*DeleteMessage, error) {
	m := new(model.User)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropUsers delete all.
func (d *DB) DropUsers() (*DeleteMessage, error) {
	var m *model.User
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
