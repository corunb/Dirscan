all: win linux mac m1

win:
	# windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o Dirscan.exe main.go
linux:
	# linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o  Dirscan_linux main.go
mac:
	# MacOS
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o Dirscan_darwin_mac main.go
m1:
	# MacOSm1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w " -trimpath -o Dirscan_darwin_arm main.go

