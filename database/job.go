package dbase

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/config"
	"github.com/NubeIO/rubix-assist/pkg/logger"
)

func (d *DB) GetJob(uuid string) (*model.Job, error) {
	m := new(model.Job)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetJobs() ([]model.Job, error) {
	var m []model.Job
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateJob(Job *model.Job) (*model.Job, error) {
	Job.UUID, _ = uuid.MakeUUID()
	Job.UUID = config.MakeTopicUUID(model.CommonNaming.Job)
	if err := d.DB.Create(&Job).Error; err != nil {
		return nil, err
	} else {
		return Job, nil
	}
}

func (d *DB) UpdateJob(uuid string, Job *model.Job) (*model.Job, error) {
	m := new(model.Job)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(Job)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return Job, query.Error
	}
}

func (d *DB) DeleteJob(uuid string) (ok bool, err error) {
	m := new(model.Job)
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

// DropJob delete all.
func (d *DB) DropJob() (bool, error) {
	var m *model.Job
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
