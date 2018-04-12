package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseAPI = "gitlab.com/api/v4/projects"
	issue   = "/issues"
)

// Client is a struct representing client for accessing gitlab API
type Client struct {
	baseURL      string
	privateToken string
}

type gitlabIssue struct {
	Title  string `json:"title"`
	WebURL string `json:"web_url"`
}

// New is a function for instantiating client with given private token
func New(projectID, token string) *Client {
	return &Client{
		baseURL:      baseAPI + "/" + projectID,
		privateToken: token,
	}
}

// GetIssue is a function for retrieving link to an issue with given IID
func (c *Client) GetIssue(iid string) error {

	req, _ := http.NewRequest("GET", c.baseURL+issue, nil)
	req.Header.Add("private-token", c.privateToken)

	q := req.URL.Query()
	q.Add("iids[]", iid)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var respIssue gitlabIssue
	json.NewDecoder(resp.Body).Decode(&respIssue)

	fmt.Println(respIssue)

	return nil
}
