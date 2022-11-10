package edgebioscli

import (
	"fmt"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli/ebmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"os"
	"path"
	"path/filepath"
)

func (inst *BiosClient) RubixEdgeUpload(body *assistmodel.FileUpload) (*model.Message, error) {
	downloadLocation := fmt.Sprintf("/data/rubix-service/apps/download/rubix-edge/%s", body.Version)
	url := fmt.Sprintf("/api/dirs/create?path=%s", downloadLocation)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))

	url = fmt.Sprintf("/api/files/upload?destination=%s", downloadLocation)
	reader, err := os.Open(body.File)
	if err != nil {
		return nil, err
	}
	resp, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&ebmodel.UploadResponse{}).
		SetFileReader("file", filepath.Base(body.File), reader).
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
