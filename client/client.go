package client

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseAPI      = "http://gitlab.com/api/v4/projects"
	issue        = "/issues"
	mergeRequest = "/merge_requests"
)

// Client is a struct representing client for accessing gitlab API
type Client struct {
	baseURL      string
	privateToken string
}

// GitlabItem is a struct representing item on gitlab eg. Issue, Merge Request
type GitlabItem struct {
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
func (c *Client) GetIssue(iid string) (GitlabItem, error) {

	req, _ := http.NewRequest("GET", c.baseURL+issue, nil)
	req.Header.Add("private-token", c.privateToken)

	q := req.URL.Query()
	q.Add("iids[]", iid)
	req.URL.RawQuery = html.UnescapeString(q.Encode())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return GitlabItem{}, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var f []interface{}
	json.Unmarshal(bodyBytes, &f)

	respMap := f[0].(map[string]interface{})

	respIssue := GitlabItem{
		Title:  respMap["title"].(string),
		WebURL: respMap["web_url"].(string),
	}

	return respIssue, nil
}

// GetIssue is a function for retrieving link to an issue with given IID
func (c *Client) GetMergeRequest(iid string) (GitlabItem, error) {

	req, _ := http.NewRequest("GET", c.baseURL+mergeRequest, nil)
	req.Header.Add("private-token", c.privateToken)

	q := req.URL.Query()
	q.Add("iids[]", iid)
	req.URL.RawQuery = html.UnescapeString(q.Encode())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return GitlabItem{}, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var f []interface{}
	json.Unmarshal(bodyBytes, &f)

	respMap := f[0].(map[string]interface{})

	respIssue := GitlabItem{
		Title:  respMap["title"].(string),
		WebURL: respMap["web_url"].(string),
	}

	return respIssue, nil
}
