package trello

type list struct {
	url string
}

func (l list) getURL() string {
	return l.url
}
