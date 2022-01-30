package main

import (
	"fmt"
	netval "github.com/THREATINT/go-net"
	"github.com/brotherpowers/ipsubnet"
	"github.com/jordan-wright/email"
	"github.com/mcnijman/go-emailaddress"
)

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
	_email, err := emailaddress.Parse("1a-foobar.com")
	if err != nil {
		fmt.Println("invalid email")
	} else {
		fmt.Println(_email.LocalPart) // foo
		fmt.Println(_email.Domain)    // bar.com
		fmt.Println(_email)           // foo@bar.com
		fmt.Println(_email.String())  // foo@bar.com
	}

}
