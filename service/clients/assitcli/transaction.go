package assitcli

import (
	"fmt"
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
)

func (inst *Client) GetTransactions() (data []model.Transaction, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Transaction{}).
		Get(path)
	return *resp.Result().(*[]model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) GetTransaction(uuid string) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Transaction{}).
		Get(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) AddTransaction(body *model.Transaction) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Transaction{}).
		Post(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) UpdateTransaction(uuid string, body *model.Transaction) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Transaction{}).
		Patch(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) DeleteTransaction(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
