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

	client.httpClient = &http.Client{}

	return client, err
}

func (c Client) appendCredentials(url string) string {
	u := url + "?token=" + c.token + "&key=" + c.key
	return u
}

func (c Client) insertPostData(r *http.Request, values map[string]string) {
	r.PostForm.Add("token", c.token)
	r.PostForm.Add("key", c.key)
	for k, v := range values {
		r.PostForm.Add(k, v)
	}
}

//Get is the primary method for accessing Trello's RESTful GET APIs
func (c Client) Get(m model) (interface{}, error) {
	var err error
	var result interface{}
	var req *http.Request
	req, err = http.NewRequest("GET", c.appendCredentials(m.getURL()), nil)
	c.httpClient.Do(req)
	return result, err
}

//Post is the primary method for accessing Trello's RESTful POST APIs
func (c Client) Post(m model,
	values map[string]string) (interface{}, error) {
	var err error
	var result interface{}
	var req *http.Request
	req, err = http.NewRequest("POST", m.getURL(), nil)
	c.httpClient.Do(req)
	return result, err
}

//Put is the primary method for accessing Trello's RESTful PUT APIs
func (c Client) Put(m model, values map[string]string) (interface{}, error) {
	var err error
	var result interface{}
	var req *http.Request
	req, err = http.NewRequest("PUT", m.getURL(), nil)
	c.httpClient.Do(req)
	return result, err
}

//Delete is the primary method for accessing Trello's RESTful DELETE APIs
func (c Client) Delete(m model) (interface{}, error) {
	var err error
	var result interface{}
	var req *http.Request
	req, err = http.NewRequest("DELETE", c.appendCredentials(m.getURL()), nil)
	c.httpClient.Do(req)
	return result, err
}

//GetURL is the public method that converts a trello API model to a url string
func GetURL(m model) string {
	return m.getURL()
}
