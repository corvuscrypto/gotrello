package trello

type staticField string

type model interface {
	getURL() string
}

func (s staticField) getURL() string {
	return string(s)
}

type calKey struct {
	url      string
	Generate staticField
}

func createCalKey(m model) calKey {
	cURL := m.getURL() + "/calendarKey"
	return calKey{
		url:      cURL,
		Generate: staticField(cURL + "/generate"),
	}
}

type emailKey struct {
	url      string
	Generate staticField
}

func createEmailKey(m model) emailKey {
	eURL := m.getURL() + "/emailKey"
	return emailKey{
		url:      eURL,
		Generate: staticField(eURL + "/generate"),
	}
}

type labelNames struct {
	url string
	Blue,
	Green,
	Orange,
	Purple,
	Red,
	Yellow staticField
}

func createLabelNames(m model) labelNames {
	lURL := m.getURL() + "/labelNames"
	return labelNames{
		url:    lURL,
		Blue:   staticField(lURL + "/blue"),
		Green:  staticField(lURL + "/green"),
		Orange: staticField(lURL + "/orange"),
		Purple: staticField(lURL + "/purple"),
		Red:    staticField(lURL + "/red"),
		Yellow: staticField(lURL + "/yellow"),
	}
}

func (l labelNames) getURL() string {
	return l.url
}
