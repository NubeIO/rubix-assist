package alerts

import (
	"errors"
)

type AlertStatus string
type AlertEntityType string
type AlertType string

const (
	Active       AlertStatus = "active"
	Acknowledged AlertStatus = "acknowledged"
	Closed       AlertStatus = "closed"
)

const (
	Gateway AlertEntityType = "gateway"
	Network AlertEntityType = "network"
	Device  AlertEntityType = "device"
	Point   AlertEntityType = "point"
	Service AlertEntityType = "service"
)

const (
	Ping      AlertType = "ping"
	Fault     AlertType = "fault"
	Threshold AlertType = "threshold"
	Flatline  AlertType = "flatline"
)

func CheckStatus(s string) error {
	switch AlertStatus(s) {
	case Active:
		return nil
	case Acknowledged:
		return nil
	case Closed:
		return nil
	}
	return errors.New("invalid alert status")
}

func CheckStatusClosed(s string) bool {
	return AlertStatus(s) == Closed
}

func CheckEntityType(s string) error {
	switch AlertEntityType(s) {
	case Gateway:
		return nil
	case Network:
		return nil
	case Device:
		return nil
	case Point:
		return nil
	case Service:
		return nil
	}
	return errors.New("invalid alert entity type")
}

func CheckAlertType(s string) error {
	switch AlertType(s) {
	case Ping:
		return nil
	case Fault:
		return nil
	case Threshold:
		return nil
	case Flatline:
		return nil
	}
	return errors.New("invalid alert type")
}
