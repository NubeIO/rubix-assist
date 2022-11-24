package amodel

import "github.com/NubeIO/lib-schema/schema"

type Network struct {
	UUID         string  `json:"uuid" gorm:"primary_key"`
	Name         string  `json:"name"  gorm:"type:varchar(255);not null"`
	LocationUUID string  `json:"location_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES locations;not null;default:null"`
	Hosts        []*Host `json:"hosts" gorm:"constraint:OnDelete:CASCADE"`
}

type LocationUUID struct {
	Type     string `json:"type" default:"string"`
	Title    string `json:"title" default:"uuid"`
	ReadOnly bool   `json:"readOnly" default:"true"`
}

type NetworkSchema struct {
	UUID         schema.UUID        `json:"uuid"`
	Name         schema.Name        `json:"name"`
	Description  schema.Description `json:"description"`
	LocationUUID LocationUUID       `json:"location_uuid"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	m.LocationUUID.Title = "location uuid"
	schema.Set(m)
	return m
}
