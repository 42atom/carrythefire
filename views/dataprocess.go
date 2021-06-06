package views

func fetchProcess() []string {
	return fetchFromHttp("process")
}
