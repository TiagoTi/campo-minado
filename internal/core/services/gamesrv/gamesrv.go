package gamesrv

import (
	"campo-minado/internal/core/domain"
	"campo-minado/internal/core/provides"
	"campo-minado/internal/core/requires"
	"fmt"
)

func NewGameService(repository requires.GamesRepository) provides.GamesService {
	return &service{"", repository}
}

type service struct {
	message    string
	repository requires.GamesRepository
}

func (s *service) Error() string {
	return s.message
}

func (s *service) Create(lines, columns, mines uint) (*domain.Game, error) {
	if lines < 1 {
		s.message = fmt.Sprintf("o tabuleiro do jogo deve ter um quantidade de linhas maior que um e recebeu %d", lines)

		return nil, s
	}

	if columns < 1 {
		s.message = fmt.Sprintf("deve ter pelo menos 2 colunas. Recebeu %d", columns)

		return nil, s
	}

	if mines < 1 {
		s.message = fmt.Sprintf("uma sessão de jogo deve possuir no mínimo uma mina e recebeu %d", mines)

		return nil, s
	}

	if mines >= lines*columns {
		s.message = fmt.Sprintf("uma sessão deve  ter menos do que %d minas e recebeu %d", lines*columns, mines)

		return nil, s
	}
	savedGame, err := s.repository.Save(domain.Game{
		Lines:   lines,
		Columns: columns,
		Mines:   mines,
	})
	if err != nil {
		return nil, err
	}
	return savedGame, nil
}
