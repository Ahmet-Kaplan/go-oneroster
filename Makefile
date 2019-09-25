build:
	CGO_ENABLED=1 \
	GOOS=linux \
	go build \
	-a \
	-ldflags '-linkmode external -extldflags "-static"' \
	-o goors \
	cmd/goors/main.go
