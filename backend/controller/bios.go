package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/git"
	"github.com/gin-gonic/gin"
)

func (base *Controller) InstallBios(ctx *gin.Context) {
	//mk download dir and or clear
	//download and unzip build into bios
	//unzip bios
	//run install
	body, err := getGitBody(ctx)
	token := resolveHeaderGitToken(ctx)
	host, useID, err := base.resolveHost(ctx)
	debug := false
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	path := fmt.Sprintf("/home/%s/rubix-bios-install", host.Username)
	opts := commandOpts{
		cmd:  path,
		host: *host,
	}
	//MAKE DIR if not existing and also clear dir
	fmt.Println("mk dir", "try and make dir")
	_, err = base.mkDir(host, opts.cmd, true, true)
	if err != nil {
		fmt.Println("mk dir", "mk dir fail")
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println("mk dir", "mk dir pass")

	opts = commandOpts{
		cmd:   "dpkg --print-architecture",
		debug: debug,
		host:  *host,
	}

	//install BUILD
	getArch, _, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println("ARCH", string(getArch))

	g := body
	g.Token = token
	g.DownloadPath = path
	command := g.BuildCURL(git.CurlReleaseDownload)
	hostName := host.Name
	if useID {
		hostName = host.ID
	}
	opts = commandOpts{
		id:    hostName,
		cmd:   command,
		debug: debug,
		host:  *host,
	}
	//DOWNLOAD BUILD
	fmt.Println("download", "try and download bios")
	_, download, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println("download", download)
	unzipCmd := "unzip -o " + g.DownloadPath + "/" + g.FolderName + " -d " + g.DownloadPath
	fmt.Println(unzipCmd)

	opts = commandOpts{
		id:    hostName,
		cmd:   unzipCmd,
		debug: debug,
		host:  *host,
	}
	//UNZIP BUILD
	_, unzip, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println("unzip", unzip)

	opts = commandOpts{
		id:    hostName,
		cmd:   fmt.Sprintf("cd %s; sudo ./rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --install --auth --device-type %s --token %s", g.DownloadPath, g.Target, token),
		debug: debug,
		host:  *host,
	}

	//install BUILD
	_, install, err := base.runCommand(opts, host.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println("install", install)

	reposeHandler("[ass]", err, ctx)

}

//func (base *Controller) InstallBios(ctx *gin.Context) {
//	getConfig := config.GetConfig()
//	body := uploadBody(ctx)
//	toPath := body.ToPath
//	if body.ToPath == "" {
//		toPath = getConfig.Path.ToPath
//	}
//	fmt.Println(toPath)
//	id := ctx.Params.ByName("id")
//	d, err := base.GetHostDB(id)
//	if err != nil {
//		reposeHandler(d, err, ctx)
//	} else {
//		c, _ := base.newClient(id)
//		defer c.Close()
//
//		commands := []string{"sudo rm -r /data",
//			"sudo rm -r /data/rubix-bios/apps/install"}
//
//		for _, command := range commands {
//			out, err := c.Run(command)
//			if err != nil {
//				fmt.Println(err)
//			} else {
//				fmt.Println(out)
//			}
//		}
//
//		reposeHandler("string(out)", err, ctx)
//	}
//
//}
