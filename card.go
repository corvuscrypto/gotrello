package trello

type card struct {
	url string
	Badges,
	CheckItemStates,
	Closed,
	DateLastActivity,
	Desc,
	DescData,
	Due,
	Email,
	IdAttachmentCover,
	IdBoard,
	IdChecklists,
	IdLabels,
	IdList,
	IdMembers,
	IdMembersVoted,
	IdShort,
	Labels,
	ManualCoverAttachment,
	Name,
	Pos,
	ShortLink,
	ShortUrl,
	Subscribed,
	Url,
	Checklists staticField
}

func createCard(m model) card {
	return card{
		url: m.getURL() + "/cards",
	}
}

func (c card) ID(id string) card {
	cardURL := c.url + "/" + id
	c.url = cardURL
	c.Badges = staticField(cardURL + "/badges")
	c.CheckItemStates = staticField(cardURL + "/checkItemStates")
	c.Checklists = staticField(cardURL + "/checklists")
	c.Closed = staticField(cardURL + "/closed")
	c.DateLastActivity = staticField(cardURL + "/dateLastActivity")
	c.Desc = staticField(cardURL + "/desc")
	c.DescData = staticField(cardURL + "/descData")
	c.Due = staticField(cardURL + "/due")
	c.Email = staticField(cardURL + "/email")
	c.IdAttachmentCover = staticField(cardURL + "/idAttachmentCover")
	c.IdBoard = staticField(cardURL + "/idBoard")
	c.IdChecklists = staticField(cardURL + "/idChecklist")
	c.IdLabels = staticField(cardURL + "/idLabels")
	c.IdList = staticField(cardURL + "/idList")
	c.IdMembers = staticField(cardURL + "/idMembers")
	c.IdMembersVoted = staticField(cardURL + "/idMembersVoted")
	c.IdShort = staticField(cardURL + "/idShort")
	c.Labels = staticField(cardURL + "/labels")
	c.ManualCoverAttachment = staticField(cardURL + "/manualCoverAttachment")
	c.Name = staticField(cardURL + "/name")
	c.Pos = staticField(cardURL + "/pos")
	c.ShortLink = staticField(cardURL + "/shortLink")
	c.ShortUrl = staticField(cardURL + "/shortUrl")
	c.Subscribed = staticField(cardURL + "/subscribed")
	c.Url = staticField(cardURL + "/url")
	return c
}

func (c card) getURL() string {
	return c.url
}
