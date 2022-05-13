package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Opts struct {
	ShellToUse string
	SetPath    string
	Commands   []string
}

type Response struct {
	Ok  bool
	Out string
	Err error
}

func Builder(args ...string) []string {
	return args
}

//str := []string{"ls", "/home/pi"}
//newStr := strings.Join(str[:], " ")
//fmt.Println(newStr)
//TODO need to make its easier to build commands as the remote.SSH only takes in one arg as a string

func Run(opts *Opts) (res *Response) {
	res = &Response{}
	if len(opts.Commands) <= 0 {
		res.Err = fmt.Errorf("no command provided")
		return res
	}
	shell := opts.ShellToUse //bash -c, "/usr/bin/ls"
	if shell == "" {
		shell = opts.Commands[0]
	}

	log.Infoln("CMD to run", exec.Command(shell, opts.Commands[1:]...).String())
	cmd := exec.Command(shell, opts.Commands[1:]...)
	cmd.Dir = opts.SetPath
	output, err := cmd.Output()
	outAsString := strings.TrimRight(string(output), "\n")
	if err != nil {
		log.Infoln("cmd", err)
		res.Err = err
		return res
	} else {
		fmt.Println("cmd", outAsString)
	}
	res.Out = outAsString
	res.Ok = true
	return res
}
