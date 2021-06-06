package remote

import (
	"log"

	"golang.org/x/crypto/ssh"
)

type Task struct {
	Client       *ssh.Client
	IP           string
	HostUserName string
	BindAddress  string
	Src          string
	FileName     string
	Dst          string
	Size         int64
}

func StartSCP(ip, bindAddress, src, dst, hostUsername, hostKeypath string) error {
	//Connect host
	sshClient, err := ConnectSSH(ip, "22", hostUsername, hostKeypath)
	if err != nil {
		return err
	}

	//Fetch plots
	plots, err := GetPlots(sshClient, src)
	if err != nil {
		return err
	}

	if len(plots) == 0 {
		log.Printf("There are no plots on %s, %s", ip, src)
	}

	//Start moving files
	for filename, size := range plots {
		err := mv(sshClient, src, filename, dst, size)
		if err != nil {
			return err
		}
	}

	return nil
}

func StartSCPSimple(ip, bindAddress, src, dst, hostUsername, hostKeypath string, workerNum int) error {
	//Connect host
	sshClient, err := ConnectSSH(ip, "22", hostUsername, hostKeypath)
	if err != nil {
		return err
	}

	//Fetch plots
	plots, err := GetPlots(sshClient, src)
	if err != nil {
		return err
	}

	if len(plots) == 0 {
		log.Printf("There are no plots on %s, %s", ip, src)
		return nil
	}
	//Change to slice
	pp := []*Task{}
	for filename, size := range plots {
		t := &Task{
			Client:       sshClient,
			IP:           ip,
			HostUserName: hostUsername,
			BindAddress:  bindAddress,
			Src:          src,
			FileName:     filename,
			Dst:          dst,
			Size:         size,
		}
		pp = append(pp, t)
	}

	t := pp[0]
	return simpleMV(t.Client, t.IP, t.HostUserName, t.BindAddress, t.Src, t.FileName, t.Dst, t.Size)
}
