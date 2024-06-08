default: cover
dev:
	chmod u+x scripts/*.sh
	@"$(CURDIR)/scripts/dev.sh"
mod:
	@go mod tidy
	@go mod vendor
test: mod
	@"$(CURDIR)/scripts/test.sh"
cover: test
	@go tool cover -html=coverage.out