package trello

import "testing"

func testWebhooks(t *testing.T) {

	testModels := []model{
		Webhooks.ID(ID),
		Webhooks.ID(ID).Active,
		Webhooks.ID(ID).CallbackURL,
		Webhooks.ID(ID).Description,
		Webhooks.ID(ID).IdModel,
		Webhooks,
	}

	testExpects := []string{
		"/webhooks/" + ID,
		"/webhooks/" + ID + "/active",
		"/webhooks/" + ID + "/callbackURL",
		"/webhooks/" + ID + "/description",
		"/webhooks/" + ID + "/idModel",
		"/webhooks",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
