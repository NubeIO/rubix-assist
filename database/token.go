package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (inst *DB) GetToken(uuid string) (*amodel.Token, error) {
	m := new(amodel.Token)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetTokens() ([]*amodel.Token, error) {
	var m []*amodel.Token
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateToken(Token *amodel.Token) (*amodel.Token, error) {
	Token.UUID = uuid.ShortUUID("tok")
	if err := inst.DB.Create(&Token).Error; err != nil {
		return nil, err
	} else {
		return Token, nil
	}
}

func (inst *DB) UpdateToken(uuid string, Token *amodel.Token) (*amodel.Token, error) {
	m := new(amodel.Token)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(Token)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Token, query.Error
	}
}

func (inst *DB) DeleteToken(uuid string) (*DeleteMessage, error) {
	m := new(amodel.Token)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropTokens() (*DeleteMessage, error) {
	var m *amodel.Token
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
