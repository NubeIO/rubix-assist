package assitcli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestClient_WiresUpload(t *testing.T) {
	plan, _ := ioutil.ReadFile("wires-example.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	fmt.Println(err)

	r, _ := client.WiresUpload("hos_7DC861C92923", data)
	fmt.Println(r)
	b, err := client.WiresBackup("rc")

	fmt.Println(b, err)

}

func TestClient_WiresBackup(t *testing.T) {

	b, err := client.WiresBackup("hos_7DC861C92923")

	fmt.Println(b, err)

}
