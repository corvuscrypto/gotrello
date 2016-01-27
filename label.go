package trello

type label struct {
	url string
	Color,
	Name,
	Blue,
	Green,
	Orange,
	Yellow,
	Red,
	Purple staticField
	Board baseBoard
}

func createLabel(m model) label {
	lURL := m.getURL() + "/labels"
	return label{
		url:    lURL,
		Blue:   staticField(lURL + "/blue"),
		Green:  staticField(lURL + "/green"),
		Orange: staticField(lURL + "/orange"),
		Yellow: staticField(lURL + "/yellow"),
		Red:    staticField(lURL + "/red"),
		Purple: staticField(lURL + "/purple"),
	}
}

func (l label) ID(id string) label {
	labelURL := l.url + "/" + id
	l.url = labelURL
	l.Board = createBaseBoard(l)
	l.Color = staticField(labelURL + "/color")
	l.Name = staticField(labelURL + "/name")
	return l
}

func (l label) getURL() string {
	return l.url
}

var Labels = label{
	url: "/labels",
}
