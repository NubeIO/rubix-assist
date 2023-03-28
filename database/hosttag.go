package base

import (
	"github.com/NubeIO/rubix-assist/amodel"
	"gorm.io/gorm/clause"
	"strings"
)

func (inst *DB) UpdateHostTags(hostUUID string, body []*amodel.HostTag) ([]*amodel.HostTag, error) {
	tx := inst.DB.Begin()
	var tags []string
	for i, b := range body {
		body[i].HostUUID = hostUUID
		tags = append(tags, b.Tag)
	}
	notIn := strings.Join(tags, ",")
	if err := tx.Where("host_uuid = ?", hostUUID).Where("tag not in (?)", notIn).Delete(&amodel.HostTag{}).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(body) > 0 {
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(body).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Where("host_uuid = ?", hostUUID).Find(&body).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return body, nil
}
