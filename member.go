package trello

type member struct {
	url string
}

func (m member) getURL() string {
	return m.url
}
