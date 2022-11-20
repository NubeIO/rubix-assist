package installer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type UploadResponse struct {
	FileName     string `json:"file_name,omitempty"`
	TmpFile      string `json:"tmp_file,omitempty"`
	UploadedFile string `json:"uploaded_file,omitempty"`
}

// Upload upload a build
func (inst *Installer) Upload(zip *multipart.FileHeader) (*UploadResponse, error) {
	tmpDir, err := inst.MakeTmpDirUpload()
	if err != nil {
		return nil, err
	}
	log.Infof("upload build to tmp dir: %s", tmpDir)
	zipSource, err := inst.SaveUploadedFile(zip, tmpDir) // save app in tmp dir
	if err != nil {
		return nil, err
	}
	return &UploadResponse{
		FileName:     zip.Filename,
		TmpFile:      tmpDir,
		UploadedFile: zipSource,
	}, nil
}

// SaveUploadedFile uploads the form file to specific dst.
// combination's of file name and the destination and will save file as: /data/my-file
// returns the filename and path as a string and any error
func (inst *Installer) SaveUploadedFile(file *multipart.FileHeader, destination string) (uploadedFile string, err error) {
	uploadedFile = path.Join(destination, file.Filename)
	fmt.Println("SaveUploadedFile destination", uploadedFile)
	src, err := file.Open()
	if err != nil {
		return uploadedFile, err
	}
	defer src.Close()
	out, err := os.Create(uploadedFile)
	if err != nil {
		return uploadedFile, err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return uploadedFile, err
}
