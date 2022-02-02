package main

import (
	"fmt"
	netval "github.com/THREATINT/go-net"
	"github.com/brotherpowers/ipsubnet"
	"github.com/jordan-wright/email"
	"github.com/mcnijman/go-emailaddress"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"net"
)

func VerifyHost(host string, remote net.Addr, key ssh.PublicKey) error {

	//
	// If you want to connect to new hosts.
	// here your should check new connections public keys
	// if the key not trusted you shuld return an error
	//

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

	//// Ask user to check if he trust the host public key.
	//if askIsHostTrusted(host, key) == false {
	//
	//	// Make sure to return error on non trusted keys.
	//	return errors.New("you typed no, aborted!")
	//}

	// Add the new host to known hosts file.
	return goph.AddKnownHost(host, remote, key, "")
}

func main() {
	//github.com/jordan-wright/email
	e := email.NewEmail()
	e.From = "Nube Alerts <apick1066@gmail.com>"
	e.To = []string{"ap@nube-io.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "apick1066@gmail.com", "Apphp2508!!!", "smtp.gmail.com"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	ip := "192.168.112.999"
	web := "nube-io.com.au"
	sub := ipsubnet.SubnetCalculator(ip, 24)
	fmt.Println(sub.GetIPAddress())
	fmt.Println(sub.GetSubnetMask()) // 255.255.254.0
	fmt.Println(netval.IsIPAddr(ip))
	fmt.Println(netval.IsURL(web))
	_, err := emailaddress.Parse("1a-foobar.com")
	if err != nil {

	} else {

	}

	client2, err := goph.NewConn(&goph.Config{
		User:     "debian",
		Addr:     "120.151.62.75",
		Port:     2221,
		Auth:     goph.Password("N00BConnect"),
		Callback: VerifyHost,
	})

	//client2, err := goph.New("debian", "120.151.62.75:2221", goph.Password("N00BConnect"))
	if err != nil {
		fmt.Println(err)
	}
	cmd, err := client2.Command("pwd")
	fmt.Println(cmd.String())
}
