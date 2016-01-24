package trello

type action struct {
	url string
	Display,
	Entities,
	Text staticField
	Board         board
	Card          card
	List          list
	Member        member
	MemberCreator memberCreator
	Organization  organization
}

func createAction(m model) action {
	actionURL := m.getURL() + "/action"

	return action{
		url: actionURL,
	}
}

func (a action) ID(id string) action {
	actionURL := a.url + "/" + id
	a.url = actionURL
	a.Display = staticField(actionURL + "/display")
	a.Entities = staticField(actionURL + "/entities")
	a.Text = staticField(actionURL + "/text")
	a.Board = createBoard(a)
	a.Card = createCard(a)
	a.List = createList(a)
	a.Member = createMember(a)
	a.MemberCreator = createMemberCreator(a)
	a.Organization = createOrganization(a)
	return a
}

func (a action) getURL() string {
	return a.url
}

var Actions = action{
	url: "/actions",
}
