package trello

type boardVisibilityRestrict struct {
	Org,
	Private,
	Public staticField
}

type organizationPrefs struct {
	url string
	AssociatedDomain,
	ExternalMembersDisabled,
	GoogleAppsVersion,
	OrgInviteRestrict,
	PermissionLevel staticField
	BoardVisibilityRestrict boardVisibilityRestrict
}

func createOrganizationPrefs(m model) organizationPrefs {
	oURL := m.getURL() + "/prefs"
	return organizationPrefs{
		url:                     oURL,
		AssociatedDomain:        staticField(oURL + "/associatedDomain"),
		ExternalMembersDisabled: staticField(oURL + "/externalMembersDisabled"),
		GoogleAppsVersion:       staticField(oURL + "/googleAppsVersion"),
		OrgInviteRestrict:       staticField(oURL + "/orgInviteRestrict"),
		PermissionLevel:         staticField(oURL + "/permissionLevel"),
		BoardVisibilityRestrict: boardVisibilityRestrict{
			Org:     staticField(oURL + "/boardVisibilityRestrict/org"),
			Private: staticField(oURL + "/boardVisibilityRestrict/private"),
			Public:  staticField(oURL + "/boardVisibilityRestrict/public"),
		},
	}
}

func (o organizationPrefs) getURL() string {
	return o.url
}

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
	Deltas,
	DisplayName,
	IdBoards,
	Invitations,
	Invited,
	Logo,
	LogoHash,
	Name,
	PowerUps,
	PremiumFeatures,
	Products,
	Url,
	Website staticField
	Actions        baseAction
	Boards         filterBoards
	Members        baseMember
	MembersInvited membersInvited
	Memberships    blankPlaceholder
	Prefs          organizationPrefs
}

func createOrganization(m model) organization {
	var oURL = m.getURL()
	switch m.(type) {
	case action, notifications:
		oURL += "/organization"
		break
	default:
		oURL += "/organizations"
	}
	o := organization{}
	o.url = oURL
	o.BillableMemberCount = staticField(oURL + "/billableMemberCount")
	o.Desc = staticField(oURL + "/desc")
	o.DescData = staticField(oURL + "/descData")
	o.Deltas = staticField(oURL + "/deltas")
	o.DisplayName = staticField(oURL + "/displayName")
	o.IdBoards = staticField(oURL + "/idBoards")
	o.Invitations = staticField(oURL + "/invitations")
	o.Invited = staticField(oURL + "/invited")
	o.Logo = staticField(oURL + "/logo")
	o.LogoHash = staticField(oURL + "/logoHash")
	o.Name = staticField(oURL + "/name")
	o.PowerUps = staticField(oURL + "/powerUps")
	o.PremiumFeatures = staticField(oURL + "/premiumFeatures")
	o.Products = staticField(oURL + "/products")
	o.Url = staticField(oURL + "/url")
	o.Website = staticField(oURL + "/website")
	o.MembersInvited = createMembersInvited(o)
	o.Memberships = createBlankPlaceholder(o, "memberships")
	o.Prefs = createOrganizationPrefs(o)
	return o
}

func (o organization) ID(id string) organization {
	orgURL := o.url + "/" + id
	o.url = orgURL
	o.BillableMemberCount = staticField(orgURL + "/billableMemberCount")
	o.Desc = staticField(orgURL + "/desc")
	o.DescData = staticField(orgURL + "/descData")
	o.Deltas = staticField(orgURL + "/deltas")
	o.DisplayName = staticField(orgURL + "/displayName")
	o.IdBoards = staticField(orgURL + "/idBoards")
	o.Invitations = staticField(orgURL + "/invitations")
	o.Invited = staticField(orgURL + "/invited")
	o.Logo = staticField(orgURL + "/logo")
	o.LogoHash = staticField(orgURL + "/logoHash")
	o.Name = staticField(orgURL + "/name")
	o.PowerUps = staticField(orgURL + "/powerUps")
	o.PremiumFeatures = staticField(orgURL + "/premiumFeatures")
	o.Products = staticField(orgURL + "/products")
	o.Url = staticField(orgURL + "/url")
	o.Website = staticField(orgURL + "/website")
	o.Actions = createBaseAction(o)
	o.Boards = createFilterBoard(o)
	o.MembersInvited = createMembersInvited(o)
	o.Members = createBaseMember(o)
	o.Memberships = createBlankPlaceholder(o, "memberships")
	o.Prefs = createOrganizationPrefs(o)
	return o
}

func (o organization) getURL() string {
	return o.url
}

var Organizations = organization{
	url: "/organizations",
}
