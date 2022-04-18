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
	_, err := emailaddress.Parse("1a-foobar.com")
	if err != nil {

	} else {

	}

	//
	//data["Date_1"] = map[string]map[string]string{}
	//data["Date_1"] = make(map[string]map[string]string, 0)
	//data["Date_1"] = make(map[string]map[string]string, 0)
	//
	//data["Date_1"]["Sistem_A"] = map[string]string{}
	//data["Date_1"]["Sistem_A"] = make(map[string]string, 0)
	//data["Date_1"]["Sistem_A"] = make(map[string]string, 0)
	//
	//data["Date_1"]["Sistem_A"]["command_1"] = "white"
	//data["Date_1"]["Sistem_A"]["command_2"] = "blue"
	//data["Date_1"]["Sistem_A"]["command_3"] = "red"
	//
	//data["Date_2"] = make(map[string]map[string]string)
	//data["Date_2"]["Sistem_A"] = make(map[string]string)
	//data["Date_2"]["Sistem_A"]["command_5"] = "violet"
	//
	//fmt.Println("data: ", data)

}
