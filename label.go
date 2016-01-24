package trello

type label struct {
	url string
}

func createLabel(m model) label {
	return label{
		url: m.getURL() + "/labels",
	}
}

func (l label) ID(id string) label {
	labelURL := l.url + "/" + id
	l.url = labelURL
	return l
}

func (l label) getURL() string {
	return l.url
}
