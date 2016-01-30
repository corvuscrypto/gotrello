package trello

type attachment struct {
	url string
}

func createAttachment(m model) attachment {
	return attachment{
		url: m.getURL() + "/attachments",
	}
}

func (a attachment) ID(id string) attachment {
	a.url = a.url + "/" + id
	return a
}

func (a attachment) getURL() string {
	return a.url
}

type idLabel struct {
	url string
}

func createIdLabel(m model) idLabel {
	return idLabel{
		url: m.getURL() + "/idLabels",
	}
}

func (i idLabel) ID(id string) idLabel {
	i.url = i.url + "/" + id
	return i
}

func (i idLabel) getURL() string {
	return i.url
}

type idMember struct {
	url string
}

func createIdMember(m model) idMember {
	return idMember{
		url: m.getURL() + "/idMembers",
	}
}

func (i idMember) ID(id string) idMember {
	i.url = i.url + "/" + id
	return i
}

func (i idMember) getURL() string {
	return i.url
}

type membersVoted struct {
	url string
}

func createMembersVoted(m model) membersVoted {
	return membersVoted{
		url: m.getURL() + "/membersVoted",
	}
}

func (i membersVoted) ID(id string) membersVoted {
	i.url = i.url + "/" + id
	return i
}

func (m membersVoted) getURL() string {
	return m.url
}

type filterCard struct {
	url string
	All,
	Closed,
	Open,
	None,
	Visible staticField
}

func createFilterCard(m model) filterCard {
	cURL := m.getURL() + "/cards"
	return filterCard{
		url:     cURL,
		All:     staticField(cURL + "/all"),
		Closed:  staticField(cURL + "/closed"),
		Open:    staticField(cURL + "/open"),
		None:    staticField(cURL + "/none"),
		Visible: staticField(cURL + "/visible"),
	}
}

func (f filterCard) ID(id string) filterCard {
	fURL := f.url + "/" + id
	f.url = fURL
	f.All = staticField(fURL + "/all")
	f.Closed = staticField(fURL + "/closed")
	f.None = staticField(fURL + "/none")
	f.Open = staticField(fURL + "/open")
	f.Visible = staticField(fURL + "/visible")
	return f
}

func (c filterCard) getURL() string {
	return c.url
}

type baseCard struct {
	filterCard
	url string
	//static fields
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
	IdChecklist,
	IdChecklists,
	IdLabels,
	IdList,
	IdMembers,
	IdMembersVoted,
	IdShort,
	Labels,
	ManualCoverAttachment,
	MarkAssociatedNotificationsRead,
	Members,
	Name,
	Pos,
	ShortLink,
	ShortUrl,
	Subscribed,
	Url staticField
}

func createBaseCard(m model) baseCard {
	var cardURL string

	//Handle special case for type
	switch m.(type) {
	case action, notifications:
		cardURL = m.getURL() + "/card"
		break
	default:
		cardURL = m.getURL() + "/cards"
	}

	c := baseCard{}
	c.filterCard = createFilterCard(m)
	c.url = cardURL
	c.Badges = staticField(cardURL + "/badges")
	c.CheckItemStates = staticField(cardURL + "/checkItemStates")
	c.Closed = staticField(cardURL + "/closed")
	c.DateLastActivity = staticField(cardURL + "/dateLastActivity")
	c.Desc = staticField(cardURL + "/desc")
	c.DescData = staticField(cardURL + "/descData")
	c.Due = staticField(cardURL + "/due")
	c.Email = staticField(cardURL + "/email")
	c.IdAttachmentCover = staticField(cardURL + "/idAttachmentCover")
	c.IdBoard = staticField(cardURL + "/idBoard")
	c.IdChecklists = staticField(cardURL + "/idChecklists")
	c.IdChecklist = staticField(cardURL + "/idChecklist")
	c.IdLabels = staticField(cardURL + "/idLabels")
	c.IdList = staticField(cardURL + "/idList")
	c.IdMembers = staticField(cardURL + "/idMembers")
	c.IdMembersVoted = staticField(cardURL + "/idMembersVoted")
	c.Labels = staticField(cardURL + "/labels")
	c.IdShort = staticField(cardURL + "/idShort")
	c.ManualCoverAttachment = staticField(cardURL + "/manualCoverAttachment")
	c.MarkAssociatedNotificationsRead = staticField(cardURL + "/markAssociatedNotificationsRead")
	c.Members = staticField(cardURL + "/members")
	c.Name = staticField(cardURL + "/name")
	c.Pos = staticField(cardURL + "/pos")
	c.ShortLink = staticField(cardURL + "/shortLink")
	c.ShortUrl = staticField(cardURL + "/shortUrl")
	c.Subscribed = staticField(cardURL + "/subscribed")
	c.Url = staticField(cardURL + "/url")
	return c
}

func (c baseCard) getURL() string {
	return c.url
}

type card struct {
	baseCard
	url string
	//submodels
	Actions      baseAction
	Attachments  attachment
	Board        baseBoard
	Checklist    checklist
	Checklists   checklists
	IdLabels     idLabel
	IdMembers    idMember
	Labels       label
	List         list
	MembersVoted membersVoted
	Stickers     blankPlaceholder
}

func (c card) ID(id string) card {
	cardURL := c.url + "/" + id
	c.url = cardURL
	c.Badges = staticField(cardURL + "/badges")
	c.CheckItemStates = staticField(cardURL + "/checkItemStates")
	c.Closed = staticField(cardURL + "/closed")
	c.DateLastActivity = staticField(cardURL + "/dateLastActivity")
	c.Desc = staticField(cardURL + "/desc")
	c.DescData = staticField(cardURL + "/descData")
	c.Due = staticField(cardURL + "/due")
	c.Email = staticField(cardURL + "/email")
	c.IdAttachmentCover = staticField(cardURL + "/idAttachmentCover")
	c.IdBoard = staticField(cardURL + "/idBoard")
	c.IdChecklist = staticField(cardURL + "/idChecklist")
	c.IdChecklists = staticField(cardURL + "/idChecklists")
	c.IdList = staticField(cardURL + "/idList")
	c.IdMembersVoted = staticField(cardURL + "/idMembersVoted")
	c.IdShort = staticField(cardURL + "/idShort")
	c.ManualCoverAttachment = staticField(cardURL + "/manualCoverAttachment")
	c.MarkAssociatedNotificationsRead = staticField(cardURL + "/markAssociatedNotificationsRead")
	c.Members = staticField(cardURL + "/members")
	c.Name = staticField(cardURL + "/name")
	c.Pos = staticField(cardURL + "/pos")
	c.ShortLink = staticField(cardURL + "/shortLink")
	c.ShortUrl = staticField(cardURL + "/shortUrl")
	c.Subscribed = staticField(cardURL + "/subscribed")
	c.Url = staticField(cardURL + "/url")
	c.Actions = createBaseAction(c)
	c.Attachments = createAttachment(c)
	c.Board = createBaseBoard(c)
	c.Checklists = createChecklists(c)
	c.Checklist = createChecklist(c)
	c.IdLabels = createIdLabel(c)
	c.IdMembers = createIdMember(c)
	c.Labels = createLabel(c)
	c.MembersVoted = createMembersVoted(c)
	c.Stickers = createBlankPlaceholder(c, "stickers")
	c.List = createList(c)
	return c
}

func (c card) getURL() string {
	return c.url
}

var Cards = card{
	url: "/cards",
}
