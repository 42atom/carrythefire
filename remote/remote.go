package remote

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

	//Start moving files
	for filename, size := range plots {
		err := mv(sshClient, src, filename, dst, size)
		if err != nil {
			return err
		}
	}

	return nil
}
