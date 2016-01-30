package trello

import "testing"

func testActions(t *testing.T) {
	a := Actions.ID(ID)
	testModels := []model{
		//base Action URLs
		Actions,
		a,
		//Action field URLs
		a.Comments,
		a.Data,
		a.Date,
		a.Display,
		a.Entities,
		a.IdMemberCreator,
		a.Text,
		a.Type,
		//Board URLs
		a.Board,
		a.Board.Closed,
		a.Board.DateLastActivity,
		a.Board.DateLastView,
		a.Board.Desc,
		a.Board.DescData,
		a.Board.IdOrganization,
		a.Board.Invitations,
		a.Board.Invited,
		a.Board.LabelNames,
		a.Board.Memberships,
		a.Board.Name,
		a.Board.Pinned,
		a.Board.PowerUps,
		a.Board.Prefs,
		a.Board.ShortLink,
		a.Board.ShortUrl,
		a.Board.Starred,
		a.Board.Subscribed,
		a.Board.Url,
		//Card URLs
		a.Card,
		a.Card.Badges,
		a.Card.CheckItemStates,
		a.Card.Closed,
		a.Card.DateLastActivity,
		a.Card.Desc,
		a.Card.DescData,
		a.Card.Due,
		a.Card.Email,
		a.Card.IdAttachmentCover,
		a.Card.IdBoard,
		a.Card.IdChecklist,
		a.Card.IdList,
		a.Card.IdMembersVoted,
		a.Card.IdShort,
		a.Card.ManualCoverAttachment,
		a.Card.Name,
		a.Card.Pos,
		a.Card.ShortLink,
		a.Card.ShortUrl,
		a.Card.Subscribed,
		a.Card.Url,
		//List URLs
		a.List,
		a.List.Closed,
		a.List.IdBoard,
		a.List.Name,
		a.List.Pos,
		a.List.Subscribed,
		//Member URLs
		a.Member,
		a.Member.AvatarHash,
		a.Member.AvatarSource,
		a.Member.Bio,
		a.Member.BioData,
		a.Member.Confirmed,
		a.Member.Email,
		a.Member.FullName,
		a.Member.GravatarHash,
		a.Member.IdBoards,
		a.Member.IdBoardsPinned,
		a.Member.IdOrganizations,
		a.Member.IdPremOrgsAdmin,
		a.Member.Initials,
		a.Member.LoginTypes,
		a.Member.MemberType,
		a.Member.OneTimeMessagesDismissed,
		a.Member.Prefs,
		a.Member.PremiumFeatures,
		a.Member.Products,
		a.Member.Status,
		a.Member.Trophies,
		a.Member.UploadedAvatarHash,
		a.Member.Url,
		a.Member.Username,
		//MemberCreator URLs
		a.MemberCreator,
		a.MemberCreator.AvatarHash,
		a.MemberCreator.AvatarSource,
		a.MemberCreator.Bio,
		a.MemberCreator.BioData,
		a.MemberCreator.Confirmed,
		a.MemberCreator.Email,
		a.MemberCreator.FullName,
		a.MemberCreator.GravatarHash,
		a.MemberCreator.IdBoards,
		a.MemberCreator.IdBoardsPinned,
		a.MemberCreator.IdOrganizations,
		a.MemberCreator.IdPremOrgsAdmin,
		a.MemberCreator.Initials,
		a.MemberCreator.LoginTypes,
		a.MemberCreator.MemberType,
		a.MemberCreator.OneTimeMessagesDismissed,
		a.MemberCreator.Prefs,
		a.MemberCreator.PremiumFeatures,
		a.MemberCreator.Products,
		a.MemberCreator.Status,
		a.MemberCreator.Trophies,
		a.MemberCreator.UploadedAvatarHash,
		a.MemberCreator.Url,
		a.MemberCreator.Username,
		//Organization URLs
		a.Organization,
		a.Organization.BillableMemberCount,
		a.Organization.Desc,
		a.Organization.DescData,
		a.Organization.DisplayName,
		a.Organization.IdBoards,
		a.Organization.Invitations,
		a.Organization.Invited,
		a.Organization.LogoHash,
		a.Organization.Memberships,
		a.Organization.Name,
		a.Organization.PowerUps,
		a.Organization.Prefs,
		a.Organization.PremiumFeatures,
		a.Organization.Products,
		a.Organization.Url,
		a.Organization.Website,
	}

	as := "/actions/" + ID
	testExpects := []string{
		//Base Action URLs
		"/actions",
		as,
		//Action field URLs
		as + "/comments",
		as + "/data",
		as + "/date",
		as + "/display",
		as + "/entities",
		as + "/idMemberCreator",
		as + "/text",
		as + "/type",
		//Board URLs
		as + "/board",
		as + "/board/closed",
		as + "/board/dateLastActivity",
		as + "/board/dateLastView",
		as + "/board/desc",
		as + "/board/descData",
		as + "/board/idOrganization",
		as + "/board/invitations",
		as + "/board/invited",
		as + "/board/labelNames",
		as + "/board/memberships",
		as + "/board/name",
		as + "/board/pinned",
		as + "/board/powerUps",
		as + "/board/prefs",
		as + "/board/shortLink",
		as + "/board/shortUrl",
		as + "/board/starred",
		as + "/board/subscribed",
		as + "/board/url",
		//Card URLs
		as + "/card",
		as + "/card/badges",
		as + "/card/checkItemStates",
		as + "/card/closed",
		as + "/card/dateLastActivity",
		as + "/card/desc",
		as + "/card/descData",
		as + "/card/due",
		as + "/card/email",
		as + "/card/idAttachmentCover",
		as + "/card/idBoard",
		as + "/card/idChecklist",
		as + "/card/idList",
		as + "/card/idMembersVoted",
		as + "/card/idShort",
		as + "/card/manualCoverAttachment",
		as + "/card/name",
		as + "/card/pos",
		as + "/card/shortLink",
		as + "/card/shortUrl",
		as + "/card/subscribed",
		as + "/card/url",
		//List URLs
		as + "/list",
		as + "/list/closed",
		as + "/list/idBoard",
		as + "/list/name",
		as + "/list/pos",
		as + "/list/subscribed",
		//Member URLs
		as + "/member",
		as + "/member/avatarHash",
		as + "/member/avatarSource",
		as + "/member/bio",
		as + "/member/bioData",
		as + "/member/confirmed",
		as + "/member/email",
		as + "/member/fullName",
		as + "/member/gravatarHash",
		as + "/member/idBoards",
		as + "/member/idBoardsPinned",
		as + "/member/idOrganizations",
		as + "/member/idPremOrgsAdmin",
		as + "/member/initials",
		as + "/member/loginTypes",
		as + "/member/memberType",
		as + "/member/oneTimeMessagesDismissed",
		as + "/member/prefs",
		as + "/member/premiumFeatures",
		as + "/member/products",
		as + "/member/status",
		as + "/member/trophies",
		as + "/member/uploadedAvatarHash",
		as + "/member/url",
		as + "/member/username",
		//MemberCreator URLs
		as + "/memberCreator",
		as + "/memberCreator/avatarHash",
		as + "/memberCreator/avatarSource",
		as + "/memberCreator/bio",
		as + "/memberCreator/bioData",
		as + "/memberCreator/confirmed",
		as + "/memberCreator/email",
		as + "/memberCreator/fullName",
		as + "/memberCreator/gravatarHash",
		as + "/memberCreator/idBoards",
		as + "/memberCreator/idBoardsPinned",
		as + "/memberCreator/idOrganizations",
		as + "/memberCreator/idPremOrgsAdmin",
		as + "/memberCreator/initials",
		as + "/memberCreator/loginTypes",
		as + "/memberCreator/memberType",
		as + "/memberCreator/oneTimeMessagesDismissed",
		as + "/memberCreator/prefs",
		as + "/memberCreator/premiumFeatures",
		as + "/memberCreator/products",
		as + "/memberCreator/status",
		as + "/memberCreator/trophies",
		as + "/memberCreator/uploadedAvatarHash",
		as + "/memberCreator/url",
		as + "/memberCreator/username",
		//Organization URLs
		as + "/organization",
		as + "/organization/billableMemberCount",
		as + "/organization/desc",
		as + "/organization/descData",
		as + "/organization/displayName",
		as + "/organization/idBoards",
		as + "/organization/invitations",
		as + "/organization/invited",
		as + "/organization/logoHash",
		as + "/organization/memberships",
		as + "/organization/name",
		as + "/organization/powerUps",
		as + "/organization/prefs",
		as + "/organization/premiumFeatures",
		as + "/organization/products",
		as + "/organization/url",
		as + "/organization/website",
	}

	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}

}
