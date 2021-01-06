GOPATH=$(shell go env GOPATH)
PATH=$(GOPATH)/bin:$(shell echo $$PATH)

gqlgen:
	@go run github.com/99designs/gqlgen generate
	