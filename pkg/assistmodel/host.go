package assistmodel

import (
	"github.com/NubeIO/lib-schema/schema"
	"time"
)

type Host struct {
	UUID                 string    `json:"uuid" gorm:"primaryKey" `
	Name                 string    `json:"name"  gorm:"type:varchar(255);not null"  `
	NetworkUUID          string    `json:"network_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	Enable               *bool     `json:"enable"`
	ProductType          string    `json:"product_type"` //edge28, rubix-compute
	IP                   string    `json:"ip"`
	Port                 int       `json:"port"`
	HTTPS                *bool     `json:"https"`
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	WiresPort            int       `json:"wires_port"`
	SSHPort              int       `json:"ssh_port"`
	SSHUsername          string    `json:"ssh_username"`
	SSHPassword          string    `json:"ssh_password" `
	RubixHTTPS           *bool     `json:"rubix_https" `
	IsLocalhost          *bool     `json:"is_localhost" `
	PingEnable           *bool     `json:"ping_enable"`
	PingFrequency        int       `json:"ping_frequency"`
	IsOffline            bool      `json:"is_offline"`
	OfflineCount         uint      `json:"offline_count"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
	BiosToken            string    `json:"-"`
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
	UUID        schema.UUID        `json:"uuid"`
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
	Enable      schema.Enable      `json:"enable"`
	Product     schema.Product     `json:"product_type"`
	NetworkUUID NetworkUUID        `json:"network_uuid"`
	IP          schema.IP          `json:"ip"`
	Port        schema.Port        `json:"port"`
	HTTPS       schema.HTTPS       `json:"https"`
	Username    schema.Username    `json:"username"`
	Password    schema.Password    `json:"password"`
	SSHPort     SSHPort            `json:"ssh_port"`
	SSHUsername SSHUsername        `json:"ssh_username"`
	SSHPassword SSHPassword        `json:"ssh_password" `
	Required    []string           `json:"required"`
}

func GetHostSchema() *HostSchema {
	m := &HostSchema{
		Required: []string{"name", "ip", "port"},
	}
	m.IP.Default = "0.0.0.0"
	m.NetworkUUID.Title = "network uuid"
	schema.Set(m)
	return m
}
