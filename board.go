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
	Url string
}

type boardStars struct {
	url string
}

//Boards is the struct representing the boards model from trello
var Boards = board{
	url: "/boards",
}

func (b board) ID(id string) board {
	return board{url: b.url + "/" + id}
}

func (b board) Field(fieldName string) board {
	return board{url: b.url + "/" + fieldName}
}

func (b board) getURL() string {
	return b.url
}
