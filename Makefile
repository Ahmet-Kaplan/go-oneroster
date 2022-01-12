build:
	go build \
		-o usulroster \
		main.go
deps: 
	cd $(CURDIR); \
	go get -v
build-linux:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64
	go build \
		-o usulroster-linux-amd64 \
		main.go
build-windows:
	CGO_ENABLED=0 \
	GOOS=windows\
	GOARCH=amd64
	go build \
		-o usulroster-windows-amd64.exe \
		main.go
