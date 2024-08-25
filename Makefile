build: build_inter build_apple
build_inter:
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o tshell main.go
	chmod +x tshell
build_apple:
	GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 go build -ldflags="-s -w" -o tshell main.go
	chmod +x tshell
