package modbus

import (
	"fmt"
	"github.com/grid-x/modbus"
	log "github.com/sirupsen/logrus"
	"time"
)

type RegType uint
type Endianness uint
type WordOrder uint
type Error string

type Serial struct {
	serialPort string // "/dev/ttyUSB0"
	baudRate   int    // 38400
	stopBits   int    // 1
	dataBits   int    // 8
	parity     string // "N"
}

type Client struct {
	Client           modbus.Client
	RTUClientHandler *modbus.RTUClientHandler
	TCPClientHandler *modbus.TCPClientHandler
	Endianness       Endianness
	WordOrder        WordOrder
	RegType          RegType
	DeviceZeroMode   bool
	Debug            bool
	PortUnavailable  bool
	IsSerial         bool
	HostIP           string
	HostPort         int
	Serial           *Serial
}

func setParity(in string) string {
	if in == SerialParity.None {
		return "N"
	} else if in == SerialParity.Odd {
		return "O"
	} else if in == SerialParity.Even {
		return "E"
	} else {
		return "N"
	}
}

func SetClient(inst *Client) (*Client, error) {
	if inst.HostIP == "" {
		inst.HostIP = "192.168.15.202"
	}
	if inst.HostPort == 0 {
		inst.HostPort = 502
	}
	if inst.IsSerial {
		serialPort := "/dev/ttyUSB0"
		baudRate := 38400
		stopBits := 1
		dataBits := 8
		parity := "N"
		handler := modbus.NewRTUClientHandler(serialPort)
		handler.BaudRate = baudRate
		handler.DataBits = dataBits
		handler.Parity = setParity(parity)
		handler.StopBits = stopBits
		handler.Timeout = 5 * time.Second

		err := handler.Connect()
		defer handler.Close()
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("setClient:  %v. port:%s", err, serialPort))
			return nil, err
		}
		mc := modbus.NewClient(handler)
		inst.RTUClientHandler = handler
		inst.Client = mc
		return inst, nil

	} else {
		url, err := JoinIpPort(inst.HostIP, inst.HostPort)
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("modbus: failed to validate device IP %s\n", url))
			return nil, err
		}

		handler := modbus.NewTCPClientHandler(url)
		err = handler.Connect()
		if err != nil {
			modbusErrorMsg(fmt.Sprintf("setClient:  %v. port:%s", err, url))
			return nil, err
		}
		defer handler.Close()
		mc := modbus.NewClient(handler)
		inst.TCPClientHandler = handler
		inst.Client = mc
		return inst, nil
	}
}

func modbusErrorMsg(args ...interface{}) {
	debugMsgEnable := true
	if debugMsgEnable {
		prefix := "Modbus: "
		log.Error(prefix, args)
	}
}

func JoinIpPort(url string, port int) (out string, err error) {
	return fmt.Sprintf("%s:%d", url, port), nil
}

var SerialParity = struct {
	None string `json:"none"`
	Odd  string `json:"odd"`
	Even string `json:"even"`
}{
	None: "none",
	Odd:  "odd",
	Even: "even",
}
