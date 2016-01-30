package trello

type memberCreator struct {
	baseMember
	url string
}

func createMemberCreator(m model) memberCreator {
	mURL := m.getURL() + "/memberCreator"
	mc := memberCreator{
		url: mURL,
	}
	mc.Actions = staticField(mURL + "/actions")
	mc.Avatar = staticField(mURL + "/avatar")
	mc.AvatarHash = staticField(mURL + "/avatarHash")
	mc.AvatarSource = staticField(mURL + "/avatarSource")
	mc.Bio = staticField(mURL + "/bio")
	mc.BioData = staticField(mURL + "/bioData")
	mc.Confirmed = staticField(mURL + "/confirmed")
	mc.Deltas = staticField(mURL + "/deltas")
	mc.Email = staticField(mURL + "/email")
	mc.FullName = staticField(mURL + "/fullName")
	mc.GravatarHash = staticField(mURL + "/gravatarHash")
	mc.IdBoards = staticField(mURL + "/idBoards")
	mc.IdBoardsPinned = staticField(mURL + "/idBoardsPinned")
	mc.IdOrganizations = staticField(mURL + "/idOrganizations")
	mc.IdPremOrgsAdmin = staticField(mURL + "/idPremOrgsAdmin")
	mc.Initials = staticField(mURL + "/initials")
	mc.LoginTypes = staticField(mURL + "/loginTypes")
	mc.MemberType = staticField(mURL + "/memberType")
	mc.OneTimeMessagesDismissed = staticField(mURL + "/oneTimeMessagesDismissed")
	mc.Prefs = createMemberPrefs(mc)
	mc.PremiumFeatures = staticField(mURL + "/premiumFeatures")
	mc.Products = staticField(mURL + "/products")
	mc.Status = staticField(mURL + "/status")
	mc.Tokens = staticField(mURL + "/tokens")
	mc.Trophies = staticField(mURL + "/trophies")
	mc.UploadedAvatarHash = staticField(mURL + "/uploadedAvatarHash")
	mc.Url = staticField(mURL + "/url")
	mc.Username = staticField(mURL + "/username")
	return mc
}

func (m memberCreator) getURL() string {
	return m.url
}
