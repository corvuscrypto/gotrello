package trello

import "testing"

func testLists(t *testing.T) {

	testModels := []model{
		Lists.ID(ID),
		Lists.ID(ID).Closed,
		Lists.ID(ID).IdBoard,
		Lists.ID(ID).Name,
		Lists.ID(ID).Pos,
		Lists.ID(ID).Subscribed,
		Lists.ID(ID).Actions,
		Lists.ID(ID).Board,
		Lists.ID(ID).Board.Closed,
		Lists.ID(ID).Board.DateLastActivity,
		Lists.ID(ID).Board.DateLastView,
		Lists.ID(ID).Board.Desc,
		Lists.ID(ID).Board.DescData,
		Lists.ID(ID).Board.IdOrganization,
		Lists.ID(ID).Board.Invitations,
		Lists.ID(ID).Board.Invited,
		Lists.ID(ID).Board.LabelNames,
		Lists.ID(ID).Board.Memberships,
		Lists.ID(ID).Board.Name,
		Lists.ID(ID).Board.Pinned,
		Lists.ID(ID).Board.PowerUps,
		Lists.ID(ID).Board.Prefs,
		Lists.ID(ID).Board.ShortLink,
		Lists.ID(ID).Board.ShortUrl,
		Lists.ID(ID).Board.Starred,
		Lists.ID(ID).Board.Subscribed,
		Lists.ID(ID).Board.Url,
		Lists.ID(ID).Cards,
		Lists.ID(ID).Cards.All,
		Lists.ID(ID).Cards.Closed,
		Lists.ID(ID).Cards.None,
		Lists.ID(ID).Cards.Open,
		Lists,
		Lists.ID(ID).ArchiveAllCards,
		Lists.ID(ID).MoveAllCards,
	}

	testExpects := []string{
		"/lists/" + ID,
		"/lists/" + ID + "/closed",
		"/lists/" + ID + "/idBoard",
		"/lists/" + ID + "/name",
		"/lists/" + ID + "/pos",
		"/lists/" + ID + "/subscribed",
		"/lists/" + ID + "/actions",
		"/lists/" + ID + "/board",
		"/lists/" + ID + "/board/closed",
		"/lists/" + ID + "/board/dateLastActivity",
		"/lists/" + ID + "/board/dateLastView",
		"/lists/" + ID + "/board/desc",
		"/lists/" + ID + "/board/descData",
		"/lists/" + ID + "/board/idOrganization",
		"/lists/" + ID + "/board/invitations",
		"/lists/" + ID + "/board/invited",
		"/lists/" + ID + "/board/labelNames",
		"/lists/" + ID + "/board/memberships",
		"/lists/" + ID + "/board/name",
		"/lists/" + ID + "/board/pinned",
		"/lists/" + ID + "/board/powerUps",
		"/lists/" + ID + "/board/prefs",
		"/lists/" + ID + "/board/shortLink",
		"/lists/" + ID + "/board/shortUrl",
		"/lists/" + ID + "/board/starred",
		"/lists/" + ID + "/board/subscribed",
		"/lists/" + ID + "/board/url",
		"/lists/" + ID + "/cards",
		"/lists/" + ID + "/cards/all",
		"/lists/" + ID + "/cards/closed",
		"/lists/" + ID + "/cards/none",
		"/lists/" + ID + "/cards/open",
		"/lists",
		"/lists/" + ID + "/archiveAllCards",
		"/lists/" + ID + "/moveAllCards",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
