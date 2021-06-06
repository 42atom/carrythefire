package views

func fetchError() []string {
	return fetchFromHttp("error-list")
}
