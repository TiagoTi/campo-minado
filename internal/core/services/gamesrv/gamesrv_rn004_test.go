package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func RN004(t *testing.T) {
	var (
		newGameError error
		game         *domain.Game
		lines        uint = 2
		columns      uint = 2
	)

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockGamRep := mocks.NewMockGamesRepository(mockCtrl)

	gameService := gamesrv.NewGameService(mockGamRep)

	assert.NotNil(t, gameService)

	game, newGameError = gameService.Create(lines, columns, 4)

	assert.Nil(t, game)
	require.Error(t, newGameError)
	assert.Equal(t, "uma sessão deve  ter menos do que 4 minas e recebeu 4", newGameError.Error())
}
