package trello

import "testing"

func testLabels(t *testing.T) {

	testModels := []model{
		Labels.ID(ID),
		Labels.ID(ID).Board,
		Labels.ID(ID).Board.Closed,
		Labels.ID(ID).Board.DateLastActivity,
		Labels.ID(ID).Board.DateLastView,
		Labels.ID(ID).Board.Desc,
		Labels.ID(ID).Board.DescData,
		Labels.ID(ID).Board.IdOrganization,
		Labels.ID(ID).Board.Invitations,
		Labels.ID(ID).Board.Invited,
		Labels.ID(ID).Board.LabelNames,
		Labels.ID(ID).Board.Memberships,
		Labels.ID(ID).Board.Name,
		Labels.ID(ID).Board.Pinned,
		Labels.ID(ID).Board.PowerUps,
		Labels.ID(ID).Board.Prefs,
		Labels.ID(ID).Board.ShortLink,
		Labels.ID(ID).Board.ShortUrl,
		Labels.ID(ID).Board.Starred,
		Labels.ID(ID).Board.Subscribed,
		Labels.ID(ID).Board.Url,
		Labels.ID(ID).Color,
		Labels.ID(ID).Name,
		Labels,
	}

	testExpects := []string{
		"/labels/" + ID,
		"/labels/" + ID + "/board",
		"/labels/" + ID + "/board/closed",
		"/labels/" + ID + "/board/dateLastActivity",
		"/labels/" + ID + "/board/dateLastView",
		"/labels/" + ID + "/board/desc",
		"/labels/" + ID + "/board/descData",
		"/labels/" + ID + "/board/idOrganization",
		"/labels/" + ID + "/board/invitations",
		"/labels/" + ID + "/board/invited",
		"/labels/" + ID + "/board/labelNames",
		"/labels/" + ID + "/board/memberships",
		"/labels/" + ID + "/board/name",
		"/labels/" + ID + "/board/pinned",
		"/labels/" + ID + "/board/powerUps",
		"/labels/" + ID + "/board/prefs",
		"/labels/" + ID + "/board/shortLink",
		"/labels/" + ID + "/board/shortUrl",
		"/labels/" + ID + "/board/starred",
		"/labels/" + ID + "/board/subscribed",
		"/labels/" + ID + "/board/url",
		"/labels/" + ID + "/color",
		"/labels/" + ID + "/name",
		"/labels",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
