package trello

type memberCreator struct {
	member
	url string
}

func createMemberCreator(m model) memberCreator {
	return memberCreator{
		url: m.getURL() + "/memberCreator",
	}
}

func (m memberCreator) getURL() string {
	return m.url
}
