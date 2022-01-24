package dbase

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/NubeIO/rubix-updater/pkg/logger"
)

func (d *DB) GetTeam(uuid string) (*model.Team, error) {
	m := new(model.Team)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetTeam error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetTeams() ([]model.Team, error) {
	var m []model.Team
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateTeam(Team *model.Team) (*model.Team, error) {
	Team.UUID = config.MakeTopicUUID(model.CommonNaming.Team)
	if err := d.DB.Create(&Team).Error; err != nil {
		return nil, err
	} else {
		return Team, nil
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
