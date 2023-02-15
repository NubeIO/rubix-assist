package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
)

func (inst *DB) GetSnapshotRestoreLogs() ([]*amodel.SnapshotRestoreLog, error) {
	var m []*amodel.SnapshotRestoreLog
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (inst *DB) CreateSnapshotRestoreLog(body *amodel.SnapshotRestoreLog) (*amodel.SnapshotRestoreLog, error) {
	body.UUID = uuid.ShortUUID("srl")
	err := inst.DB.Create(&body).Error
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (inst *DB) UpdateSnapshotRestoreLog(uuid string, host *amodel.SnapshotRestoreLog) (*amodel.SnapshotRestoreLog, error) {
	m := amodel.SnapshotRestoreLog{}
	err := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(host).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (inst *DB) DeleteSnapshotRestoreLog(uuid string) (*DeleteMessage, error) {
	var m *amodel.SnapshotRestoreLog
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}
