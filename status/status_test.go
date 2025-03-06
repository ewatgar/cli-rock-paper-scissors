package status

import (
	"testing"

	"github.com/ewatgar/cli-rock-paper-scissors/option"
)

func TestStatusString(t *testing.T) {
	t.Run("TestStringWin", func(t *testing.T) {
		want := "win"
		got := Win.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStringLose", func(t *testing.T) {
		want := "lose"
		got := Lose.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStringStalemate", func(t *testing.T) {
		want := "stalemate"
		got := Stalemate.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}

func TestStatusFromString(t *testing.T) {
	t.Run("TestFromStringWin", func(t *testing.T) {
		want := Win
		got, err := FromString("win")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestFromStringLose", func(t *testing.T) {
		want := Lose
		got, err := FromString("lose")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestFromStringStalemate", func(t *testing.T) {
		want := Stalemate
		got, err := FromString("stalemate")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestFromStringError", func(t *testing.T) {
		got, err := FromString("invalid")

		if !(got == 0 && err != nil) {
			t.Error("invalid option should throw exception")
		}
	})
}

// paper > rock > scissors > paper
func TestPlayerWin(t *testing.T) {
	t.Run("TestWinPlayerRockRivalScissors", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Scissors
		want := Win
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestWinPlayerScissorsRivalPaper", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Paper
		want := Win
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestWinPlayerPaperRivalRock", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Rock
		want := Win
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}

func TestPlayerLose(t *testing.T) {
	t.Run("TestLosePlayerScissorsRivalRock", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Rock
		want := Lose
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestLosePlayerPaperRivalScissors", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Scissors
		want := Lose
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestLosePlayerRockRivalPaper", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Paper
		want := Lose
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

}

func TestPlayerStalemate(t *testing.T) {
	t.Run("TestStalematePlayerRockRivalRock", func(t *testing.T) {
		playerOption := option.Rock
		rivalOption := option.Rock
		want := Stalemate
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStalematePlayerPaperRivalPaper", func(t *testing.T) {
		playerOption := option.Paper
		rivalOption := option.Paper
		want := Stalemate
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStalematePlayerScissorsRivalScissors", func(t *testing.T) {
		playerOption := option.Scissors
		rivalOption := option.Scissors
		want := Stalemate
		got := CheckWinStatus(playerOption, rivalOption)

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}
