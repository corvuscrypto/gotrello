package trello

import "testing"

var ID = "4f53aed56525abf3e12e2e45"

func TestGetURL(t *testing.T) {
	// test for Action
	testActions(t)
	// test for Board
	testBoards(t)
}
