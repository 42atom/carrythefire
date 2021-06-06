package remote

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func GetPlots(sshClient *ssh.Client, src string) (map[string]int64, error) {
	client, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}

	plots := map[string]int64{}
	fs, err := client.ReadDir(src)
	if err != nil {
		return nil, err
	}

	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		//Check plot file
		ext := filepath.Ext(f.Name())
		if ext == ".plot" {
			plots[f.Name()] = f.Size()
		}
	}

	return plots, nil
}

func mv(sshClient *ssh.Client, src, filename, dst string, size int64) error {
	csize, err := copy(sshClient, src, filename, dst, size)
	if err != nil {
		return err
	}
	if csize == size {
		err := remove(sshClient, src, filename)
		if err != nil {
			return err
		}
	}
	return err
}

func copy(sshClient *ssh.Client, src, filename, dst string, size int64) (int64, error) {
	srcPath := fmt.Sprintf("%s/%s", src, filename)
	dstPath := fmt.Sprintf("%s/%s", dst, filename)

	//Check dst exist and the size is collect
	if fi, err := os.Stat(dstPath); err == nil {
		//exists, check size
		if fi.Size() == size {
			log.Printf("File %s already exists", dstPath)
			return fi.Size(), nil
		}
	}

	//TODO: Disk alert

	now := time.Now()
	scpClient, err := newSCPClient(sshClient)
	if err != nil {
		return 0, nil
	}

	//Creat dst folder if not exists
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err := os.Mkdir(dst, 0777)
		if err != nil {
			return 0, err
		}
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return 0, err
	}
	err = scpClient.CopyFromRemote(dstFile, srcPath)
	if err != nil {
		return 0, err
	}
	fInfo, err := dstFile.Stat()
	if err != nil {
		return 0, err
	}

	log.Printf("Copied %s to %s, size: %d GB, elapse: %s\n", srcPath, dstPath, fInfo.Size()/(1024*1024*1024), time.Since(now))
	return fInfo.Size(), dstFile.Close()
}

func remove(sshClient *ssh.Client, src, filename string) error {
	client, err := sftp.NewClient(sshClient)
	if err != nil {
		return err
	}

	srcPath := fmt.Sprintf("%s/%s", src, filename)
	return client.Remove(srcPath)
}
