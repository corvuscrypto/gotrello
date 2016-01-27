package trello

import "testing"

func testTypes(t *testing.T) {

 testModels := []model{
  Types.ID(ID),

 }

 testExpects := []string{
  "/types/"+ID,
 }
 for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
