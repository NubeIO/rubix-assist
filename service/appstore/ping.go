package appstore

import (
	"fmt"
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
