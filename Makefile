dep: ## update dependencies
	@dep ensure -vendor-only

build: dep ## build the binary
	@go build -i cmd/ensure.go