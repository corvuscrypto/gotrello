package trello

import "net/http"

//Client is the object that allows API calls to be made
type Client struct {
	token      string
	key        string
	httpClient *http.Client
}

//NewClient takes an app key and secret key (Available from Trello)
//and attempts to authorize and return a usable client object
func NewClient(token string, key string) (*Client, error) {
	var err error
	var client = &Client{}

	return client, err
}

func (c Client) Get() {

}
