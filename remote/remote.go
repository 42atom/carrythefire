package remote

import "log"

func StartSCP(ip, src, dst, hostUsername, hostKeypath string) error {
	//Connect host
	sshClient, err := connectSSH(ip, "22", hostUsername, hostKeypath)
	if err != nil {
		return err
	}

	//Fetch plots
	plots, err := getPlots(sshClient, src)
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
