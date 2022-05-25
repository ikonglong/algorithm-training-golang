package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicTacToe(t *testing.T) {
	board := []string{"O X", " XO", "X O"}
	assert.Equal(t, "X", tictactoe(board))
}

func TestTicTacToe_bug(t *testing.T) {
	board := []string{"O X", " XO", "X O"}
	// bug: O->1, X->2。第一行，[O, ' ', 'X'] 和也为 3
	assert.Equal(t, "X", tictactoe_bug(board))
}
