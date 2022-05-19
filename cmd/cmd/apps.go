package cmd

import (
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/NubeIO/rubix-assist/service/installer"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "manage rubix service apps",
	Long:  `do things like install an app, the device must have internet access to download the apps`,
	Run:   runApps,
}

func runApps(cmd *cobra.Command, args []string) {

	inst := &installer.Installer{
		Token:    flgApp.token,
		Owner:    flgApp.owner,
		Repo:     flgApp.repo,
		Arch:     flgApp.arch,
		Tag:      flgApp.tag,
		DestPath: flgApp.destPath,
		Target:   flgApp.target,
	}

	install := installer.New(inst)

	downloadInstall := install.DownloadInstall()
	pprint.PrintJOSN(downloadInstall)
	pprint.Print(downloadInstall)

}

var flgApp struct {
	token    string
	owner    string
	repo     string
	arch     string
	tag      string
	destPath string
	target   string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgApp.token, "token", "", "github oauth2 token value (optional)")
	flagSet.StringVarP(&flgApp.owner, "owner", "", "NubeIO", "github repository (OWNER/name)")
	flagSet.StringVarP(&flgApp.repo, "repo", "", "rubix-bios", "github repository (owner/NAME)")
	flagSet.StringVar(&flgApp.tag, "tag", "latest", "version of build")
	flagSet.StringVar(&flgApp.destPath, "dest", "/data", "destination path")
	flagSet.StringVar(&flgApp.target, "target", "", "rename destination file (optional)")
	flagSet.StringVar(&flgApp.arch, "arch", "amd64", "arch keyword")

}
