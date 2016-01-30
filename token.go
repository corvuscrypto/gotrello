package trello

type token struct {
	url string
	DateCreated,
	DateExpires,
	IdMember,
	Identifier,
	Permissions staticField
	Member   baseMember
	Webhooks webhook
}

func createToken(m model) token {
	tURL := m.getURL() + "/token"
	return token{
		url:         tURL,
		DateCreated: staticField(tURL + "/dateCreated"),
		DateExpires: staticField(tURL + "/dateExpires"),
		IdMember:    staticField(tURL + "/idMember"),
		Identifier:  staticField(tURL + "/identifier"),
		Permissions: staticField(tURL + "/permissions"),
		Member:      createBaseMember(m),
		Webhooks:    createWebhook(m),
	}
}

func (t token) ID(id string) token {
	tURL := t.url + "/" + id
	return token{
		url:         tURL,
		DateCreated: staticField(tURL + "/dateCreated"),
		DateExpires: staticField(tURL + "/dateExpires"),
		IdMember:    staticField(tURL + "/idMember"),
		Identifier:  staticField(tURL + "/identifier"),
		Permissions: staticField(tURL + "/permissions"),
		Member:      createBaseMember(t),
		Webhooks:    createWebhook(t),
	}
}

func (t token) getURL() string {
	return t.url
}

var Tokens = token{
	url: "/tokens",
}
