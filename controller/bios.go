package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/rubix-assist/service/installer"
	"github.com/gin-gonic/gin"
)

type InstallResp struct {
	RespDownload *installer.RespDownload `json:"response_download"`
	RespBuilder  *installer.RespBuilder  `json:"response_builder"`
	RespInstall  *installer.RespInstall  `json:"response_install"`
}

func (inst *Controller) InstallBios(ctx *gin.Context) {
	body, err := getAppsInstallBody(ctx)
	token := resolveHeaderGitToken(ctx)
	host, _, err := inst.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	fmt.Println(body)
	fmt.Println(token)

	if nils.BoolIsNil(host.IsLocalhost) {
		body.Token = token
		newInstall := installer.New(body)
		resp := &InstallResp{}

		//DOWNLOAD
		download, err := newInstall.Download()
		resp.RespDownload = download
		inst.publishMSG(&WsMsg{Topic: "app-download", Message: resp})
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}
		//Build service file
		build, err := newInstall.Build()
		resp.RespBuilder = build
		inst.publishMSG(&WsMsg{Topic: "app-build", Message: resp})
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}
		//Install
		install, err := newInstall.Install("nubeio-rubix-bios", "/data/nubeio-rubix-bios.service")
		resp.RespInstall = install
		inst.publishMSG(&WsMsg{Topic: "app-install", Message: resp})
		fmt.Println(err, "Install")
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}

		reposeHandler(resp, nil, ctx)
		return

	} else {
		reposeHandler(nil, errors.New("host must be localhost"), ctx)
		return
	}

}
