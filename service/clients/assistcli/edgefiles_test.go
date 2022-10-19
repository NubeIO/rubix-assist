package assistcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"gopkg.in/yaml.v3"
	"testing"
)

type testYml struct {
	Auth bool `json:"auth" yaml:"auth"`
}

func TestClient_EdgeFileExists(t *testing.T) {
	message, err := client.EdgeFileExists("rc", "/data/flow-framework/config/config.yml")
	fmt.Println(message)
	fmt.Println(err)
}

func TestClient_EdgeReadFile(t *testing.T) {
	message, err := client.EdgeReadFile("rc", "/data/flow-framework/config/config.yml")
	dataFile := testYml{}
	err = yaml.Unmarshal(message, &dataFile)
	fmt.Println(err)
	pprint.PrintJSON(dataFile)
}
