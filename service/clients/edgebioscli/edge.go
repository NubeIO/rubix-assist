package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/service/clients/assitcli/nresty"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli/ebmodel"
	"io"
	"path"
)

func (inst *BiosClient) Upload(version, zipFileName string, reader io.Reader) (*model.Message, error) {
	downloadLocation := fmt.Sprintf("/data/rubix-service/apps/download/rubix-edge/%s", version)
	url := fmt.Sprintf("/api/files/upload?destination=%s", downloadLocation)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&ebmodel.UploadResponse{}).
		SetFileReader("file", zipFileName, reader).
		Post(url))
	if err != nil {
		return nil, err
	}
	upload := resp.Result().(*ebmodel.UploadResponse)

	url = fmt.Sprintf("/api/zip/unzip?source=%s&destination=%s", upload.Destination, downloadLocation)
	resp, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]string{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	unzippedFiles := resp.Result().(*[]string)

	url = fmt.Sprintf("/api/files/delete?file=%s", upload.Destination)
	resp, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	if err != nil {
		return nil, err
	}

	for _, f := range *unzippedFiles {
		from := path.Join(downloadLocation, f)
		to := path.Join(downloadLocation, "app")
		url = fmt.Sprintf("/api/files/move?from=%s&to=%s", from, to)
		resp, err = nresty.FormatRestyResponse(inst.Rest.R().
			SetResult(&model.Message{}).
			Post(url))
		if err != nil {
			return nil, err
		}
	}
	return &model.Message{Message: "successfully uploaded the rubix-edge in edge device"}, nil
}
