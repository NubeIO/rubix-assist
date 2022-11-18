package base

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (inst *DB) GetUser(uuid string) (*model.User, error) {
	m := new(model.User)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetUsers() ([]*model.User, error) {
	var m []*model.User
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateUser(User *model.User) (*model.User, error) {
	User.UUID, _ = uuid.MakeUUID()
	if err := inst.DB.Create(&User).Error; err != nil {
		return nil, err
	} else {
		return User, nil
	}
}

func (inst *DB) UpdateUser(uuid string, User *model.User) (*model.User, error) {
	m := new(model.User)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(User)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return User, query.Error
	}
}

func (inst *DB) DeleteUser(uuid string) (*DeleteMessage, error) {
	m := new(model.User)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropUsers() (*DeleteMessage, error) {
	var m *model.User
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
