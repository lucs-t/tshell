VERSION:=1.0.0

build_linux:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o tshell main.go
	chmod +x tshell
build_mac_arm64:
	GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 go build -ldflags="-s -w" -o tshell main.go
	chmod +x tshell
build_mac_amd64:
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o tshell main.go
	chmod +x tshell

pack: pack_linux pack_mac_arm64 pack_mac_amd64
pack_linux: build_linux
	tar -czvf tshell-$(VERSION)-linux.tar.gz tshell
pack_mac_arm64: build_mac_arm64
	tar -czvf tshell-$(VERSION)-macos-arm64.tar.gz tshell
pack_mac_amd64: build_mac_amd64
	tar -czvf tshell-$(VERSION)-macos-amd64.tar.gz tshell

