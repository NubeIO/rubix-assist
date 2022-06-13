package edgeapi

import (
	"github.com/NubeIO/edge/service/client"
	base "github.com/NubeIO/rubix-assist/database"
	"github.com/NubeIO/rubix-assist/service/events"
)

type Manager struct {
	em     *client.Client
	DB     *base.DB
	Events *events.Events
}

func New(apps *Manager) *Manager {
	apps.Events.EventBus.RegisterTopic(automater)
	apps.registerAutomater()
	return apps
}

func (inst *Manager) reset(url string, port int) *client.Client {
	return client.New(url, port)
}

type App struct {
	LocationName string `json:"locationName"`
	NetworkName  string `json:"networkName"`
	HostName     string `json:"hostName"`
	HostUUID     string `json:"hostUUID"`
	AppName      string `json:"appName"`
	Version      string `json:"version"`
}

func (inst *Manager) response() *client.Response {
	return &client.Response{}
}
