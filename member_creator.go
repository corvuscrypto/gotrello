package trello

type memberCreator struct {
	url string
}

func (m memberCreator) getURL() string {
	return m.url
}
