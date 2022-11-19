package helpers

import (
	"errors"
	"fmt"
	"strings"
)

func CheckVersion(version string) error {
	if version[0:1] != "v" { // make sure have a v at the start v0.1.1
		return errors.New(fmt.Sprintf("incorrect provided: %s version number try: v1.2.3", version))
	}
	p := strings.Split(version, ".")
	if len(p) != 3 {
		return errors.New(fmt.Sprintf("incorrect length provided: %s version number try: v1.2.3", version))
	}
	return nil
}
