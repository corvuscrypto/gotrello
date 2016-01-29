package trello

type blankPlaceholder struct {
	url string
}

func createBlankPlaceholder(m model, name string) blankPlaceholder {
	return blankPlaceholder{
		url: m.getURL() + "/" + name,
	}
}

func (b blankPlaceholder) ID(id string) blankPlaceholder {
	bURL := b.getURL() + "/" + id
	return blankPlaceholder{
		url: bURL,
	}
}

func (b blankPlaceholder) getURL() string {
	return b.url
}
