package trello

type notificationAll struct {
	url  string
	Read staticField
}

func createNotificationAll(m model) notificationAll {
	nURL := m.getURL() + "/all"
	return notificationAll{
		url:  nURL,
		Read: staticField(nURL + "/read"),
	}
}

type notifications struct {
	url string
	Data,
	Date,
	IdMemberCreator,
	Type,
	Display,
	Entities,
	Read,
	Unread staticField
	Board         baseBoard
	Card          baseCard
	List          list
	Member        baseMember
	MemberCreator memberCreator
	Organization  organization
	All           notificationAll
}

func createNotification(m model) notifications {
	nURL := m.getURL() + "/notifications"
	n := notifications{
		url: nURL,
	}
	n.Data = staticField(nURL + "/data")
	n.Date = staticField(nURL + "/date")
	n.IdMemberCreator = staticField(nURL + "/idMemberCreator")
	n.Type = staticField(nURL + "/type")
	n.Display = staticField(nURL + "/display")
	n.Entities = staticField(nURL + "/entities")
	n.Read = staticField(nURL + "/read")
	n.Unread = staticField(nURL + "/unread")
	n.All = createNotificationAll(m)
	return n
}

func (n notifications) ID(id string) notifications {
	nURL := n.url + "/" + id
	return notifications{
		url:             nURL,
		Data:            staticField(nURL + "/data"),
		Date:            staticField(nURL + "/date"),
		IdMemberCreator: staticField(nURL + "/idMemberCreator"),
		Type:            staticField(nURL + "/type"),
		Display:         staticField(nURL + "/display"),
		Entities:        staticField(nURL + "/entities"),
		Read:            staticField(nURL + "/read"),
		Unread:          staticField(nURL + "/unread"),
		Board:           createBaseBoard(n),
		Card:            createBaseCard(n),
		List:            createList(n),
		Member:          createBaseMember(n),
		MemberCreator:   createMemberCreator(n),
		Organization:    createOrganization(n),
		All:             createNotificationAll(n),
	}
}

func (n notifications) getURL() string {
	return n.url
}

var Notifications = notifications{
	url: "/notifications",
}
