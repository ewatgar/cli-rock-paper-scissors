package main

import (
	"testing"

	"github.com/ewatgar/cli-rock-paper-scissors/option"
	"github.com/ewatgar/cli-rock-paper-scissors/status"
)

// paper > rock > scissors > paper
func TestPlayerWin(t *testing.T) {
	t.Run("TestWinPlayerRockRivalScissors", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Scissors
		want := status.Win
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestWinPlayerScissorsRivalPaper", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Paper
		want := status.Win
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestWinPlayerPaperRivalRock", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Rock
		want := status.Win
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}

func TestPlayerLose(t *testing.T) {
	t.Run("TestLosePlayerScissorsRivalRock", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Rock
		want := status.Lose
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestLosePlayerPaperRivalScissors", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Scissors
		want := status.Lose
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestLosePlayerRockRivalPaper", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Paper
		want := status.Lose
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

}

func TestPlayerStalemate(t *testing.T) {
	t.Run("TestStalematePlayerRockRivalRock", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Rock
		want := status.Stalemate
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStalematePlayerPaperRivalPaper", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Paper
		want := status.Stalemate
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStalematePlayerScissorsRivalScissors", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Scissors
		want := status.Stalemate
		got := status.CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}
