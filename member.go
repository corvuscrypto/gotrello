package trello

type member struct {
	url string
	AvatarHash,
	AvatarSource,
	Bio,
	BioData,
	Confirmed,
	Email,
	FullName,
	GravatarHash,
	IdBoards,
	IdBoardsPinned,
	IdOrganizations,
	IdPremOrgsAdmin,
	Initials,
	LoginTypes,
	MemberType,
	OneTimeMessagesDismissed,
	PremiumFeatures,
	Products,
	Status,
	Trophies,
	UploadedAvatarHash,
	Url,
	Username,
	Actions,
	Deltas,
	Tokens,
	Avatar staticField
}

func createMember(m model) member {
	mURL := m.getURL()
	switch m.(type) {
	case action:
		mURL += "/member"
		break
	default:
		mURL += "/members"
	}
	return member{
		url: mURL,
	}
}

func (m member) ID(id string) member {
	memberURL := m.url + "/" + id
	m.url = memberURL
	m.Actions = staticField(memberURL + "/actions")
	m.Avatar = staticField(memberURL + "/avatar")
	m.AvatarHash = staticField(memberURL + "/avatarHash")
	m.AvatarSource = staticField(memberURL + "/avatarSource")
	m.Bio = staticField(memberURL + "/bio")
	m.BioData = staticField(memberURL + "/bioData")
	m.Confirmed = staticField(memberURL + "/confirmed")
	m.Deltas = staticField(memberURL + "/deltas")
	m.Email = staticField(memberURL + "/email")
	m.FullName = staticField(memberURL + "/fullName")
	m.GravatarHash = staticField(memberURL + "/gravatarHash")
	m.IdBoards = staticField(memberURL + "/idBoards")
	m.IdBoardsPinned = staticField(memberURL + "/idBoardsPinned")
	m.IdOrganizations = staticField(memberURL + "/idOrganizations")
	m.IdPremOrgsAdmin = staticField(memberURL + "/idPremOrgsAdmin")
	m.Initials = staticField(memberURL + "/initials")
	m.LoginTypes = staticField(memberURL + "/loginTypes")
	m.MemberType = staticField(memberURL + "/memberType")
	m.OneTimeMessagesDismissed = staticField(memberURL + "/oneTimeMessagesDismissed")
	m.PremiumFeatures = staticField(memberURL + "/premiumFeatures")
	m.Products = staticField(memberURL + "/products")
	m.Status = staticField(memberURL + "/status")
	m.Tokens = staticField(memberURL + "/tokens")
	m.Trophies = staticField(memberURL + "/trophies")
	m.UploadedAvatarHash = staticField(memberURL + "/uploadedAvatarHash")
	m.Url = staticField(memberURL + "/url")
	m.Username = staticField(memberURL + "/username")
	return m
}

func (m member) getURL() string {
	return m.url
}
