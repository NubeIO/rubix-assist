package base

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
)

func (inst *DB) GetSnapshotLog() ([]*amodel.SnapshotLog, error) {
	var m []*amodel.SnapshotLog
	if err := inst.DB.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) CreateSnapshotLog(body *amodel.SnapshotLog) (*amodel.SnapshotLog, error) {
	fmt.Println("creating...", body)
	if err := inst.DB.Create(&body).Error; err != nil {
		fmt.Println("failed...", body)
		return nil, err
	}
	fmt.Println("created...", body)
	return body, nil
}

func (inst *DB) UpdateSnapshotLog(file string, body *amodel.SnapshotLog) (*amodel.SnapshotLog, error) {
	var logs []*amodel.SnapshotLog
	if err := inst.DB.Where("file = ?", file).Find(&logs).Error; err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		body.File = file
		return inst.CreateSnapshotLog(body)
	}
	var m *amodel.SnapshotLog
	if err := inst.DB.Where("file = ?", file).Find(&m).Updates(body).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) DeleteSnapshotLog(file string) (*DeleteMessage, error) {
	var m *amodel.SnapshotLog
	query := inst.DB.Where("file = ? ", file).Delete(&m)
	return deleteResponse(query)
}

// DeleteSnapshotLogs avoids discrepancies between snapshot raw files and logs
// discrepancies happens when snapshot gets deleted manually
func (inst *DB) DeleteSnapshotLogs(files []string) (*DeleteMessage, error) {
	var m *amodel.SnapshotLog
	query := inst.DB.Where("file NOT IN ?", files).Delete(&m)
	return deleteResponse(query)
}
