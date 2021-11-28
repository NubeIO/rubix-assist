package main

import (

	"github.com/NubeIO/rubix-updater/utils/git"

)


func main()  {

	s := &git.Git{
		Owner: "NubeIO",
		Repo: "flow-framework",
	}
	s.GetBuilds()

}
