package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/service/github"
	"github.com/gin-gonic/gin"
)




func git(token string)  {
	owner := "NubeIO"
	repo := "flow-framework"
	tag := "latest"
	//deviceType := "armv7"
	//dir := "/"
	//pluginsDir := "/data/flow-framework/data/plugins"

	a := github.New()
	var version string
	if tag == "latest" {
		_repo := fmt.Sprintf("%s/%s", owner, repo)
		fmt.Println("repo", _repo)
		tags, err := a.GetLatestReleaseTag(_repo)
		if err != nil {
			fmt.Println("error:", _repo)
		}
		version = tags
	} else {
		version = tag
	}
	fmt.Println("version:", version)
	//downloads, err := github.GetBuilds(owner, repo, version, token)
	//if err != nil {
	//	fmt.Println("error: main github.RetrieveAssets:", err)
	//	//return err
	//}
	//
	//
	////fmt.Println(dir)
	//fmt.Println(downloads)
	//fmt.Println(111)
	//fmt.Println(111)
	//fmt.Println(downloads)

}
func (base *Controller) GitGetRelease(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	getToken, err := base.dbGetToken(id)
	if err != nil {
		return
	}
	fmt.Println(2222, getToken.Token)
	//git(token.Token)

	token := getToken.Token

	owner := "NubeIO"
	repo := "flow-framework"
	tag := "latest"
	//deviceType := "armv7"
	//dir := "/"
	//pluginsDir := "/data/flow-framework/data/plugins"

	a := github.New()
	var version string
	if tag == "latest" {
		_repo := fmt.Sprintf("%s/%s", owner, repo)
		fmt.Println("repo", _repo)
		tags, err := a.GetLatestReleaseTag(_repo)
		if err != nil {
			fmt.Println("error:", _repo)
		}
		version = tags
	} else {
		version = tag
	}
	fmt.Println("version:", version)
	downloads, err := github.GetAssetsInfo(owner, repo, version, token)
	if err != nil {
		fmt.Println("error: main github.RetrieveAssets:", err)
		//return err
	}
	//fmt.Println(dir)
	//fmt.Println(downloads)
	fmt.Println(111)
	fmt.Println(111)
	//downloads.Assets

	reposeHandler(downloads.Assets, err, ctx)
}


