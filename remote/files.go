package remote

import (
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getPlots(sshClient *ssh.Client, src string) []string {
	client, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}
	w := client.Walk("/home/project/")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}
	return nil
}

func mv(ip, src, filename string) error {
	err := copy(ip, src, filename)
	if err != nil {
		return err
	}
	err = remove(ip, src, filename)
	return err
}

func copy(ip, src, filename string) error {

	return nil
}

func remove(ip, src, filename string) error {
	return nil
}
