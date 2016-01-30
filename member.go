package trello

type boardsInvited struct {
	baseBoard
	url string
}

func createBoardsInvited(m model) boardsInvited {
	base := createBaseBoard(m)
	b := boardsInvited{
		baseBoard: base,
	}
	bURL := m.getURL() + "/boardsInvited"
	b.url = bURL
	b.Actions = staticField(bURL + "/actions")
	b.BoardStars = staticField(bURL + "/boardStars")
	b.Checklists = staticField(bURL + "/checklists")
	b.Closed = staticField(bURL + "/closed")
	b.DateLastActivity = staticField(bURL + "/dateLastActivity")
	b.DateLastView = staticField(bURL + "/dateLastView")
	b.Deltas = staticField(bURL + "/deltas")
	b.Desc = staticField(bURL + "/desc")
	b.DescData = staticField(bURL + "/descData")
	b.IdOrganization = staticField(bURL + "/idOrganization")
	b.Invitations = staticField(bURL + "/invitations")
	b.Invited = staticField(bURL + "/invited")
	b.MarkAsViewed = staticField(bURL + "/markAsViewed")
	b.Name = staticField(bURL + "/name")
	b.Pinned = staticField(bURL + "/pinned")
	b.ShortLink = staticField(bURL + "/shortLink")
	b.ShortUrl = staticField(bURL + "/shortUrl")
	b.Starred = staticField(bURL + "/starred")
	b.Subscribed = staticField(bURL + "/subscribed")
	b.Url = staticField(bURL + "/url")
	b.LabelNames = createLabelNames(b)
	b.Memberships = createMemberships(b)
	b.PowerUps = createPowerUps(b)
	b.Prefs = createBoardPrefs(b)
	return b
	return b
}

func (b boardsInvited) getURL() string {
	return b.url
}

type memberNotifications struct {
	url string
	AddAdminToBoard,
	AddAdminToOrganization,
	AddedAttachmentToCard,
	AddedMemberToCard,
	AddedToBoard,
	AddedToCard,
	AddedToOrganization,
	CardDueSoon,
	ChangeCard,
	CloseBoard,
	CommentCard,
	CreatedCard,
	DeclinedInvitationToBoard,
	DeclinedInvitationToOrganization,
	InvitedToBoard,
	InvitedToOrganization,
	MakeAdminOfBoard,
	MakeAdminOfOrganization,
	MemberJoinedTrello,
	MentionedOnCard,
	RemovedFromBoard,
	RemovedFromCard,
	RemovedFromOrganization,
	RemovedMemberFromCard,
	UnconfirmedInvitedToBoard,
	UnconfirmedInvitedToOrganization,
	UpdateCheckItemStateOnCard staticField
}

func createMemberNotifications(m model) memberNotifications {
	mURL := m.getURL() + "/notifications"
	return memberNotifications{
		url:                              mURL,
		AddAdminToBoard:                  staticField(mURL + "/addAdminToBoard"),
		AddAdminToOrganization:           staticField(mURL + "/addAdminToOrganization"),
		AddedAttachmentToCard:            staticField(mURL + "/addedAttachmentToCard"),
		AddedMemberToCard:                staticField(mURL + "/addedMemberToCard"),
		AddedToBoard:                     staticField(mURL + "/addedToBoard"),
		AddedToCard:                      staticField(mURL + "/addedToCard"),
		AddedToOrganization:              staticField(mURL + "/addedToOrganization"),
		CardDueSoon:                      staticField(mURL + "/cardDueSoon"),
		ChangeCard:                       staticField(mURL + "/changeCard"),
		CloseBoard:                       staticField(mURL + "/closeBoard"),
		CommentCard:                      staticField(mURL + "/commentCard"),
		CreatedCard:                      staticField(mURL + "/createdCard"),
		DeclinedInvitationToBoard:        staticField(mURL + "/declinedInvitationToBoard"),
		DeclinedInvitationToOrganization: staticField(mURL + "/declinedInvitationToOrganization"),
		InvitedToBoard:                   staticField(mURL + "/invitedToBoard"),
		InvitedToOrganization:            staticField(mURL + "/invitedToOrganization"),
		MakeAdminOfBoard:                 staticField(mURL + "/makeAdminOfBoard"),
		MakeAdminOfOrganization:          staticField(mURL + "/makeAdminOfOrganization"),
		MemberJoinedTrello:               staticField(mURL + "/memberJoinedTrello"),
		MentionedOnCard:                  staticField(mURL + "/mentionedOnCard"),
		RemovedFromBoard:                 staticField(mURL + "/removedFromBoard"),
		RemovedFromCard:                  staticField(mURL + "/removedFromCard"),
		RemovedFromOrganization:          staticField(mURL + "/removedFromOrganization"),
		RemovedMemberFromCard:            staticField(mURL + "/removedMemberFromCard"),
		UnconfirmedInvitedToBoard:        staticField(mURL + "/unconfirmedInvitedToBoard"),
		UnconfirmedInvitedToOrganization: staticField(mURL + "/unconfirmedInvitedToOrganization"),
		UpdateCheckItemStateOnCard:       staticField(mURL + "/updateCheckItemStateOnCard"),
	}
}

func (m memberNotifications) getURL() string {
	return m.url
}

type baseMember struct {
	url string
	Admins,
	All,
	AvatarHash,
	AvatarSource,
	Bio,
	BioData,
	Confirmed,
	Deactivated,
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
	None,
	Normal,
	OneTimeMessagesDismissed,
	Owners,
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
	Cards,
	Avatar staticField
	Prefs memberPrefs
}

func createBaseMember(m model) baseMember {
	mURL := m.getURL()
	bm := baseMember{}
	switch m.(type) {
	case action, token, notifications:
		mURL += "/member"
		break
	default:
		mURL += "/members"
	}
	bm.url = mURL
	bm.Admins = staticField(mURL + "/admins")
	bm.All = staticField(mURL + "/all")
	bm.AvatarHash = staticField(mURL + "/avatarHash")
	bm.AvatarSource = staticField(mURL + "/avatarSource")
	bm.Bio = staticField(mURL + "/bio")
	bm.BioData = staticField(mURL + "/bioData")
	bm.Confirmed = staticField(mURL + "/confirmed")
	bm.Email = staticField(mURL + "/email")
	bm.FullName = staticField(mURL + "/fullName")
	bm.GravatarHash = staticField(mURL + "/gravatarHash")
	bm.IdBoards = staticField(mURL + "/idBoards")
	bm.IdBoardsPinned = staticField(mURL + "/idBoardsPinned")
	bm.IdOrganizations = staticField(mURL + "/idOrganizations")
	bm.IdPremOrgsAdmin = staticField(mURL + "/idPremOrgsAdmin")
	bm.Initials = staticField(mURL + "/initials")
	bm.LoginTypes = staticField(mURL + "/loginTypes")
	bm.MemberType = staticField(mURL + "/memberType")
	bm.None = staticField(mURL + "/none")
	bm.Normal = staticField(mURL + "/normal")
	bm.OneTimeMessagesDismissed = staticField(mURL + "/oneTimeMessagesDismissed")
	bm.Owners = staticField(mURL + "/owners")
	bm.PremiumFeatures = staticField(mURL + "/premiumFeatures")
	bm.Prefs = createMemberPrefs(bm)
	bm.Products = staticField(mURL + "/products")
	bm.Status = staticField(mURL + "/status")
	bm.Trophies = staticField(mURL + "/trophies")
	bm.UploadedAvatarHash = staticField(mURL + "/uploadedAvatarHash")
	bm.Url = staticField(mURL + "/url")
	bm.Username = staticField(mURL + "/username")
	bm.Actions = staticField(mURL + "/actions")
	bm.Deltas = staticField(mURL + "/deltas")
	bm.Tokens = staticField(mURL + "/tokens")
	bm.Avatar = staticField(mURL + "/avatar")
	return bm
}

func (m baseMember) ID(id string) baseMember {
	mURL := m.url + "/" + id
	m.url = mURL

	m.Admins = staticField(mURL + "/admins")
	m.All = staticField(mURL + "/all")
	m.AvatarHash = staticField(mURL + "/avatarHash")
	m.AvatarSource = staticField(mURL + "/avatarSource")
	m.Bio = staticField(mURL + "/bio")
	m.BioData = staticField(mURL + "/bioData")
	m.Confirmed = staticField(mURL + "/confirmed")
	m.Deactivated = staticField(mURL + "/deactivated")
	m.Email = staticField(mURL + "/email")
	m.FullName = staticField(mURL + "/fullName")
	m.GravatarHash = staticField(mURL + "/gravatarHash")
	m.IdBoards = staticField(mURL + "/idBoards")
	m.IdBoardsPinned = staticField(mURL + "/idBoardsPinned")
	m.IdOrganizations = staticField(mURL + "/idOrganizations")
	m.IdPremOrgsAdmin = staticField(mURL + "/idPremOrgsAdmin")
	m.Initials = staticField(mURL + "/initials")
	m.LoginTypes = staticField(mURL + "/loginTypes")
	m.MemberType = staticField(mURL + "/memberTypes")
	m.None = staticField(mURL + "/none")
	m.Normal = staticField(mURL + "/normal")
	m.OneTimeMessagesDismissed = staticField(mURL + "/oneTimeMessagesDismissed")
	m.Owners = staticField(mURL + "/owners")
	m.PremiumFeatures = staticField(mURL + "/premiumFeatures")
	m.Prefs = createMemberPrefs(m)
	m.Products = staticField(mURL + "/products")
	m.Status = staticField(mURL + "/status")
	m.Trophies = staticField(mURL + "/trophies")
	m.UploadedAvatarHash = staticField(mURL + "/uploadedAvatarHash")
	m.Url = staticField(mURL + "/url")
	m.Username = staticField(mURL + "/username")
	m.Actions = staticField(mURL + "/actions")
	m.Deltas = staticField(mURL + "/deltas")
	m.Tokens = staticField(mURL + "/tokens")
	m.Avatar = staticField(mURL + "/avatar")
	m.Cards = staticField(mURL + "/cards")
	return m
}

func (m baseMember) getURL() string {
	return m.url
}

type memberPrefs struct {
	url string
	ColorBlind,
	Locale,
	MinutesBetweenSummaries staticField
}

func createMemberPrefs(m model) memberPrefs {
	mURL := m.getURL() + "/prefs"
	return memberPrefs{
		url:                     mURL,
		ColorBlind:              staticField(mURL + "/colorBlind"),
		Locale:                  staticField(mURL + "/locale"),
		MinutesBetweenSummaries: staticField(mURL + "/minutesBetweenSummaries"),
	}
}

func (m memberPrefs) getURL() string {
	return m.url
}

type savedSearches struct {
	url string
	Name,
	Pos,
	Query staticField
}

func createSavedSearches(m model) savedSearches {
	sURL := m.getURL() + "/savedSearches"
	return savedSearches{
		url: sURL,
	}
}

func (s savedSearches) ID(id string) savedSearches {
	sURL := s.url + "/" + id
	return savedSearches{
		url:   sURL,
		Name:  staticField(sURL + "/name"),
		Pos:   staticField(sURL + "/pos"),
		Query: staticField(sURL + "/query"),
	}
}

func (s savedSearches) getURL() string {
	return s.url
}

type organizationsInvited struct {
	organization
	url string
}

func createOrganizationsInvited(m model) organizationsInvited {
	oURL := m.getURL() + "/organizationsInvited"
	o := organizationsInvited{
		organization: createOrganization(m),
	}
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
	o.Memberships = createBlankPlaceholder(o, "memberships")
	o.Prefs = createOrganizationPrefs(o)
	return o
}

func (o organizationsInvited) getURL() string {
	return o.url
}

type member struct {
	baseMember
	url                    string
	Cards                  baseCard
	BoardBackgrounds       blankPlaceholder
	Boards                 filterBoards
	BoardStars             boardStars
	BoardsInvited          boardsInvited
	CustomBoardBackgrounds blankPlaceholder
	CustomEmoji            blankPlaceholder
	CustomStickers         blankPlaceholder
	Notifications          memberNotifications
	Organizations          filterOrganization
	OrganizationsInvited   organizationsInvited
	SavedSearches          savedSearches
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
		baseMember: createBaseMember(m),
		url:        mURL,
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
	m.Cards = createBaseCard(m)
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
	m.Prefs = createMemberPrefs(m)
	m.PremiumFeatures = staticField(memberURL + "/premiumFeatures")
	m.Products = staticField(memberURL + "/products")
	m.Status = staticField(memberURL + "/status")
	m.Tokens = staticField(memberURL + "/tokens")
	m.Trophies = staticField(memberURL + "/trophies")
	m.UploadedAvatarHash = staticField(memberURL + "/uploadedAvatarHash")
	m.Url = staticField(memberURL + "/url")
	m.Username = staticField(memberURL + "/username")
	m.BoardBackgrounds = createBlankPlaceholder(m, "boardBackgrounds")
	m.Boards = createFilterBoard(m)
	m.BoardStars = createBoardStars(m)
	m.BoardsInvited = createBoardsInvited(m)
	m.CustomBoardBackgrounds = createBlankPlaceholder(m, "customBoardBackgrounds")
	m.CustomEmoji = createBlankPlaceholder(m, "customEmoji")
	m.CustomStickers = createBlankPlaceholder(m, "customStickers")
	m.Notifications = createMemberNotifications(m)
	m.Organizations = createFilterOrganization(m)
	m.OrganizationsInvited = createOrganizationsInvited(m)
	m.SavedSearches = createSavedSearches(m)
	return m
}

func (m member) getURL() string {
	return m.url
}

var Members = member{
	url: "/members",
}
