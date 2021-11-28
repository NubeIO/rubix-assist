package git

import (
	"bufio"
	"fmt"
	"github.com/NubeIO/rubix-updater/utils/command"
	"strings"
)

var Targets = struct {
	Amd64 string `json:"amd64"`
	Armv7 string `json:"armv7"`
}{
	Amd64: "amd64",
	Armv7: "armv7",
}

type Git struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Zip    string `json:"zip"`
	Target string `json:"target"`
	Token  string `json:"token"`
}

func split(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

const (
	urlBase = iota
	urlReleasesTags
	urlReleases
	urlReleasesLatest

	curlReleasesTags
	curlReleases
	curlReleasesLatest

)



func (g *Git) buildURL(url int) string {
	base := fmt.Sprintf("https://api.github.com/repos/%s/%s", g.Owner, g.Repo)
	switch url {
	case urlBase:
		return base
	case urlReleasesTags:
		return fmt.Sprintf("%s/tags", base)
	case urlReleases:
		return fmt.Sprintf("%s/releases", base)
	case urlReleasesLatest:
		return fmt.Sprintf("%s/releases/latest", base)
	default:
		return base
	}
}


func (g *Git) buildCURL (selection int) string {
	switch selection {
	case curlReleasesTags:
		return fmt.Sprintf("curl -s %s | grep -E 'name' | cut -d '\"' -f 4 -",g.buildURL(urlReleasesTags))
	case curlReleases:
		//curl -s https://api.github.com/repos/NubeIO/flow-framework/releases/latest | grep -E 'browser_download_url' | grep amd64  | cut -d '"' -f 4 -
		return fmt.Sprintf("curl -s %s | grep -E 'name' | cut -d '\"' -f 4 -",g.buildURL(urlReleasesTags))
	case curlReleasesLatest:
		return fmt.Sprintf("curl -s %s | grep -E 'name' | cut -d '\"' -f 4 -",g.buildURL(urlReleasesTags))
	default:
		return ""
	}
}



func (g *Git) cmd() {
	fmt.Println(g.buildCURL(curlReleasesTags))

	cmd, err := command.RunCMD(g.buildCURL(curlReleasesTags), false)
	if err != nil {
		//return
	}
	res := split(string(cmd))
	fmt.Println(res)

}

func (g *Git) GetReleasesTags() ([]string, error) {
	cmd, err := command.RunCMD(g.buildCURL(curlReleasesTags), false)
	if err != nil {
		return nil, nil
	}
	return split(string(cmd)), nil
}

//## download all
//# curl -s https://api.github.com/repos/NubeIO/flow-framework/releases/latest | grep -E 'browser_download_url' | grep amd64 | cut -d '"' -f 4 | wget -qi -
//
//## list
//# curl -s https://api.github.com/repos/NubeIO/flow-framework/releases/latest | grep -E 'browser_download_url' | grep amd64  | cut -d '"' -f 4 -
//
//# download a zip
//# curl -fsSL https://github.com/NubeIO/flow-framework/releases/download/v0.1.8/system-0.1.8-c4f9e7a8.amd64.zip  -O
//
//# download with token
//# curl -fsSL -H 'Authorization: token hp_jCSKmylkV937Vy6aEPyOTZNlHhLGHN0Xzld' https://github.com/NubeIO/flow-framework/releases/download/v0.1.8/system-0.1.8-c4f9e7a8.amd64.zip -O

//## list all releases

