package installer

import (
	"context"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"time"
)

type Installer struct {
	Token     string
	Owner     string
	Repo      string
	Asset     string
	Arch      string
	Tag       string
	DestPath  string
	Target    string
	gitClient *git.Client
}

func New(inst *Installer) *Installer {
	opts := &git.AssetOptions{
		Owner:    inst.Owner,
		Repo:     inst.Repo,
		Tag:      inst.Tag,
		Arch:     inst.Arch,
		DestPath: inst.DestPath,
		Target:   inst.Target,
	}
	ctx := context.Background()
	fmt.Println(11111)
	inst.gitClient = git.NewClient(inst.Token, opts, ctx)
	fmt.Println(11111)
	return inst
}

type InstallResp struct {
	GitResp     *git.DownloadResponse `json:"git_resp"`
	GitError    error                 `json:"git_error"`
	BuilderErr  error                 `json:"builder_err"`
	InstallResp *ctl.InstallResp      `json:"install_resp"`
}

func (inst *Installer) DownloadInstall() *InstallResp {
	ret := &InstallResp{}
	fmt.Println(11111)
	//download and unzip to /data
	resp, err := inst.gitClient.DownloadInstall()
	ret.GitResp = resp
	if err != nil {
		ret.GitError = err
		return ret
	}

	newService := "nubeio-rubix-bios"
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/rubix-bios-app"
	execCmd := "/data/rubix-bios-app/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"

	bld := &builder.SystemDBuilder{
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: newService,
			Path:     "/tmp",
		},
	}

	err = bld.Build()
	if err != nil {
		ret.BuilderErr = err
	}

	path := "/tmp/nubeio-rubix-bios.service"

	timeOut := 30
	service := ctl.New(newService, path)
	opts := systemctl.Options{Timeout: timeOut}
	installOpts := ctl.InstallOpts{
		Options: opts,
	}
	service.InstallOpts = installOpts
	installResp := service.Install()
	ret.InstallResp = installResp
	fmt.Println("full install error", err)
	if err != nil {
		fmt.Println("full install error", err)
	}

	time.Sleep(8 * time.Second)

	status, err := systemctl.Status(newService, systemctl.Options{})
	if err != nil {
		log.Errorf("service found: %s: %v", newService, err)
	}
	fmt.Println(status)

	//res, err := service.Remove()
	//fmt.Println("full install error", err)
	//if err != nil {
	//	fmt.Println("full install error", err)
	//}
	//pprint.PrintJOSN(res)

	return ret
}
