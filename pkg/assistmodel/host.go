package assistmodel

import (
	"github.com/NubeIO/lib-schema/schema"
)

type Host struct {
	UUID          string  `json:"uuid" gorm:"primaryKey" `
	NetworkUUID   string  `json:"network_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	Name          string  `json:"name"  gorm:"type:varchar(255);not null"  `
	Enable        *bool   `json:"enable"`
	Description   *string `json:"description"`
	IP            string  `json:"ip"`
	BiosPort      int     `json:"bios_port"`
	Port          int     `json:"port"`
	HTTPS         *bool   `json:"https"`
	PingEnable    *bool   `json:"ping_enable"`
	PingFrequency int     `json:"ping_frequency"`
	IsOffline     bool    `json:"is_offline"`
	OfflineCount  uint    `json:"offline_count"`
	Message       *string `json:"message"`
	ExternalToken string  `json:"external_token"`
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

type HostSchema struct {
	UUID          schema.UUID        `json:"uuid"`
	NetworkUUID   NetworkUUID        `json:"network_uuid"`
	Name          schema.Name        `json:"name"`
	Enable        schema.Enable      `json:"enable"`
	Description   schema.Description `json:"description"`
	IP            schema.Host        `json:"ip"`
	BiosPort      schema.Port        `json:"bios_port"`
	Port          schema.Port        `json:"port"`
	HTTPS         schema.HTTPS       `json:"https"`
	ExternalToken schema.Token       `json:"external_token"`
	Required      []string           `json:"required"`
}

func GetHostSchema() *HostSchema {
	m := &HostSchema{
		Required: []string{"ip", "port"},
	}
	m.IP.Default = "0.0.0.0"
	m.BiosPort.Title = "bios port"
	m.BiosPort.Default = 1659
	m.BiosPort.ReadOnly = true
	m.Port.Default = 1661
	m.Port.ReadOnly = true
	m.NetworkUUID.Title = "network uuid"
	m.ExternalToken.Title = "external token"
	schema.Set(m)
	return m
}
