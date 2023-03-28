package base

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-assist/amodel"
)

func (inst *DB) CreateHostComment(body *amodel.HostComment) (*amodel.HostComment, error) {
	body.UUID = uuid.ShortUUID("hcm")
	if err := inst.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func (inst *DB) UpdateHostComment(uuid string, body *amodel.HostComment) (*amodel.HostComment, error) {
	m := new(amodel.HostComment)
	err := inst.DB.Where("uuid = ?", uuid).Find(&m).Updates(body).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (inst *DB) DeleteHostComment(uuid string) (*DeleteMessage, error) {
	var m *amodel.HostComment
	query := inst.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}
