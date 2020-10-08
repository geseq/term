format:
	@echo "Formatting code to comply with goimports"
	goimports -w ./

vet:
	go vet ./...

lint:
	golint -set_exit_status ./...

test: format vet lint
	go test -race -cover ./...
