package edgecli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

var jwt = ""
var token = "$2a$10$2y3froQ8HlLacbnIEO..b.Hwx0XmP8fAo9SKcE4eS6fuRQA/h0IVW"

func TestClient_Login(t *testing.T) {
	cli := New(&Client{})
	login, err := cli.Login(&user.User{
		Username: "admin",
		Password: "N00BWires",
	})
	fmt.Println(err)
	pprint.PrintJSON(login)
	if err != nil {
		return
	}
	jwt = login.AccessToken
}

func TestClient_GenerateToken(t *testing.T) {
	cli := New(&Client{})
	jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjM1NzMyMzEsImlhdCI6MTY2MDk4MTIzMSwic3ViIjoiYWRtaW4ifQ.2E71U_B47bphYjB5AZDVDXJr0SYtXPXWVytO4jg6-Wk"
	login, err := cli.GenerateToken(jwt, &TokenCreate{Name: "test", Blocked: nils.NewFalse()})
	fmt.Println(err)
	pprint.PrintJSON(login)
	if err != nil {
		return
	}
}
func TestClient_GetTokens(t *testing.T) {
	cli := New(&Client{})
	cli.SetTokenHeader(token)
	login, err := cli.GetTokens()
	fmt.Println(err)
	pprint.PrintJSON(login)
	if err != nil {
		return
	}
}

func TestClient_GetUser(t *testing.T) {
	cli := New(&Client{})
	cli.SetTokenHeader(token)
	login, err := cli.GetUser()
	pprint.PrintJSON(login)
	if err != nil {
		return
	}
}
