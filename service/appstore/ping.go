package appstore

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/service/clients/edgecli"
	"net"
	"time"
)

// AssistPing ping from the assist service
func (inst *Store) AssistPing(hostUUID, hostName string) bool {
	host, err := inst.getHost(hostUUID, hostName)
	if err != nil {
		return false
	}
	ip_ := fmt.Sprintf("%s:%d", host.IP, host.Port)
	conn, err := net.DialTimeout("tcp", ip_, 1000*time.Millisecond)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

// EdgePing ping from the edge device
func (inst *Store) EdgePing(hostUUID, hostName string, body *edgecli.PingBody) (bool, error) {
	client, err := inst.getClient(hostUUID, hostName)
	if err != nil {
		return false, err
	}
	return client.Ping(body)
}
