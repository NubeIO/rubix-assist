package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/git"
	"github.com/gin-gonic/gin"
)

func (base *Controller) GitGetRelease(ctx *gin.Context) {
	body, err := getGitBody(ctx)
	token := resolveHeaderGitToken(ctx)
	host, useID, err := base.resolveHost(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	g := body
	g.Token = token
	command := g.BuildCURL(git.CurlReleasesLatest)
	hostName := host.Name
	if useID {
		hostName = host.UUID
	}
	opts := commandOpts{
		uuid:    hostName,
		cmd:   command,
		debug: true,
		host:  *host,
	}
	runCommand, _, err := base.runCommand(opts, g.IsLocalhost)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	res := g.ResultSplit(string(runCommand))
	reposeHandler(res, err, ctx)

}
