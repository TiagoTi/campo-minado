default: test
mod:
	@go mod tidy
	@go mod vendor
dev:
	@go install go.uber.org/mock/mockgen@latest
generate: mod
	@"$(CURDIR)/scripts/generate.sh"
test: mod
	@"$(CURDIR)/scripts/test.sh"	