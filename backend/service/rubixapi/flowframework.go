package rubixapi

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/rubix-updater/model"

)



func (a *RestClient) LocalStorage(r Req, update bool, proxy bool) (*model.LocalstorageFlowNetwork, error) {
	r.URL = fmt.Sprintf("/api/localstorage_flow_network")
	if proxy {
		r.URL = fmt.Sprintf("/ff/api/localstorage_flow_network")
	}
	if update {
		r.Method = PATCH
	} else {
		r.Method = GET
	}
	request, err := Request(r)
	if err != nil {
		return nil, err
	}
	s, _ := json.MarshalIndent(r.Body, "", "\t")
	fmt.Print(string(s))
	fmt.Println(request.String())
	res := new(model.LocalstorageFlowNetwork)
	err = json.Unmarshal(request.Bytes(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func urlBuilder(url string, proxy bool, uuid string) (out string) {
	out = ""
	if proxy {
		if uuid != "" {
			out = fmt.Sprintf("/ff/%s/%s", url, uuid)
		} else {
			out = fmt.Sprintf("/ff/%s", url)
		}
	} else {
		if uuid != "" {
			out = fmt.Sprintf("/ff/%s/%s", url, uuid)
		} else {
			out = fmt.Sprintf("/%s/", url)
		}
	}
	return out
}


func (a *RestClient) Points(r Req, method int, proxy bool, uuid string) (interface{}, error) {
	r.URL = urlBuilder("api/points", proxy, uuid)
	r.Method = method
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.RequestBuilder.BaseURL+r.URL)
	s, _ := json.MarshalIndent(r.Body, "", "\t")
	fmt.Println("BODY")
	fmt.Println(string(s))
	request, err := Request(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(request.String())
	var res interface{}
	err = json.Unmarshal(request.Bytes(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
