package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (inst *DB) GetTeam(uuid string) (*amodel.Team, error) {
	m := new(amodel.Team)
	if err := inst.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetTeam error: %v", err)
		return nil, err
	}
	return m, nil
}

func (inst *DB) GetTeams() ([]*amodel.Team, error) {
	var m []*amodel.Team
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateTeam(body *amodel.Team) (*amodel.Team, error) {
	body.UUID = uuid.ShortUUID("tea")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (inst *DB) UpdateTeam(uuid string, Team *amodel.Team) (*amodel.Team, error) {
	m := new(amodel.Team)
	query := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(Team)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Team, query.Error
	}
}

func (inst *DB) DeleteTeam(uuid string) (*DeleteMessage, error) {
	m := new(amodel.Team)
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (inst *DB) DropTeams() (*DeleteMessage, error) {
	var m *amodel.Team
	query := inst.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
