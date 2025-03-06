package status

import (
	"errors"

	"github.com/ewatgar/cli-rock-paper-scissors/option"
)

type GameStatus int

const (
	Win GameStatus = iota
	Lose
	Stalemate
)

func (s GameStatus) String() string {
	return [...]string{"win", "lose", "stalemate"}[s]
}

func FromString(s string) (GameStatus, error) {
	statuses := map[string]GameStatus{
		"win":       Win,
		"lose":      Lose,
		"stalemate": Stalemate,
	}
	if status, ok := statuses[s]; ok {
		return status, nil
	}
	return 0, errors.New("invalid game status string")
}

func CheckWinStatus(player option.GameOption, rival option.GameOption) GameStatus {
	if player == rival {
		return Stalemate
	}

	var status GameStatus
	// paper > rock > scissors > paper
	switch player {
	case option.Rock:
		if rival == option.Scissors {
			status = Win
		}
		if rival == option.Paper {
			status = Lose
		}
	case option.Paper:
		if rival == option.Rock {
			status = Win
		}
		if rival == option.Scissors {
			status = Lose
		}
	case option.Scissors:
		if rival == option.Paper {
			status = Win
		}
		if rival == option.Rock {
			status = Lose
		}
	}

	return status
}
