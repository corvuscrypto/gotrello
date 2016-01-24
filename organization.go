package trello

type organization struct {
	url string
}

func (o organization) getURL() string {
	return o.url
}
