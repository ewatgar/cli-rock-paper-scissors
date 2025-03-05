package option

import "errors"

type GameOption int

const (
	Rock GameOption = iota
	Paper
	Scissors
)

func (o GameOption) String() string {
	return [...]string{"rock", "paper", "scissors"}[o]
}

func FromString(s string) (GameOption, error) {
	options := map[string]GameOption{
		"rock":     Rock,
		"paper":    Paper,
		"scissors": Scissors,
	}
	if option, ok := options[s]; ok {
		return option, nil
	}
	return 0, errors.New("invalid game option string")
}
