.PHONY: go-mod test run clean

go-mod:
	go mod download
test:
	 go test ./...
run:
	go run apps/blog/main.go

docker:
	sudo docker compose up
clean:
	go clean
	rm -f myapp

