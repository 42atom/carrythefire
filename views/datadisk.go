package views

func fetchDisk() []string {
	return fetchFromHttp("disk-usage")
}
