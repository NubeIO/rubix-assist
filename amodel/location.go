package amodel

import "github.com/NubeIO/lib-schema/schema"

type Location struct {
	UUID        string     `json:"uuid" gorm:"primary_key"`
	Name        string     `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Description string     `json:"description"`
	Networks    []*Network `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
}

type LocationProperties struct {
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
}

func GetLocationProperties() *LocationProperties {
	m := &LocationProperties{}
	m.Name.Min = 0
	schema.Set(m)
	return m
}

type LocationSchema struct {
	Required   []string            `json:"required"`
	Properties *LocationProperties `json:"properties"`
}

func GetLocationSchema() *LocationSchema {
	m := &LocationSchema{
		Required:   []string{},
		Properties: GetLocationProperties(),
	}
	return m
}
