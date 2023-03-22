package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
)

func (inst *DB) GetSnapshotCreateLogs(hostUUID string) ([]*amodel.SnapshotCreateLog, error) {
	var m []*amodel.SnapshotCreateLog
	if err := inst.DB.Where("host_uuid = ?", hostUUID).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) CreateSnapshotCreateLog(body *amodel.SnapshotCreateLog) (*amodel.SnapshotCreateLog, error) {
	body.UUID = uuid.ShortUUID("scl")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func (inst *DB) UpdateSnapshotCreateLog(uuid string, body *amodel.SnapshotCreateLog) (*amodel.SnapshotCreateLog, error) {
	var m *amodel.SnapshotCreateLog
	if err := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(body).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) DeleteSnapshotCreateLog(uuid string) (*DeleteMessage, error) {
	var m *amodel.SnapshotCreateLog
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}
