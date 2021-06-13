build:
	go build -o bin/yymr ./cmd/main.go
	go build -o bin/yymr-pt ./cmd/wr/main.go
	go build -o bin/yymr-exec ./cmd/exec/main.go

