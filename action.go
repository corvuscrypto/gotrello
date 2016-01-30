package trello

type baseAction struct {
	url string
	Comments,
	Data,
	Date,
	Display,
	Entities,
	IdMemberCreator,
	Type,
	Text staticField
}

func createBaseAction(m model) baseAction {
	aURL := m.getURL() + "/actions"
	return baseAction{
		url:             aURL,
		Comments:        staticField(aURL + "/comments"),
		Data:            staticField(aURL + "/data"),
		Date:            staticField(aURL + "/date"),
		IdMemberCreator: staticField(aURL + "/idMemberCreator"),
		Type:            staticField(aURL + "/type"),
		Display:         staticField(aURL + "/display"),
		Entities:        staticField(aURL + "/entities"),
		Text:            staticField(aURL + "/text"),
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

func (a baseAction) getURL() string {
	return a.url
}

type action struct {
	baseAction
	url           string
	Board         baseBoard
	Card          baseCard
	List          list
	Member        member
	MemberCreator memberCreator
	Organization  organization
}

func (a action) ID(id string) action {
	actionURL := a.url + "/" + id
	a.url = actionURL
	a.Comments = staticField(actionURL + "/comments")
	a.Data = staticField(actionURL + "/data")
	a.Date = staticField(actionURL + "/date")
	a.IdMemberCreator = staticField(actionURL + "/idMemberCreator")
	a.Display = staticField(actionURL + "/display")
	a.Entities = staticField(actionURL + "/entities")
	a.Text = staticField(actionURL + "/text")
	a.Type = staticField(actionURL + "/type")
	a.Board = createBaseBoard(a)
	a.Card = createBaseCard(a)
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
