package trello

type filterOrganization struct {
	url string
	All,
	Members,
	None,
	Public staticField
}

func createFilterOrganization(m model) filterOrganization {
	oURL := m.getURL() + "/organizations"
	return filterOrganization{
		url:     oURL,
		All:     staticField(oURL + "/all"),
		Members: staticField(oURL + "/members"),
		None:    staticField(oURL + "/none"),
		Public:  staticField(oURL + "/public"),
	}
}

func (o filterOrganization) getURL() string {
	return o.url
}

type organization struct {
	url string
	BillableMemberCount,
	Desc,
	DescData,
	DisplayName,
	IdBoards,
	Invitations,
	Invited,
	LogoHash,
	Memberships,
	Name,
	PowerUps,
	Prefs,
	PremiumFeatures,
	Products,
	Url,
	Website staticField
	// Actions action
	Boards filterBoards
}

func createOrganization(m model) organization {
	var oURL = m.getURL()
	switch m.(type) {
	case action:
		oURL += "/organization"
		break
	default:
		oURL += "/organizations"
	}
	return organization{
		url:                 oURL,
		BillableMemberCount: staticField(oURL + "/billableMemberCount"),
		Desc:                staticField(oURL + "/desc"),
		DescData:            staticField(oURL + "/descData"),
		DisplayName:         staticField(oURL + "/displayName"),
		IdBoards:            staticField(oURL + "/idBoards"),
		Invitations:         staticField(oURL + "/invitations"),
		Invited:             staticField(oURL + "/invited"),
		LogoHash:            staticField(oURL + "/logoHash"),
		Memberships:         staticField(oURL + "/memberships"),
		Name:                staticField(oURL + "/name"),
		PowerUps:            staticField(oURL + "/powerUps"),
		Prefs:               staticField(oURL + "/prefs"),
		PremiumFeatures:     staticField(oURL + "/premiumFeatures"),
		Products:            staticField(oURL + "/products"),
		Url:                 staticField(oURL + "/url"),
		Website:             staticField(oURL + "/website"),
	}
}

func (o organization) ID(id string) organization {
	orgURL := o.url + "/" + id
	o.url = orgURL
	o.BillableMemberCount = staticField(orgURL + "/billableMemberCount")
	o.Desc = staticField(orgURL + "/desc")
	o.DescData = staticField(orgURL + "/descData")
	o.DisplayName = staticField(orgURL + "/displayName")
	o.IdBoards = staticField(orgURL + "/idBoards")
	o.Invitations = staticField(orgURL + "/Invitations")
	o.Invited = staticField(orgURL + "/invited")
	o.LogoHash = staticField(orgURL + "/logoHash")
	o.Memberships = staticField(orgURL + "/memberships")
	o.Name = staticField(orgURL + "/name")
	o.PowerUps = staticField(orgURL + "/powerUps")
	o.PremiumFeatures = staticField(orgURL + "/premiumFeatures")
	o.Products = staticField(orgURL + "/products")
	o.Url = staticField(orgURL + "/url")
	o.Website = staticField(orgURL + "/website")
	// o.Actions = createAction(o) TODO: break the actions struct up into a base struct
	o.Boards = createFilterBoard(o)
	return o
}

func (o organization) getURL() string {
	return o.url
}

var Organizations = organization{
	url: "/organizations",
}
