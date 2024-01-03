
default:
	go build -o build/hashpass

fmt:
	go fmt main.go

linux-x64:
	GOOS=linux GOARCH=amd64 go build -o build/hashpass-linux-amd64 main.go

macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o build/hashpass-macos-arm64 main.go

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o build/hashpass-linux-arm64  main.go

