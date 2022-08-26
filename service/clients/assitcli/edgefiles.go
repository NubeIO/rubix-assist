package assitcli

//// EdgeWalk list all files/dirs in a dir
//func (inst *Client) EdgeWalk(hostIDName, path string) ([]string, error) {
//	url := fmt.Sprintf("/api/edge/files/walk/?path=%s", path)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&[]string{}).
//		Get(url))
//	if err != nil {
//		return nil, err
//	}
//	return *resp.Result().(*[]string), nil
//}
//
//// EdgeListFiles list all files/dirs in a dir
//func (inst *Client) EdgeListFiles(hostIDName, path string) ([]string, error) {
//	url := fmt.Sprintf("/api/edge/files/list/?path=%s", path)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&[]string{}).
//		Get(url))
//	if err != nil {
//		return nil, err
//	}
//	return *resp.Result().(*[]string), nil
//}
//
//// EdgeRenameFile rename a file
//// use the full name of file and path
//func (inst *Client) EdgeRenameFile(hostIDName, oldNameAndPath, newNameAndPath string) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/rename/?old=%s&new=%s", oldNameAndPath, newNameAndPath)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Post(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeCopyFile copy a file
//// use the full name of file and path
//func (inst *Client) EdgeCopyFile(hostIDName, from, to string) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/copy/?from=%s&to=%s", from, to)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Post(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeMoveFile move a file
//// use the full name of file and path
//func (inst *Client) EdgeMoveFile(hostIDName, from, to string) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/move/?from=%s&to=%s", from, to)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Post(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeDeleteFile delete a file
//// use the full name of file and path
//func (inst *Client) EdgeDeleteFile(hostIDName, path string) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/delete/?path=%s", path)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Delete(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeDeleteAllFiles delete all file's in a dir
//// use the full name of file and path
//func (inst *Client) EdgeDeleteAllFiles(hostIDName, path string) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/delete/all/?path=%s", path)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Delete(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeDeleteDir delete a dir
//// use the full name of file and path
//func (inst *Client) EdgeDeleteDir(hostIDName, path string, recursively bool) (*Message, error) {
//	url := fmt.Sprintf("/api/edge/files/delete/?path=%s&recursively=%s", path, "false")
//	if recursively {
//		url = fmt.Sprintf("/api/edge/files/delete/?path=%s&recursively=%s", path, "true")
//	}
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&Message{}).
//		Delete(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*Message), nil
//}
//
//// EdgeUploadFile upload file
//func (inst *Client) EdgeUploadFile(hostIDName, destination, fileName string, file io.Reader) (*EdgeUploadResponse, error) {
//	url := fmt.Sprintf("/api/edge/files/upload/?destination=%s", destination)
//	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&EdgeUploadResponse{}).
//		SetFileReader("file", fileName, file).
//		Post(url))
//	if err != nil {
//		return nil, err
//	}
//	return resp.Result().(*EdgeUploadResponse), nil
//}
//
//func (inst *Client) EdgeUploadLocalFile(hostIDName, path, fileName, destination string) (*EdgeUploadResponse, error) {
//	fileAndPath := filepath.FromSlash(fmt.Sprintf("%s/%s", path, fileName))
//	reader, err := os.Open(fileAndPath)
//	if err != nil {
//		return nil, errors.New(fmt.Sprintf("error open file:%s err:%s", fileAndPath, err.Error()))
//	}
//	resp, err := inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetResult(&EdgeUploadResponse{}).
//		SetFileReader("file", fileName, reader).
//		Post(fmt.Sprintf("/api/edge/files/upload/?destination=%s", destination))
//	if err != nil {
//		return nil, err
//	}
//	if resp.StatusCode() > 299 {
//		return nil, errors.New(resp.String())
//	}
//	return resp.Result().(*EdgeUploadResponse), nil
//}
//
//// EdgeDownloadFile download a file
//func (inst *Client) EdgeDownloadFile(hostIDName, path, file, destination string) (*DownloadResponse, error) {
//	url := fmt.Sprintf("/api/edge/files/download/?path=%s&file=%s", path, file)
//	_, err := nresty.FormatRestyResponse(inst.Rest.R().
//		SetHeader("host_uuid", hostIDName).
//		SetHeader("host_name", hostIDName).
//		SetOutput(destination).
//		Post(url))
//	if err != nil {
//		return nil, err
//	}
//	return &DownloadResponse{FileName: file, Path: path, Destination: destination}, nil
//}
