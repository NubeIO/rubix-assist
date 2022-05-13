package remote

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/rubix-assist/service/remote/command"
	"strings"
)

type Node struct {
	IsInstalled      bool   `json:"is_installed"`
	InstalledVersion string `json:"installed_version"`
}

func (inst *Admin) NodeGetVersion() (res *command.Response, node *Node) {
	cmd := "nodejs -v"
	inst.SSH.CMD.Commands = command.Builder(cmd)
	res = inst.SSH.RunCommand()
	if res.Err != nil {
		return
	}
	cmdOut := res.Out
	if strings.Contains(cmdOut, "v") {
		node.InstalledVersion = str.RemoveNewLine(cmdOut)
		node.IsInstalled = true
		return res, node
	} else {
		node.InstalledVersion = cmdOut
		node.IsInstalled = false
		return res, node
	}
}

type NodeJSInstall struct {
	AlreadyInstalled bool   `json:"already_installed"`
	InstalledOk      bool   `json:"installed_ok"`
	TextOut          string `json:"text_out"`
}

//func (inst *Admin) InstallNode14() (NodeJSInstall NodeJSInstall, err error) {
//	cmd := "sudo apt update -y \\\n  && curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash - \\\n  && sudo apt-get install -y nodejs \\\n  && nodejs -v"
//	inst.Host.CMD = cmd
//	res = inst.Host.RunCommand()
//	cmdOut := res.Out
//	if err != nil {
//		log.Error("ufw: Install Error: ", err)
//		NodeJSInstall.TextOut = out
//		NodeJSInstall.InstalledOk = ok
//		return NodeJSInstall, err
//	}
//	_, _, err = inst.NodeGetVersion()
//	if err != nil {
//		log.Error("node: NodeGetVersion Error: ", err)
//		NodeJSInstall.TextOut = ""
//		NodeJSInstall.InstalledOk = false
//		return NodeJSInstall, err
//	}
//	NodeJSInstall.AlreadyInstalled = true
//	return NodeJSInstall, err
//
//}
