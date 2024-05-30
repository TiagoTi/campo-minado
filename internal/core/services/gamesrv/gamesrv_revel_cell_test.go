package gamesrv_test

import (
	"testing"

	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRevelCell(t *testing.T) {
	t.Skip()
	t.Parallel()
	var (
		lines       uint
		columns     uint
		mines       uint
		gameService provides.GamesService
	)

	lines = 2

	columns = 2

	mines = 1

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)

	gameService = gamesrv.NewGameService(mockGamRep)

	assert.NotNil(t, gameService)

	game, newGameError := gameService.Create(lines, columns, mines)
	assert.NotNil(t, game)
	require.NoError(t, newGameError)
	t.Run("O jogador só pode realizar revelações em posições válidas do tabuleiro.", RN005)
}
