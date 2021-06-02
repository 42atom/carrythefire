package remote

import (
	"fmt"
	"time"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func connectSSH(ip, port, username, keypath string) (*ssh.Client, error) {

	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	clientConfig, err := auth.PrivateKey(username, keypath, ssh.InsecureIgnoreHostKey())
	if err != nil {
		return nil, err
	}
	clientConfig.Timeout = time.Second * 15
	host := fmt.Sprintf("%s:%s", ip, port)
	return ssh.Dial("tcp", host, &clientConfig)
}

func newSCPClient(sshClient *ssh.Client) (*scp.Client, error) {
	client, err := scp.NewClientBySSH(sshClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
