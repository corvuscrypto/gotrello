package trello

type board struct {
	url string
	//Lonely Fields
	BoardStars,
	Checklists,
	Closed,
	DateLastActivity,
	DateLastView,
	Deltas,
	Desc,
	DescData,
	IdOrganization,
	Invitations,
	Invited,
	Name,
	Pinned,
	ShortLink,
	ShortUrl,
	Starred,
	Subscribed,
	Url staticField
}

//Boards is the struct representing the boards model from trello
var Boards = board{
	url: "/boards",
}

func createBoard(m model) board {
	boardURL := m.getURL()
	return board{}
}

func (b board) getURL() string {
	return b.url
}
