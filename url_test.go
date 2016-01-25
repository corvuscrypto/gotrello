package trello

import "testing"

var ID = "4f53aed56525abf3e12e2e45"

func TestGetURL(t *testing.T) {
	// test for Action
	testActions(t)
}

func testActions(t *testing.T) {
	a := Actions.ID(ID)
	testModels := []model{
		Actions,
		a,
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
		a.Card.IdChecklists,
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
		a.Comments,
		a.Data,
		a.Date,
		a.Display,
		a.Entities,
		a.IdMemberCreator,
		a.List,
		a.Member,
		a.MemberCreator,
		a.Organization,
		a.Text,
		a.Type,
	}

	as := "/actions/" + ID
	testExpects := []string{
		"/actions",
		as,
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
		as + "/card/iShort",
		as + "/card/manualCoverAttachment",
		as + "/card/name",
		as + "/card/pos",
		as + "/card/shortLink",
		as + "/card/shortUrl",
		as + "/card/subscribed",
		as + "/card/url",
		as + "/comments",
		as + "/data",
		as + "/date",
		as + "/display",
		as + "/entities",
		as + "/idMemberCreator",
		as + "/list",
		as + "/member",
		as + "/memberCreator",
		as + "/organization",
		as + "/text",
		as + "/type",
	}

	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}

}
