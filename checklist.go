package trello

type checkItem struct {
	url string
	Name,
	Pos,
	State,
	ConvertToCard staticField
}

func createCheckItems(m model) checkItem {
	return checkItem{
		url: m.getURL() + "/checkItem",
	}
}

func (c checkItem) ID(id string) checkItem {
	c.url = c.url + "/" + id
	c.Name = staticField(c.url + "/name")
	c.Pos = staticField(c.url + "/pos")
	c.State = staticField(c.url + "/state")
	c.ConvertToCard = staticField(c.url + "/convertToCard")
	return c
}

func (c checkItem) getURL() string {
	return c.url
}

type checklist struct {
	url string
	IdBoard,
	IdCard,
	Name,
	Pos staticField
	//submodels
	Board      baseBoard
	Cards      filterCard
	CheckItems checkItem
	CheckItem  checkItem
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
	c.Board = createBaseBoard(c)
	c.Cards = createFilterCard(c)
	c.CheckItems = createCheckItems(c)
	c.CheckItem = c.CheckItems
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

var Checklists = checklist{
	url: "/checklists",
}
