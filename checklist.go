package trello

type checkItems struct {
	url string
}

func createCheckItems(m model) checkItems {
	return checkItems{
		url: m.getURL() + "/checkItems",
	}
}

func (c checkItems) ID(id string) checkItems {
	c.url = c.url + "/" + id
	return c
}

type checklist struct {
	url string
	IdBoard,
	IdCard,
	Name,
	Pos staticField
	//submodels
	Board      baseBoard
	Cards      baseCard
	CheckItems checkItems
}

func createChecklist(m model) checklist {
	return checklist{
		url: m.getURL() + "/checklist",
	}
}

func (c checklist) ID(id string) checklist {
	cURL := c.url + "/" + id
	c.url = cURL
	c.IdBoard = staticField(cURL + "/idBoard")
	c.IdCard = staticField(cURL + "/idCard")
	c.Name = staticField(cURL + "/name")
	c.Pos = staticField(cURL + "/pos")
	return c
}

func (c checklist) getURL() string {
	return c.url
}

type checklists struct {
	url string
}

func createChecklists(m model) checklists {
	return checklists{
		url: m.getURL() + "/checklists",
	}
}
func (c checklists) ID(id string) checklists {
	c.url = c.url + "/" + id
	return c
}

func (c checklists) getURL() string {
	return c.url
}

var Checklist = checklist{
	url: "/checklists",
}
