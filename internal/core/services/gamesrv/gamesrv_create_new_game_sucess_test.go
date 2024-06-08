package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func CreateNewGameSucess(t *testing.T) {
	var (
		lines       int
		columns     int
		mines       int
		gameService provides.GamesService
	)

	lines = 2

	columns = 2

	mines = 1

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
	mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)

	mockGamRep.EXPECT().Save(domain.Game{Lines: lines, Columns: columns, Mines: mines}).Return(&domain.Game{
		ID:      "123",
		Lines:   lines,
		Columns: columns,
		Mines:   mines,
	}, nil)

	gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

	assert.NotNil(t, gameService)

	game, newGameError := gameService.Create(lines, columns, mines)
	assert.NotNil(t, game)
	require.NoError(t, newGameError)
	assert.Equal(t, "123", game.ID)
	assert.Equal(t, lines, game.Lines)
	assert.Equal(t, columns, game.Columns)
	assert.Equal(t, mines, game.Mines)
}
