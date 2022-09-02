package ip

import (
	"fmt"
	"net"
	"testing"
)

func TestCheckHTTP(t *testing.T) {
	list := []string{"192.178.13.12", "http://192.178.13.12:8080", "aaaa", "ssss"}
	for _, s := range list {
		a := isValidURL(s)
		fmt.Println(a.IP)
	}
	i, err := net.LookupHost("http://0.0.0.0:1662/")
	fmt.Println(i, err)
}
