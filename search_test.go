package trello

import "testing"

func testSearchs(t *testing.T) {

	testModels := []model{
		Search,
		Search.Members,
	}

	testExpects := []string{
		"/search",
		"/search/members",
	}
	for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
