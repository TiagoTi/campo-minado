default: test
mod:
	@go mod tidy
dev:
	@"$(CURDIR)/scripts/dev.sh"
lint:
	@"$(CURDIR)/scripts/golangci-lint.sh"
generate: mod
	@"$(CURDIR)/scripts/generate.sh"
test: mod
	@"$(CURDIR)/scripts/test.sh"
cover: test
	@go tool cover -html=coverage.out
template:
	@"$(CURDIR)/scripts/template.sh"