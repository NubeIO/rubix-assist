package dbase

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist-model/model"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetTeam(uuid string) (*model.Team, error) {
	m := new(model.Team)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetTeam error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTeams() ([]*model.Team, error) {
	var m []*model.Team
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateTeam(body *model.Team) (*model.Team, error) {
	body.UUID = uuid.ShortUUID("tea")
	if err := d.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (d *DB) UpdateTeam(uuid string, Team *model.Team) (*model.Team, error) {
	m := new(model.Team)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(Team)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Team, query.Error
	}
}

func (d *DB) DeleteTeam(uuid string) (ok bool, err error) {
	m := new(model.Team)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}

// DropTeams delete all.
func (d *DB) DropTeams() (bool, error) {
	var m *model.Team
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
