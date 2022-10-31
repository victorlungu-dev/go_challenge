test:
	go test ./...

cover:
	sh ./cover.sh

build:
	docker build --target release -t go_chall:latest .

