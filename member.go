package trello

type boardsInvited struct {
	baseBoard
	url string
}

func createBoardsInvited(m model) boardsInvited {
	base := createBaseBoard(m)
	return boardsInvited{
		base,
		m.getURL() + "boardsInvited",
	}
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
		AddedMemberToCard:                staticField(mURL + "/AddedMemberToCard"),
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
		UnconfirmedInvitedToBoard:        staticField(mURL + "/uncomfirmedInvitedToBoard"),
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
	mURL := m.getURL() + "/members"
	return baseMember{
		Admins:          staticField(mURL + "/admins"),
		All:             staticField(mURL + "/all"),
		AvatarHash:      staticField(mURL + "/avatarHash"),
		AvatarSource:    staticField(mURL + "/avatarSource"),
		Bio:             staticField(mURL + "/bio"),
		BioData:         staticField(mURL + "/bioData"),
		Confirmed:       staticField(mURL + "/confirmed"),
		Email:           staticField(mURL + "/email"),
		FullName:        staticField(mURL + "/fullName"),
		GravatarHash:    staticField(mURL + "/gravatarHash"),
		IdBoards:        staticField(mURL + "/idBoards"),
		IdBoardsPinned:  staticField(mURL + "/idBoardsPinned"),
		IdOrganizations: staticField(mURL + "/idOrganizations"),
		IdPremOrgsAdmin: staticField(mURL + "/idPremOrgsAdmin"),
		Initials:        staticField(mURL + "/initials"),
		LoginTypes:      staticField(mURL + "/loginTypes"),
		MemberType:      staticField(mURL + "/memberTypes"),
		None:            staticField(mURL + "/none"),
		Normal:          staticField(mURL + "/normal"),
		OneTimeMessagesDismissed: staticField(mURL + "/oneTimeMessagesDismissed"),
		Owners:             staticField(mURL + "/owners"),
		PremiumFeatures:    staticField(mURL + "/premiumFeatures"),
		Prefs:              createMemberPrefs(m),
		Products:           staticField(mURL + "/products"),
		Status:             staticField(mURL + "/status"),
		Trophies:           staticField(mURL + "/trophies"),
		UploadedAvatarHash: staticField(mURL + "/uploadedAvatarHash"),
		Url:                staticField(mURL + "/url"),
		Username:           staticField(mURL + "/username"),
		Actions:            staticField(mURL + "/actions"),
		Deltas:             staticField(mURL + "/deltas"),
		Tokens:             staticField(mURL + "/tokens"),
		Avatar:             staticField(mURL + "/avatar"),
	}
}

func (m baseMember) ID(id string) baseMember {
	mURL := m.url + "/" + id
	return baseMember{
		Admins:          staticField(mURL + "/admins"),
		All:             staticField(mURL + "/all"),
		AvatarHash:      staticField(mURL + "/avatarHash"),
		AvatarSource:    staticField(mURL + "/avatarSource"),
		Bio:             staticField(mURL + "/bio"),
		BioData:         staticField(mURL + "/bioData"),
		Confirmed:       staticField(mURL + "/confirmed"),
		Deactivated:     staticField(mURL + "/deactivated"),
		Email:           staticField(mURL + "/email"),
		FullName:        staticField(mURL + "/fullName"),
		GravatarHash:    staticField(mURL + "/gravatarHash"),
		IdBoards:        staticField(mURL + "/idBoards"),
		IdBoardsPinned:  staticField(mURL + "/idBoardsPinned"),
		IdOrganizations: staticField(mURL + "/idOrganizations"),
		IdPremOrgsAdmin: staticField(mURL + "/idPremOrgsAdmin"),
		Initials:        staticField(mURL + "/initials"),
		LoginTypes:      staticField(mURL + "/loginTypes"),
		MemberType:      staticField(mURL + "/memberTypes"),
		None:            staticField(mURL + "/none"),
		Normal:          staticField(mURL + "/normal"),
		OneTimeMessagesDismissed: staticField(mURL + "/oneTimeMessagesDismissed"),
		Owners:             staticField(mURL + "/owners"),
		PremiumFeatures:    staticField(mURL + "/premiumFeatures"),
		Prefs:              createMemberPrefs(m),
		Products:           staticField(mURL + "/products"),
		Status:             staticField(mURL + "/status"),
		Trophies:           staticField(mURL + "/trophies"),
		UploadedAvatarHash: staticField(mURL + "/uploadedAvatarHash"),
		Url:                staticField(mURL + "/url"),
		Username:           staticField(mURL + "/username"),
		Actions:            staticField(mURL + "/actions"),
		Deltas:             staticField(mURL + "/deltas"),
		Tokens:             staticField(mURL + "/tokens"),
		Avatar:             staticField(mURL + "/avatar"),
		Cards:              staticField(mURL + "/cards"),
	}
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
		Locale:                  staticField(mURL + "/local"),
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
	OrganizationsInvited   organization
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
	m.OrganizationsInvited = createOrganization(m)
	m.SavedSearches = createSavedSearches(m)
	return m
}

func (m member) getURL() string {
	return m.url
}

var Members = member{
	url: "/members",
}
