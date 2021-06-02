package remote

import "fmt"

func StartSCP(ip, src, dst, hostUsername, hostKeypath string) error {
	sshClient, err := connectSSH(ip, "22", hostUsername, hostKeypath)
	if err != nil {
		return err
	}

	plots := getPlots(sshClient, src)
	fmt.Println(plots)

	return nil
}
