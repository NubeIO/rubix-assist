package edgecli

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	pprint "github.com/NubeIO/rubix-assist/pkg/helpers/print"
	"testing"
)

func TestClient_ListFiles(t *testing.T) {
	cli := New("", 0)
	apps, err := cli.ListFiles("/data")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)
}

func TestClient_RenameFile(t *testing.T) {
	cli := New("", 0)
	dir, err := fileutils.HomeDir()
	if err != nil {
		return
	}
	old := fmt.Sprintf("%s/test/test.txt", dir)
	newName := fmt.Sprintf("%s/test/test2.txt", dir)
	fmt.Println(old, newName)
	apps, err := cli.RenameFile(old, newName)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)
}

func TestClient_CopyFile(t *testing.T) {
	cli := New("", 0)
	dir, err := fileutils.HomeDir()
	if err != nil {
		return
	}
	old := fmt.Sprintf("%s/test/test2.txt", dir)
	newName := fmt.Sprintf("%s/test/test2/test2.txt", dir)
	fmt.Println(old, newName)
	apps, err := cli.CopyFile(old, newName)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)
}

func TestClient_MoveFile(t *testing.T) {
	cli := New("", 0)
	dir, err := fileutils.HomeDir()
	if err != nil {
		return
	}
	old := fmt.Sprintf("%s/test/test2.txt", dir)
	newName := fmt.Sprintf("%s/test/test2/test2.txt", dir)
	fmt.Println(old, newName)
	apps, err := cli.MoveFile(old, newName)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(apps)
}

func TestClient_DownloadFile(t *testing.T) {
	cli := New("", 0)
	dir, err := fileutils.HomeDir()
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/test", dir)
	fileName := "test.txt"
	dest := "/home/aidan/test/test33.txt"
	message, err := cli.DownloadFile(path, fileName, dest)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(message)

}

func TestClient_DeleteAllFiles(t *testing.T) {
	cli := New("", 0)
	dir, err := fileutils.HomeDir()
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/test", dir)
	message, err := cli.DeleteAllFiles(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	pprint.PrintJOSN(message)

}
