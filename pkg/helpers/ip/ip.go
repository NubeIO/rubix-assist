package ip

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Valid struct {
	IP          bool  `json:"ip"`
	URL         bool  `json:"url"`
	ValidSchema bool  `json:"valid_schema"`
	LookupHost  bool  `json:"lookup_host"`
	Error       error `json:"error"`
}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}
}

func isValidURL(toCheck string) *Valid {
	resp := &Valid{}
	// Check it's an Absolute URL or absolute path
	uri, err := url.ParseRequestURI(CheckHTTP(toCheck))
	if err != nil {
		resp.Error = err
	} else {
		resp.URL = true
	}
	// Check it's an acceptable scheme
	resp.ValidSchema = true
	switch uri.Scheme {
	case "http":
	case "https":
	default:
		resp.ValidSchema = false
	}
	// Check it's a valid domain name
	_, err = net.LookupHost(uri.Host)
	if err != nil {
		resp.Error = err
		resp.LookupHost = true
	}

	if checkIPAddress(toCheck) {
		resp.IP = true
	}

	return resp
}

func IsHTTPS(isHTTPS bool) string {
	if isHTTPS {
		return "https://"
	} else {
		return "http://"
	}
}

func CheckHTTP(address string) string {
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return "http://" + address
	}
	return address
}

func CheckURL(ip string, port int) error {
	_, err := url.Parse(CheckHTTP(fmt.Sprintf("%s:%d", ip, port)))
	return err
}

func Builder(ip string, port int) (*url.URL, error) {
	return url.ParseRequestURI(CheckHTTP(fmt.Sprintf("%s:%d", ip, port)))
}
