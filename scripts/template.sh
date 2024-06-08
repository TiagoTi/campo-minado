#!/bin/bash
VERSION=005
cat <<- EOF > ./internal/core/services/gamesrv/gamesrv_rn${VERSION}_test.go
package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/services/gamesrv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RN$VERSION(t *testing.T) {
	var (
		newGameError error
		game         *domain.Game
		lines        int = 3
		columns      int = 3
        mines        int = 1
	)

	gameService := gamesrv.NewGameService()

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(lines, columns, mines)

	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "-", newGameError.Error())
}
EOF
