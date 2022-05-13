package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestCMD(t *testing.T) {

	str := []string{"/home/pi", "ls"}
	fmt.Println(strings.Join(str[:], " "))
	out := Run(&Opts{SetPath: "/home/aidan", Commands: []string{"ls"}})
	fmt.Println(out.Out)

}
