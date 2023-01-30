package amodel

import "github.com/NubeIO/lib-schema/schema"

type Network struct {
	UUID         string  `json:"uuid" gorm:"primary_key"`
	Name         string  `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_networks_name_location_uuid"`
	LocationUUID string  `json:"location_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES locations;not null;default:null;uniqueIndex:idx_networks_name_location_uuid"`
	Description  string  `json:"description"`
	Hosts        []*Host `json:"hosts" gorm:"constraint:OnDelete:CASCADE"`
}

type NetworkProperties struct {
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
}

func GetNetworkProperties() *NetworkProperties {
	m := &NetworkProperties{}
	schema.Set(m)
	return m
}

type NetworkSchema struct {
	Required   []string           `json:"required"`
	Properties *NetworkProperties `json:"properties"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{
		Required:   []string{"name"},
		Properties: GetNetworkProperties(),
	}
	return m
}
