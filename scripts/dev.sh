#!/bin/bash
go install go.uber.org/mock/mockgen@latest
mkdir -p internal/core/{domain,ports,services/gamesrv}
go install  github.com/tommy-muehle/go-mnd/v2/cmd/mnd
asdf reshim golang