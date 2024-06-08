package gamesrv_test

import (
	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func CreateNewGameErrorDatabase(t *testing.T) {
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

	mockGamRep.EXPECT().Save(gomock.Any()).Return(nil, fmt.Errorf(`erro ao gravar`))

	gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

	assert.NotNil(t, gameService)

	game, newGameError := gameService.Create(lines, columns, mines)
	assert.Nil(t, game)
	require.Error(t, newGameError)
}
