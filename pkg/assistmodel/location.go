package assistmodel

import "github.com/NubeIO/lib-schema/schema"

type Location struct {
	UUID        string     `json:"uuid" gorm:"primary_key"`
	Name        string     `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Description string     `json:"description"`
	Networks    []*Network `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
}

type LocationSchema struct {
	UUID        schema.UUID        `json:"uuid"`
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
}

func GetLocationSchema() *LocationSchema {
	m := &LocationSchema{}
	schema.Set(m)
	return m
}
