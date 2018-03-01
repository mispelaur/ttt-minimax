package main

import (
	"reflect"
	"testing"
)

func TestIsWinner(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		board          []int
		player         int
		playerIsWinner bool
	}{
		{
			name:           "Computer doesn't win if there's a tie",
			board:          []int{computer, user, computer, user, computer, user, user, computer, user},
			player:         computer,
			playerIsWinner: false,
		},
		{
			name:           "User doesn't win if there's a tie",
			board:          []int{computer, user, computer, user, computer, user, user, computer, user},
			player:         user,
			playerIsWinner: false,
		},
		{
			name:           "Computer wins when it should",
			board:          []int{computer, 0, user, computer, user, user, computer, 0, 0},
			player:         computer,
			playerIsWinner: true,
		},
		{
			name:           "Computer loses when it should",
			board:          []int{user, 0, computer, user, computer, computer, user, 0, 0},
			player:         computer,
			playerIsWinner: false,
		},
		{
			name:           "User wins when they should",
			board:          []int{user, 0, computer, user, computer, computer, user, 0, 0},
			player:         user,
			playerIsWinner: true,
		},
		{
			name:           "User loses when they should",
			board:          []int{computer, 0, user, computer, user, user, computer, 0, 0},
			player:         user,
			playerIsWinner: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			val := isWinner(c.board, c.player)

			if val != c.playerIsWinner {
				t.Fatalf("expected board %v to yield value %v for player %v, but instead got %v", c.board, c.playerIsWinner, c.player, val)
			}
		})
	}
}

func TestFindEmpties(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		board   []int
		empties []int
	}{
		{
			name:    "No empty spaces",
			board:   []int{computer, user, computer, user, computer, user, user, computer, user},
			empties: []int{},
		},
		{
			name:    "Some empty spaces",
			board:   []int{user, 0, computer, user, computer, computer, user, 0, 0},
			empties: []int{1, 7, 8},
		},
		{
			name:    "All spaces empty",
			board:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			empties: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			val := findEmpties(c.board)

			if len(c.empties) == 0 && len(val) == 0 {
				return
			}
			if !reflect.DeepEqual(val, c.empties) {
				t.Fatalf("Got %v but expected %v", val, c.empties)
			}
		})
	}
}

func TestChooseComputerMove(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		board           []int
		acceptableMoves []int
	}{
		{
			name:            "Block user from winning diagonal",
			board:           []int{0, 0, 0, user, user, 0, computer, computer, user},
			acceptableMoves: []int{0},
		},
		{
			name:            "Block user from winning horizontal",
			board:           []int{user, user, 0, 0, computer, 0, 0, computer, user},
			acceptableMoves: []int{2},
		},
		{
			name:            "Block user from winning vertical",
			board:           []int{computer, user, 0, 0, user, 0, 0, 0, 0},
			acceptableMoves: []int{7},
		},
		{
			name:            "Computer win diagonal",
			board:           []int{computer, user, user, 0, 0, 0, user, 0, computer},
			acceptableMoves: []int{4},
		},
		{
			name:            "Computer win horizontal",
			board:           []int{user, 0, 0, 0, user, user, computer, 0, computer},
			acceptableMoves: []int{7},
		},
		{
			name:            "Computer win vertical",
			board:           []int{user, 0, 0, 0, computer, 0, 0, computer, user},
			acceptableMoves: []int{1},
		},
		{
			name:            "First user move in top right corner",
			board:           []int{0, 0, 0, 0, 0, 0, 0, 0, user},
			acceptableMoves: []int{4, 5, 7},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			val := chooseComputerMove(c.board)

			pass := false
			for _, option := range c.acceptableMoves {
				if val == option {
					pass = true
					break
				}
			}
			if !pass {
				t.Fatalf("Expected next computer move for board %v to be one of %v, but instead got %v", c.board, c.acceptableMoves, val)
			}
		})
	}
}
