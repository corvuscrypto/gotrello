package trello

import "testing"

func testTokens(t *testing.T) {

	testModels := []model{
		Tokens.ID(ID),
		Tokens.ID(ID).DateCreated,
		Tokens.ID(ID).DateExpires,
		Tokens.ID(ID).IdMember,
		Tokens.ID(ID).Identifier,
		Tokens.ID(ID).Permissions,
		Tokens.ID(ID).Member,
		Tokens.ID(ID).Member.AvatarHash,
		Tokens.ID(ID).Member.AvatarSource,
		Tokens.ID(ID).Member.Bio,
		Tokens.ID(ID).Member.BioData,
		Tokens.ID(ID).Member.Confirmed,
		Tokens.ID(ID).Member.Email,
		Tokens.ID(ID).Member.FullName,
		Tokens.ID(ID).Member.GravatarHash,
		Tokens.ID(ID).Member.IdBoards,
		Tokens.ID(ID).Member.IdBoardsPinned,
		Tokens.ID(ID).Member.IdOrganizations,
		Tokens.ID(ID).Member.IdPremOrgsAdmin,
		Tokens.ID(ID).Member.Initials,
		Tokens.ID(ID).Member.LoginTypes,
		Tokens.ID(ID).Member.MemberType,
		Tokens.ID(ID).Member.OneTimeMessagesDismissed,
		Tokens.ID(ID).Member.Prefs,
		Tokens.ID(ID).Member.PremiumFeatures,
		Tokens.ID(ID).Member.Products,
		Tokens.ID(ID).Member.Status,
		Tokens.ID(ID).Member.Trophies,
		Tokens.ID(ID).Member.UploadedAvatarHash,
		Tokens.ID(ID).Member.Url,
		Tokens.ID(ID).Member.Username,
		Tokens.ID(ID).Webhooks,
		Tokens.ID(ID).Webhooks.ID(ID),
	}

	testExpects := []string{
		"/tokens/" + ID,
		"/tokens/" + ID + "/dateCreated",
		"/tokens/" + ID + "/dateExpires",
		"/tokens/" + ID + "/idMember",
		"/tokens/" + ID + "/identifier",
		"/tokens/" + ID + "/permissions",
		"/tokens/" + ID + "/member",
		"/tokens/" + ID + "/member/avatarHash",
		"/tokens/" + ID + "/member/avatarSource",
		"/tokens/" + ID + "/member/bio",
		"/tokens/" + ID + "/member/bioData",
		"/tokens/" + ID + "/member/confirmed",
		"/tokens/" + ID + "/member/email",
		"/tokens/" + ID + "/member/fullName",
		"/tokens/" + ID + "/member/gravatarHash",
		"/tokens/" + ID + "/member/idBoards",
		"/tokens/" + ID + "/member/idBoardsPinned",
		"/tokens/" + ID + "/member/idOrganizations",
		"/tokens/" + ID + "/member/idPremOrgsAdmin",
		"/tokens/" + ID + "/member/initials",
		"/tokens/" + ID + "/member/loginTypes",
		"/tokens/" + ID + "/member/memberType",
		"/tokens/" + ID + "/member/oneTimeMessagesDismissed",
		"/tokens/" + ID + "/member/prefs",
		"/tokens/" + ID + "/member/premiumFeatures",
		"/tokens/" + ID + "/member/products",
		"/tokens/" + ID + "/member/status",
		"/tokens/" + ID + "/member/trophies",
		"/tokens/" + ID + "/member/uploadedAvatarHash",
		"/tokens/" + ID + "/member/url",
		"/tokens/" + ID + "/member/username",
		"/tokens/" + ID + "/webhooks",
		"/tokens/" + ID + "/webhooks/" + ID,
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
