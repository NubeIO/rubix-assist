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

	client := New("0.0.0.0", 1662)

	r, _ := client.WiresUpload("rc", data)
	fmt.Println(r)
	b, err := client.WiresBackup("rc")

	fmt.Println(b, err)

}
