package assistapi

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
	"testing"
)

type testYml struct {
	Auth bool `json:"auth" yaml:"auth"`
}

func TestClient_EdgeReadConfig(t *testing.T) {

	cli := NewAuth(&Client{
		rest: &resty.Client{},
		URL:  "0.0.0.0",
		Port: 1662,
	})

	config, err := cli.EdgeReadConfig("rc", "flow-framework", "config.yml")
	if err != nil {
		return
	}
	data := testYml{}
	err = yaml.Unmarshal(config.Data, &data)
	fmt.Println(err)
	pprint.PrintJOSN(data)

}
