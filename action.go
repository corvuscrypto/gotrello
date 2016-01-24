package trello

type baseAction struct {
	url string
	Comments,
	Display,
	Entities,
	Text staticField
}

func createBaseAction(m model) baseAction {
	aURL := m.getURL() + "/actions"
	return baseAction{
		url:      aURL,
		Comments: staticField(aURL + "/comments"),
		Display:  staticField(aURL + "/display"),
		Entities: staticField(aURL + "/entities"),
		Text:     staticField(aURL + "/text"),
	}
}

func (a baseAction) ID(id string) baseAction {
	aURL := a.url + "/" + id
	a.Comments = staticField(aURL + "/comments")
	a.Display = staticField(aURL + "/display")
	a.Entities = staticField(aURL + "/entities")
	a.Text = staticField(aURL + "/text")
	return a
}

type action struct {
	baseAction
	url           string
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

//Actions is the variable through which all chaining for other model fields can be accessed
var Actions = action{
	url: "/actions",
}
