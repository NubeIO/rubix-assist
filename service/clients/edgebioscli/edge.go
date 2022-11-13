package edgebioscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-assist/model"
	"github.com/NubeIO/rubix-assist/pkg/assistmodel"
	"github.com/NubeIO/rubix-assist/pkg/global"
	"github.com/NubeIO/rubix-assist/service/clients/edgebioscli/ebmodel"
	"github.com/NubeIO/rubix-assist/service/clients/helpers/nresty"
	"github.com/NubeIO/rubix-assist/service/systemctl"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

const rubixEdgeName = "rubix-edge"

func (inst *BiosClient) RubixEdgeUpload(body *assistmodel.FileUpload) (*model.Message, error) {
	downloadLocation := fmt.Sprintf("/data/rubix-service/apps/download/%s/%s", rubixEdgeName, body.Version)
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

func (inst *BiosClient) RubixEdgeInstall(version string) (*model.Message, error) {
	// delete installed files
	installationDirectory := fmt.Sprintf("/data/rubix-service/apps/install/%s", rubixEdgeName)
	url := fmt.Sprintf("/api/files/delete-all?path=%s", installationDirectory)
	_, _ = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Delete(url))
	log.Println("deleted installed files, if any")

	downloadedFile := fmt.Sprintf("/data/rubix-service/apps/download/%s/%s/app", rubixEdgeName, version)
	installationFile := fmt.Sprintf("/data/rubix-service/apps/install/%s/%s/app", rubixEdgeName, version)

	// create installation directory
	installationDirectoryWithVersion := filepath.Dir(installationFile)
	url = fmt.Sprintf("/api/dirs/create?path=%s", installationDirectoryWithVersion)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	log.Info("created installation directory")

	// move downloaded file to installation directory
	url = fmt.Sprintf("/api/files/move?from=%s&to=%s", downloadedFile, installationFile)
	_, err = nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.Message{}).
		Post(url))
	if err != nil {
		return nil, err
	}
	log.Info("moved downloaded file to installation directory")

	tmpDir, absoluteServiceFileName, err := systemctl.GenerateServiceFile(&systemctl.ServiceFile{
		Name:                        rubixEdgeName,
		Version:                     version,
		ExecStart:                   "app -p 1661 -r /data -a rubix-edge -d data -c config --prod server",
		AttachWorkingDirOnExecStart: true,
	}, global.App)
	if err != nil {
		return nil, err
	}
	log.Info("created service file locally")

	serviceFileName := global.App.GetServiceNameFromAppName(rubixEdgeName)
	if err != nil {
		return nil, err
	}
	const serviceDir = "/lib/systemd/system"
	const serviceDirSoftLink = "/etc/systemd/system/multi-user.target.wants"
	serviceFile := path.Join(serviceDir, serviceFileName)
	symlinkServiceFile := path.Join(serviceDirSoftLink, serviceFileName)
	url = fmt.Sprintf("/api/files/upload?destination=%s", serviceDir)
	reader, err := os.Open(absoluteServiceFileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error open service file: %s err: %s", absoluteServiceFileName, err.Error()))
	}
	if _, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetFileReader("file", serviceFileName, reader).
		SetResult(&ebmodel.UploadResponse{}).
		Post(url)); err != nil {
		return nil, err
	}
	log.Info("service file is uploaded successfully")

	url = fmt.Sprintf("/api/syscall/unlink?path=%s", symlinkServiceFile)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("soft un-linked %s", symlinkServiceFile)

	url = fmt.Sprintf("/api/syscall/link?path=%s&link=%s", serviceFile, symlinkServiceFile)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("soft linked %s to %s", serviceFile, symlinkServiceFile)

	url = "/api/systemctl/daemon-reload"
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("daemon reloaded")

	url = fmt.Sprintf("/api/systemctl/enable?unit=%s", serviceFileName)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("enabled service %s", serviceFileName)

	url = fmt.Sprintf("/api/systemctl/restart?unit=%s", serviceFileName)
	if _, err = nresty.FormatRestyResponse(inst.Rest.R().Post(url)); err != nil {
		log.Error(err)
	}
	log.Infof("started service %s", serviceFileName)

	err = fileutils.RmRF(tmpDir)
	if err != nil {
		log.Errorf("delete tmp generated service file %s", absoluteServiceFileName)
	}
	log.Infof("deleted tmp generated local service file %s", absoluteServiceFileName)

	return &model.Message{Message: "successfully installed the rubix-edge in edge device"}, nil
}
