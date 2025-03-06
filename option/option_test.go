package option

import "testing"

func TestOptionString(t *testing.T) {
	t.Run("TestStringRock", func(t *testing.T) {
		want := "rock"
		got := Rock.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStringPaper", func(t *testing.T) {
		want := "paper"
		got := Paper.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestStringScissors", func(t *testing.T) {
		want := "scissors"
		got := Scissors.String()

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}

func TestOptionFromString(t *testing.T) {
	t.Run("TestFromStringRock", func(t *testing.T) {
		want := Rock
		got, err := FromString("rock")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestFromStringPaper", func(t *testing.T) {
		want := Paper
		got, err := FromString("paper")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})

	t.Run("TestFromStringScissors", func(t *testing.T) {
		want := Scissors
		got, err := FromString("scissors")

		if err != nil {
			t.Error("error:", err)
		}

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}
