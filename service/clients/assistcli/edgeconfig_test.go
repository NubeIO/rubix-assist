package assistcli

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestClient_EdgeReadConfig(t *testing.T) {
	resp, err := client.EdgeReadConfig("rc", "flow-framework", "config.yml")
	dataFile := testYml{}
	err = yaml.Unmarshal(resp.Data, &dataFile)
	fmt.Println(err)
	pprint.PrintJSON(dataFile)
}
