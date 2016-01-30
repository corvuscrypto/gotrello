package trello

type search struct {
	url     string
	Members staticField
}

func (s search) getURL() string {
	return s.url
}

var Search = search{
	url:     "/search",
	Members: staticField("/search/members"),
}
