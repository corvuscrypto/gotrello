package trello

import "testing"

func testSessions(t *testing.T) {

 testModels := []model{
  Sessions.Socket,
Sessions.ID(ID),
Sessions.ID(ID).Status,
Sessions,

 }

 testExpects := []string{
  "/sessions/socket",
"/sessions/"+ID,
"/sessions/"+ID+"/status",
"/sessions",
 }
 for i := range testModels {
		m := testModels[i]
		e := testExpects[i]
		if s := GetURL(m); s != e {
			t.Errorf("test model# %d: Expected \"%s\", Got \"%s\"\n", i, e, s)
		}
	}
}
