package trello

type webhook struct {
	url string
	Active,
	CallbackURL,
	Description,
	IdModel staticField
}

func createWebhook(m model) webhook {
	wURL := m.getURL() + "/webhooks"
	return webhook{
		url:         wURL,
		Active:      staticField(wURL + "/active"),
		CallbackURL: staticField(wURL + "/callbackURL"),
		Description: staticField(wURL + "/description"),
		IdModel:     staticField(wURL + "/idModel"),
	}
}

func (w webhook) ID(id string) webhook {
	wURL := w.url + "/" + id
	return webhook{
		url:         wURL,
		Active:      staticField(wURL + "/active"),
		CallbackURL: staticField(wURL + "/callbackURL"),
		Description: staticField(wURL + "/description"),
		IdModel:     staticField(wURL + "/idModel"),
	}
}

func (w webhook) getURL() string {
	return w.url
}

var Webhooks = webhook{
	url: "/webhooks",
}
