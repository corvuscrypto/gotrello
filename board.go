package trello

type membersInvited struct {
	member
	url string
}

type memberships struct {
	url string
}

func createMemberships(m model) memberships {
	return memberships{
		url: m.getURL() + "/memberships",
	}
}

func (m memberships) ID(id string) memberships {
	m.url = m.url + "/" + id
	return m
}

type myPrefs struct {
	url string
	EmailPosition,
	IdEmailList,
	ShowListGuide,
	ShowSidebar,
	ShowSidebarActivity,
	ShowSidebarBoardActions,
	ShowSidebarMembers staticField
}

func createMyPrefs(m model) myPrefs {
	mpURL := m.getURL() + "/myPrefs"
	return myPrefs{
		url:                     mpURL,
		EmailPosition:           staticField(mpURL + "/emailPosition"),
		IdEmailList:             staticField(mpURL + "/idEmailList"),
		ShowListGuide:           staticField(mpURL + "/showListGuide"),
		ShowSidebar:             staticField(mpURL + "/showSidebar"),
		ShowSidebarActivity:     staticField(mpURL + "/showSidebarActivity"),
		ShowSidebarBoardActions: staticField(mpURL + "/showSidebarBoardActions"),
		ShowSidebarMembers:      staticField(mpURL + "/showSidebarMembers"),
	}
}

type boardPrefs struct {
	url string
	Background,
	CalendarFeedEnabled,
	CardAging,
	CardCovers,
	Comments,
	Invitations,
	PermissionLevel,
	SelfJoin,
	Voting staticField
}

func createBoardPrefs(m model) boardPrefs {
	pURL := m.getURL() + "/prefs"
	return boardPrefs{
		url:                 pURL,
		Background:          staticField(pURL + "/background"),
		CalendarFeedEnabled: staticField(pURL + "/calendarFeedEnabled"),
		CardAging:           staticField(pURL + "/cardAging"),
		CardCovers:          staticField(pURL + "/cardCovers"),
		Comments:            staticField(pURL + "/comments"),
		Invitations:         staticField(pURL + "/invitations"),
		PermissionLevel:     staticField(pURL + "/permissionLevel"),
		SelfJoin:            staticField(pURL + "/selfJoin"),
		Voting:              staticField(pURL + "/voting"),
	}
}

type powerUps struct {
	url string
	Calendar,
	CardAging,
	Recap,
	Voting staticField
}

func createPowerUps(m model) powerUps {
	pURL := m.getURL() + "/powerUps"
	return powerUps{
		url:       pURL,
		Calendar:  staticField(pURL + "/calendar"),
		CardAging: staticField(pURL + "/cardAging"),
		Recap:     staticField(pURL + "/recap"),
		Voting:    staticField(pURL + "/voting"),
	}
}

type baseBoard struct {
	url string
	//static fields
	Actions,
	BoardStars,
	Checklists,
	Closed,
	DateLastActivity,
	DateLastView,
	Deltas,
	Desc,
	DescData,
	IdOrganization,
	Invitations,
	Invited,
	Name,
	Pinned,
	ShortLink,
	ShortUrl,
	Starred,
	Subscribed,
	Url staticField
}

func createBaseBoard(m model) baseBoard {
	bURL := m.getURL() + "/board"
	return baseBoard{
		url:              bURL,
		Actions:          staticField(bURL + "/actions"),
		BoardStars:       staticField(bURL + "/boardstars"),
		Checklists:       staticField(bURL + "/checklists"),
		Closed:           staticField(bURL + "/closed"),
		DateLastActivity: staticField(bURL + "/dateLastActivity"),
		DateLastView:     staticField(bURL + "/dateLastView"),
		Deltas:           staticField(bURL + "/deltas"),
		Desc:             staticField(bURL + "/desc"),
		DescData:         staticField(bURL + "/descData"),
		IdOrganization:   staticField(bURL + "/idOrganization"),
		Invitations:      staticField(bURL + "/invitations"),
		Invited:          staticField(bURL + "/invited"),
		Name:             staticField(bURL + "/name"),
		Pinned:           staticField(bURL + "/pinned"),
		ShortLink:        staticField(bURL + "/shortLink"),
		ShortUrl:         staticField(bURL + "/shortUrl"),
		Starred:          staticField(bURL + "/starred"),
		Subscribed:       staticField(bURL + "/subscribed"),
		Url:              staticField(bURL + "/url"),
	}
}

type board struct {
	baseBoard
	url string
	//submodels
	CalendarKey    calKey
	EmailKey       emailKey
	Cards          card
	Labels         label
	LabelNames     labelNames
	Lists          list
	Members        member
	MembersInvited membersInvited
	Memberships    memberships
	MyPrefs        myPrefs
	Organization   organization
	PowerUps       powerUps
	Prefs          boardPrefs
}

//Boards is the struct representing the boards model from trello
var Boards = board{
	url: "/boards",
}

func createBoard(m model) board {
	return board{
		url: m.getURL() + "/boards",
	}
}

func (b board) ID(id string) board {
	boardURL := b.url + "/" + id
	b.url = boardURL
	b.Actions = staticField(boardURL + "/actions")
	b.BoardStars = staticField(boardURL + "/boardstars")
	b.Checklists = staticField(boardURL + "/checklists")
	b.Closed = staticField(boardURL + "/closed")
	b.DateLastActivity = staticField(boardURL + "/dateLastActivity")
	b.DateLastView = staticField(boardURL + "/dateLastView")
	b.Deltas = staticField(boardURL + "/deltas")
	b.Desc = staticField(boardURL + "/desc")
	b.DescData = staticField(boardURL + "/descData")
	b.IdOrganization = staticField(boardURL + "/idOrganization")
	b.Invitations = staticField(boardURL + "/invitations")
	b.Invited = staticField(boardURL + "/invited")
	b.Name = staticField(boardURL + "/name")
	b.Pinned = staticField(boardURL + "/pinned")
	b.ShortLink = staticField(boardURL + "/shortLink")
	b.ShortUrl = staticField(boardURL + "/shortUrl")
	b.Starred = staticField(boardURL + "/starred")
	b.Subscribed = staticField(boardURL + "/subscribed")
	b.Url = staticField(boardURL + "/url")
	return b
}

func (b board) getURL() string {
	return b.url
}
