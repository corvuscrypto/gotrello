package trello

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
	PremiumFeatures,
	Products,
	Url,
	Website staticField
}

func createOrganization(m model) organization {
	return organization{
		url: m.getURL() + "/organizations",
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
	return o
}

func (o organization) getURL() string {
	return o.url
}
