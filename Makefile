.PHONY: clear build
clear:
	rm -rf build
build: clear
	GOOS=linux GOARCH=amd64 go build -o ./build/plot-carrier ./main.go