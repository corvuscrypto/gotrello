package trello

type boardStars struct {
	url string
}

func (b boardStars) getURL() string {
	return b.url
}
