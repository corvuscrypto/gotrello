package trello

type session struct {
	url string
	Socket,
	Status staticField
}

func (s session) ID(id string) session {
	sURL := s.getURL() + "/" + id
	return session{
		url:    sURL,
		Socket: staticField(sURL + "/socket"),
		Status: staticField(sURL + "/status"),
	}
}

func (s session) getURL() string {
	return s.url
}

var Sessions = session{
	url:    "/sessions",
	Socket: staticField("/sessions/socket"),
}
