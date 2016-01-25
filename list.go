package trello

type list struct {
	url string
	Actions,
	Closed,
	IdBoard,
	Name,
	Pos,
	Subscribed,
	ArchiveAllCards,
	MoveAllCards staticField
}

func createList(m model) list {
	lURL := m.getURL()
	switch m.(type) {
	case action:
		lURL += "/list"
		break
	default:
		lURL += "/lists"
	}
	return list{
		url:             lURL,
		Actions:         staticField(lURL + "/actions"),
		Closed:          staticField(lURL + "/closed"),
		IdBoard:         staticField(lURL + "/idBoard"),
		Name:            staticField(lURL + "/name"),
		Pos:             staticField(lURL + "/pos"),
		Subscribed:      staticField(lURL + "/subscribed"),
		ArchiveAllCards: staticField(lURL + "/archiveAllCards"),
		MoveAllCards:    staticField(lURL + "/moveAllCards"),
	}
}

func (l list) ID(id string) list {
	listURL := l.url + "/" + id
	l.url = listURL
	l.Actions = staticField(listURL + "/actions")
	l.Closed = staticField(listURL + "/closed")
	l.IdBoard = staticField(listURL + "/idBoard")
	l.Name = staticField(listURL + "/name")
	l.Pos = staticField(listURL + "/pos")
	l.Subscribed = staticField(listURL + "/subscribed")
	l.ArchiveAllCards = staticField(listURL + "/archiveAllCards")
	l.MoveAllCards = staticField(listURL + "/moveAllCards")
	return l
}

func (l list) getURL() string {
	return l.url
}
