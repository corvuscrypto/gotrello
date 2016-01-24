package trello

type checklist struct {
	url string
}

func createChecklist(m model) checklist {
	return checklist{
		url: m.getURL() + "/checklist",
	}
}

func (c checklist) ID(id string) checklist {
	cURL := c.url + "/" + id
	c.url = cURL
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
