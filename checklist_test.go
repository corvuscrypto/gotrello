package trello

import "testing"

func testChecklists(t *testing.T) {

	testModels := []model{
		Checklists.ID(ID),
		Checklists.ID(ID).IdBoard,
		Checklists.ID(ID).IdCard,
		Checklists.ID(ID).Name,
		Checklists.ID(ID).Pos,
		Checklists.ID(ID).Board,
		Checklists.ID(ID).Board.Closed,
		Checklists.ID(ID).Board.DateLastActivity,
		Checklists.ID(ID).Board.DateLastView,
		Checklists.ID(ID).Board.Desc,
		Checklists.ID(ID).Board.DescData,
		Checklists.ID(ID).Board.IdOrganization,
		Checklists.ID(ID).Board.Invitations,
		Checklists.ID(ID).Board.Invited,
		Checklists.ID(ID).Board.LabelNames,
		Checklists.ID(ID).Board.Memberships,
		Checklists.ID(ID).Board.Name,
		Checklists.ID(ID).Board.Pinned,
		Checklists.ID(ID).Board.PowerUps,
		Checklists.ID(ID).Board.Prefs,
		Checklists.ID(ID).Board.ShortLink,
		Checklists.ID(ID).Board.ShortUrl,
		Checklists.ID(ID).Board.Starred,
		Checklists.ID(ID).Board.Subscribed,
		Checklists.ID(ID).Board.Url,
		Checklists.ID(ID).Cards,
		Checklists.ID(ID).Cards.All,
		Checklists.ID(ID).Cards.Closed,
		Checklists.ID(ID).Cards.None,
		Checklists.ID(ID).Cards.Open,
		Checklists.ID(ID).CheckItems,
		Checklists.ID(ID).CheckItems.ID(ID),
		Checklists,
	}

	testExpects := []string{
		"/checklists/" + ID,
		"/checklists/" + ID + "/idBoard",
		"/checklists/" + ID + "/idCard",
		"/checklists/" + ID + "/name",
		"/checklists/" + ID + "/pos",
		"/checklists/" + ID + "/board",
		"/checklists/" + ID + "/board/closed",
		"/checklists/" + ID + "/board/dateLastActivity",
		"/checklists/" + ID + "/board/dateLastView",
		"/checklists/" + ID + "/board/desc",
		"/checklists/" + ID + "/board/descData",
		"/checklists/" + ID + "/board/idOrganization",
		"/checklists/" + ID + "/board/invitations",
		"/checklists/" + ID + "/board/invited",
		"/checklists/" + ID + "/board/labelNames",
		"/checklists/" + ID + "/board/memberships",
		"/checklists/" + ID + "/board/name",
		"/checklists/" + ID + "/board/pinned",
		"/checklists/" + ID + "/board/powerUps",
		"/checklists/" + ID + "/board/prefs",
		"/checklists/" + ID + "/board/shortLink",
		"/checklists/" + ID + "/board/shortUrl",
		"/checklists/" + ID + "/board/starred",
		"/checklists/" + ID + "/board/subscribed",
		"/checklists/" + ID + "/board/url",
		"/checklists/" + ID + "/cards",
		"/checklists/" + ID + "/cards/all",
		"/checklists/" + ID + "/cards/closed",
		"/checklists/" + ID + "/cards/none",
		"/checklists/" + ID + "/cards/open",
		"/checklists/" + ID + "/checkItems",
		"/checklists/" + ID + "/checkItems/" + ID,
		"/checklists",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
