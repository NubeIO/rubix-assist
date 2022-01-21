package main

import (
	"fmt"
	"github.com/jordan-wright/email"
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

	res := 22 % 20
	fmt.Println(res)

	res = 20 % 20
	fmt.Println(res)

	res = 200 % 10
	fmt.Println(res)

	res = 201 % 20
	fmt.Println(res)
}
