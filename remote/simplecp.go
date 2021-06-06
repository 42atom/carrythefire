package remote

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"golang.org/x/crypto/ssh"
)

func simpleMV(sshClient *ssh.Client, ip, hostUsername, bindAddress, src, filename, dst string, size int64) error {
	log.Println("Start simple scp...")
	csize, err := simpleCopy(ip, hostUsername, bindAddress, src, filename, dst, size)
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

func simpleCopy(ip, hostUsername, bindAddress, src, filename, dst string, size int64) (int64, error) {
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

	now := time.Now()

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

	args := fmt.Sprintf("%s@%s:%s", hostUsername, ip, srcPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("scp", args, dst)
	if bindAddress != "" {
		cmd = exec.Command("scp", "-o", fmt.Sprintf("BindAddress=%s", bindAddress), args, dst)
	}
	//fmt.Println(cmd.String())
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return 0, fmt.Errorf(fmt.Sprintf("%s:%s", err.Error(), stderr.String()))
	}

	fInfo, err := dstFile.Stat()
	if err != nil {
		return 0, err
	}

	log.Printf("Copied %s to %s, size: %d GB, elapse: %s\n", srcPath, dstPath, fInfo.Size()/(1024*1024*1024), time.Since(now))
	return fInfo.Size(), dstFile.Close()
}
