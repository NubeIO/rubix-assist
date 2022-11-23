package edgecli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/amodel"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"gopkg.in/yaml.v3"
	"testing"
)

type testYml struct {
	Auth bool `json:"auth" yaml:"auth"`
}

func TestClient_ReadFile(t *testing.T) {
	cli := New(&Client{})
	data, err := cli.ReadFile("/data/flow-framework/config/.env")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func TestClient_ReadFileToYml(t *testing.T) {
	cli := New(&Client{})
	message, err := cli.ReadFile("/data/flow-framework/config/config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(message))
	data := testYml{}
	err = yaml.Unmarshal(message, &data)
	fmt.Println(err)
	pprint.PrintJSON(data)
}

type testJson struct {
	ImageVersion string `json:"image_version"`
	Arch         string `json:"arch"`
	Product      string `json:"product"`
}

func TestClient_WriteFileYml(t *testing.T) {
	data := testYml{
		Auth: false,
	}
	cli := New(&Client{})
	body := &amodel.WriteFile{
		FilePath: "/data/flow-framework/config/config.yml",
		Body:     data,
	}
	message, err := cli.WriteFileYml(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJSON(message)
}

func TestClient_WriteFile(t *testing.T) {
	data := `
PORT=1313
SECRET_KEY=__SECRET_KEY__
`
	cli := New(&Client{})
	body := &amodel.WriteFile{
		FilePath:     "/data/rubix-wires/config/.env",
		BodyAsString: data,
	}
	message, err := cli.WriteString(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJSON(message)
}
