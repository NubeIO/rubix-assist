package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/git"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/dirs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (base *Controller) InstallBios(ctx *gin.Context) {
	//mk download dir and or clear
	//download and unzip build into bios
	//unzip bios
	//run install
	body, err := getGitBody(ctx)
	token := resolveHeaderGitToken(ctx)
	host, _, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	path := fmt.Sprintf("/home/%s/rubix-bios-install", host.Username)
	_host, _ := base.hostCopy(host)
	_dirs := dirs.Dirs{
		Host:          _host,
		Name:          path,
		CheckIfExists: true,
		IfExistsClear: true,
	}

	//MAKE DIR if not existing and also clear dir
	log.Println("mk dir ", "try and make dir")
	_, err = _dirs.MKDir()
	if err != nil {
		fmt.Println("mk dir", "mk dir fail")
		reposeHandler(nil, err, ctx)
		return
	}
	log.Println("mk dir ", "mk dir pass")

	g := body
	g.Token = token
	g.DownloadPath = path
	_dirs.Host.CommandOpts.CMD = g.BuildCURL(git.CurlReleaseDownload)
	log.Println("download ", _dirs.Host.CommandOpts.CMD)
	//DOWNLOAD BUILD
	log.Println("download ", "try and download bios")
	_, download, err := _dirs.Host.RunCommand()
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	log.Println("download ", download)

	//UNZIP BUILD
	_dirs.Host.CommandOpts.CMD = "unzip -o " + g.DownloadPath + "/" + g.FolderName + " -d " + g.DownloadPath
	_, unzip, err := _dirs.Host.RunCommand()
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	log.Println("unzip ", unzip)

	_dirs.Host.CommandOpts.CMD = "rm /data/rubix-service/config/app.json"
	//rm /data/rubix-service/config/apps.json
	_, deleteConfigFile, _ := _dirs.Host.RunCommand()
	log.Println("deleteConfigFile ", deleteConfigFile, _dirs.Host.CommandOpts.CMD)

	_dirs.Host.CommandOpts.CMD = "rm /data/rubix-service/config/apps.json"
	_, deleteConfigFile, _ = _dirs.Host.RunCommand()
	log.Println("deleteConfigFile ", deleteConfigFile, _dirs.Host.CommandOpts.CMD)

	//INSTALL BUILD
	_dirs.Host.CommandOpts.CMD = fmt.Sprintf("cd %s; sudo ./rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --install --auth --device-type %s --token %s", g.DownloadPath, g.Target, token)
	log.Println("cmd ", _dirs.Host.CommandOpts.CMD)
	_, install, err := _dirs.Host.RunCommand()
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	log.Println("install ", install)
	reposeHandler("installed", err, ctx)
	return
}
