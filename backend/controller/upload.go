package controller

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/melbahja/goph"
	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

var (
	err        error
	auth       goph.Auth
	client     *goph.Client
	addr       string
	user       string
	port       uint
	key        string
	cmd        string
	pass       bool
	passphrase bool
	timeout    time.Duration
	agent      bool
	sftpc      *sftp.Client
)

type serverSettings struct {
	Addr, Key, User, Catalog, Password string
	Port                               uint
}

type Git struct {
	Token string `json:"token"`
}

type Upload struct {
	FromPath   string   `json:"from_path"`
	ToPath     string   `json:"to_path"`
	UnZipPath  string   `json:"unzip_path"`
	DeletePath string   `json:"delete_path"`
	ClearDIR   bool     `json:"clear_dir"`
	Unzip      bool     `json:"unzip"`
	Zips       []string `json:"zips"`
}

type Dir struct {
	Path string `json:"path"`
}

func VerifyHost(host string, remote net.Addr, key ssh.PublicKey) error {
	// hostFound: is host in known hosts file.
	// err: error if key not in known hosts file OR host in known hosts file but key changed!
	hostFound, err := goph.CheckKnownHost(host, remote, key, "")
	// Host in known hosts but key mismatch!
	// Maybe because of MAN IN THE MIDDLE ATTACK!
	if hostFound && err != nil {

		return err
	}

	// handshake because public key already exists.
	if hostFound && err == nil {

		return nil
	}
	return goph.AddKnownHost(host, remote, key, "")
}

func (base *Controller) newClient(id string) (c *goph.Client, err error) {
	h, err := base.GetHostDB(id)
	if err != nil {
		return nil, err
	} else {
		var cli serverSettings
		cli.Addr = h.IP
		cli.User = h.Username
		cli.Password = h.Password
		cli.Port = uint(h.Port)
		c, err = goph.NewConn(&goph.Config{
			User:     cli.User,
			Addr:     cli.Addr,
			Port:     cli.Port,
			Auth:     goph.Password(cli.Password),
			Callback: VerifyHost,
		})
		return c, err
	}

}

func (base *Controller) newRemoteClient(host model.Host) (c *goph.Client, err error) {
	var cli serverSettings
	cli.Addr = host.IP
	cli.User = host.Username
	cli.Password = host.Password
	cli.Port = uint(host.Port)
	c, err = goph.NewConn(&goph.Config{
		User:     cli.User,
		Addr:     cli.Addr,
		Port:     cli.Port,
		Auth:     goph.Password(cli.Password),
		Callback: VerifyHost,
	})
	return c, err

}

func getSftp(client *goph.Client) *sftp.Client {
	var err error
	if sftpc == nil {
		sftpc, err = client.NewSftp()
		if err != nil {
			panic(err)
		}
	}
	return sftpc
}

func gitBody(ctx *gin.Context) (dto *Git) {
	err = ctx.ShouldBindJSON(&dto)
	return dto
}

func uploadBody(ctx *gin.Context) (dto *Upload) {
	err = ctx.ShouldBindJSON(&dto)
	return dto
}

func dirBody(ctx *gin.Context) (dto *Dir) {
	err = ctx.ShouldBindJSON(&dto)
	return dto
}

func (base *Controller) uploadZip(id string, body *Upload) error {
	getConfig := config.GetConfig()
	fromPath := body.FromPath
	toPath := body.ToPath
	unZipPath := body.UnZipPath

	if body.FromPath == "" {
		fromPath = getConfig.Path.FromPath
	}
	if body.ToPath == "" {
		toPath = getConfig.Path.ToPath
	}
	if body.UnZipPath == "" {
		unZipPath = getConfig.Path.UnZipPath
	}
	c, _ := base.newClient(id)
	defer c.Close()

	for i, zip := range body.Zips {
		fmt.Println(toPath)
		fmt.Println(fromPath)
		fmt.Println(i, zip)

		fp := fmt.Sprintf("%s/%s", fromPath, zip)
		tp := fmt.Sprintf("%s/%s", toPath, zip)

		err := c.Upload(fp, tp)
		if err != nil {
			log.Println("UPLOAD ZIP ERROR", err)
		}
		command := fmt.Sprintf("sudo unzip %s -d %s", tp, unZipPath)
		out, err := c.Run(command)
		fmt.Println(command)
		if err != nil {
			//log.Fatal(err)
		}
		fmt.Println(string(out))
	}

	return nil
}

func (base *Controller) UploadZip(ctx *gin.Context) {
	body := uploadBody(ctx)
	id := ctx.Params.ByName("id")
	err := base.uploadZip(id, body)
	if err != nil {
		reposeHandler(nil, err, ctx)
	} else {
		reposeHandler("ok", err, ctx)
	}

}

func (base *Controller) Unzip(ctx *gin.Context) {
	getConfig := config.GetConfig()
	body := uploadBody(ctx)
	toPath := body.ToPath
	if body.ToPath == "" {
		toPath = getConfig.Path.ToPath
	}
	id := ctx.Params.ByName("id")
	d, err := base.GetHostDB(id)
	if err != nil {
		reposeHandler(d, err, ctx)
	} else {
		c, _ := base.newClient(id)
		defer c.Close()
		for _, zip := range body.Zips {
			tp := fmt.Sprintf("%s/%s", toPath, zip)
			command := fmt.Sprintf("unzip %s -d %s", tp, toPath)
			out, err := c.Run(command)
			fmt.Println(command)
			if err != nil {
				//log.Fatal(err)
			}
			fmt.Println(string(out))
		}
		reposeHandler("string(out)", err, ctx)
	}

}
