package edgecli

import (
	"fmt"
	"github.com/NubeIO/lib-date/datelib"
)

type Message struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Time struct {
	*datelib.Time
}

func (inst *Client) GetTime() (*Time, *Message, error) {
	path := fmt.Sprintf("%s/%s", paths.System.path, "time")
	resp, err := inst.Rest.R().
		SetResult(&Time{}).
		SetError(&Message{}).
		Get(path)
	if resp.IsError() || err != nil {
		return nil, resp.Error().(*Message), err
	} else {
		return resp.Result().(*Time), nil, nil
	}
}
