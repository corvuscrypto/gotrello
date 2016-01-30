package trello

import "testing"

var ID = "4f53aed56525abf3e12e2e45"

func TestGetURL(t *testing.T) {
	// test for Action
	testActions(t)
	// test for Board
	testBoards(t)
	// test for Cards
	testCards(t)
	// test for Checklists
	testChecklists(t)
	// test for Labels
	testLabels(t)
	// test for Lists
	testLists(t)
	// test for Members
	testMembers(t)
	// test for Notifications
	testNotifications(t)
	// test for Organizations
	testOrganizations(t)
	// test for Search
	testSearchs(t)
	// test for Session
	testSessions(t)
	// test for Tokens
	testTokens(t)
	// test for Types
	testTypes(t)
	// test for Webhooks
	testWebhooks(t)
}
