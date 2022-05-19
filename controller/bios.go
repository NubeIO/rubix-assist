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
		download, err := newInstall.Download()

		resp.RespDownload = download
		inst.publishMSG(&WsMsg{Topic: "app-download", Message: resp})
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}
		build, err := newInstall.Build()
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}
		resp.RespBuilder = build
		inst.publishMSG(&WsMsg{Topic: "app-build", Message: resp})
		install, err := newInstall.Install("rubix-bios", "/tmp/nubeio-rubix-bios.service")
		if err != nil {
			reposeWithCode(404, resp, nil, ctx)
			return
		}
		resp.RespInstall = install
		inst.publishMSG(&WsMsg{Topic: "app-install", Message: resp})

		reposeHandler(resp, nil, ctx)
		return

	} else {
		reposeHandler(nil, errors.New("host must be localhost"), ctx)
		return
	}

	return
}
