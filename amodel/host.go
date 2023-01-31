package amodel

import (
	"github.com/NubeIO/lib-schema/schema"
)

type Host struct {
	UUID           string `json:"uuid" gorm:"primaryKey"`
	GlobalUUID     string `json:"global_uuid"`
	Name           string `json:"name"  gorm:"type:varchar(255);not null;uniqueIndex:idx_hosts_name_network_uuid"`
	NetworkUUID    string `json:"network_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null;uniqueIndex:idx_hosts_name_network_uuid"`
	Enable         *bool  `json:"enable"`
	Description    string `json:"description"`
	IP             string `json:"ip"`
	BiosPort       int    `json:"bios_port"`
	Port           int    `json:"port"`
	HTTPS          *bool  `json:"https"`
	IsOnline       *bool  `json:"is_online"`
	IsValidToken   *bool  `json:"is_valid_token"`
	ExternalToken  string `json:"external_token"`
	VirtualIP      string `json:"virtual_ip"`
	ReceivedBytes  int    `json:"received_bytes"`
	SentBytes      int    `json:"sent_bytes"`
	ConnectedSince string `json:"connected_since"`
}

type NetworkUUID struct {
	Type     string `json:"type" default:"string"`
	Title    string `json:"title" default:"uuid"`
	ReadOnly bool   `json:"readOnly" default:"true"`
}

type SSHUsername struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"ssh username"`
	Min     int    `json:"minLength" default:"1"`
	Max     int    `json:"maxLength" default:"50"`
	Default string `json:"default" default:"admin"`
}

type SSHPassword struct {
	Type  string `json:"type" default:"string"`
	Title string `json:"title" default:"ssh password"`
}

type SSHPort struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"rubix port"`
	Min     int    `json:"minLength" default:"2"`
	Max     int    `json:"maxLength" default:"65535"`
	Default int    `json:"default" default:"22"`
	Help    string `json:"help" default:"ip port, eg port 8080 192.168.15.10:8080"`
}

type HostProperties struct {
	Name          schema.Name        `json:"name"`
	Enable        schema.Enable      `json:"enable"`
	Description   schema.Description `json:"description"`
	IP            schema.Host        `json:"ip"`
	BiosPort      schema.Port        `json:"bios_port"`
	Port          schema.Port        `json:"port"`
	HTTPS         schema.HTTPS       `json:"https"`
	ExternalToken schema.Token       `json:"external_token"`
}

func GetHostProperties() *HostProperties {
	m := &HostProperties{}
	m.Name.Min = 0
	m.IP.Default = "0.0.0.0"
	m.BiosPort.Title = "bios port"
	m.BiosPort.Default = 1659
	m.BiosPort.ReadOnly = true
	m.Port.Default = 1661
	m.Port.ReadOnly = true
	m.ExternalToken.Title = "external token"
	schema.Set(m)
	return m
}

type HostSchema struct {
	Required   []string        `json:"required"`
	Properties *HostProperties `json:"properties"`
}

func GetHostSchema() *HostSchema {
	m := &HostSchema{
		Required:   []string{"ip", "bios_port", "port"},
		Properties: GetHostProperties(),
	}
	return m
}
