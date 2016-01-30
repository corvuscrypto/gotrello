package trello

type search struct {
	url     string
	Members staticField
}

func createSearch(m model) search {
	sURL := m.getURL() + "/search"
	return search{
		url:     sURL,
		Members: staticField(sURL + "/members"),
	}
}

func (s search) getURL() string {
	return s.url
}

var Search = search{
	url:     "/search",
	Members: staticField("/search/members"),
}
