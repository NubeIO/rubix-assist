package ip

import (
	"fmt"
	"testing"
)

func TestCheckHTTP(t *testing.T) {

	list := []string{"192.178.13.12", "http://192.178.13.12:8080", "aaaa", "ssss"}

	for _, s := range list {
		a := isValidURL(s)
		fmt.Println(a.IP)
		//pprint.PrintJOSN(a)
	}

}
