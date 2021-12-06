package dbase

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetToken(id string) (*model.Token, error) {
	m := new(model.Token)
	if err := d.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTokens() ([]model.Token, error) {
	var m []model.Token
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateToken(Token *model.Token) (*model.Token, error) {
	Token.ID, _ = uuid.MakeUUID()
	if err := d.DB.Create(&Token).Error; err != nil {
		return nil, err
	} else {
		return Token, nil
	}
}

func (d *DB) UpdateToken(id string, Token *model.Token) (*model.Token, error) {
	m := new(model.Token)
	query := d.DB.Where("id = ?", id).Find(&m).Updates(Token)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Token, query.Error
	}
}

func (d *DB) DeleteToken(id string) (ok bool, err error) {
	m := new(model.Token)
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

// DropTokens delete all.
func (d *DB) DropTokens() (bool, error) {
	var m *model.Token
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
