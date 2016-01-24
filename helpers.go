package trello

type staticField string

type model interface {
	getURL() string
}

func (s staticField) getURL() string {
	return string(s)
}
