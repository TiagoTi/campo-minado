package gamesrv_test

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires/mocks"
	"campo-minado/internal/core/services/gamesrv"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetGame(t *testing.T) {
	t.Parallel()
	var (
		gameService provides.GamesService
	)

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	t.Run("get game with sucess", func(t *testing.T) {
		mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
		mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)
		mockGamRep.EXPECT().Get("123").Return(&domain.Game{
			ID:      "123",
			Lines:   4,
			Columns: 3,
			Mines:   2,
		}, nil)

		gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

		assert.NotNil(t, gameService)

		game, newGameError := gameService.Get("123")
		assert.NotNil(t, game)
		require.NoError(t, newGameError)

		assert.Equal(t, "123", game.ID)
		assert.Equal(t, int(4), game.Lines)
		assert.Equal(t, int(3), game.Columns)
		assert.Equal(t, int(2), game.Mines)
	})
	t.Run("fail to get game", func(t *testing.T) {
		mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
		mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)

		gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

		assert.NotNil(t, gameService)

		game, newGameError := gameService.Get("")
		assert.Nil(t, game)
		require.Error(t, newGameError)
		assert.Equal(t, "id do jogo deve ser informado", newGameError.Error())

	})

	t.Run("to get game with error", func(t *testing.T) {
		mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
		mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)
		mockGamRep.EXPECT().Get("123").Return(nil, fmt.Errorf(`erro de conexão`))

		gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

		assert.NotNil(t, gameService)

		game, newGameError := gameService.Get("123")
		assert.Nil(t, game)
		require.Error(t, newGameError)
		assert.Equal(t, "erro de conexão", newGameError.Error())
	})

	t.Run("game not exists", func(t *testing.T) {
		mockGamRep := mocks.NewMockGamesRepository(mockCtrl)
		mockBoardGenerator := mocks.NewMockBoardGenerator(mockCtrl)
		mockGamRep.EXPECT().Get("123").Return(nil, nil)

		gameService = gamesrv.NewGameService(mockGamRep, mockBoardGenerator)

		assert.NotNil(t, gameService)

		game, newGameError := gameService.Get("123")
		assert.Nil(t, game)
		require.Error(t, newGameError)
		assert.Equal(t, "jogo não existe", newGameError.Error())
	})

}
