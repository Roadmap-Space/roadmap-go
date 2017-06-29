package roadmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var apiClient *Client

// HTTPTimeout is the maximum number of second to wait for the API responses.
var HTTPTimeout = 25 * time.Second

// Client is used to perform API http request
type Client struct {
	email string
	token string
	c     *http.Client

	BasePath string

	Roadmaps *Roadmaps
	Feedback *Feedback
	Ideas    *Ideas
	Stories  *Stories
	Items    *Items
}

// New creates a new Client
func New(email, token string) *Client {
	basePath := "https://app.roadmap.space/v1"
	if os.Getenv("RM_DEBUG") == "1" {
		basePath = "http://localhost:8070/v1"
	}

	httpClient := &http.Client{Timeout: HTTPTimeout}
	apiClient = &Client{
		email:    email,
		token:    token,
		c:        httpClient,
		BasePath: basePath,
		Roadmaps: &Roadmaps{EndpointURL: "/roadmaps"},
		Feedback: &Feedback{EndpointURL: "/feedback"},
		Ideas:    &Ideas{EndpointURL: "/ideas"},
		Stories:  &Stories{EndpointURL: "/stories"},
		Items:    &Items{EndpointURL: "/items"},
	}

	return apiClient
}

func (api *Client) get(path string, result interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", api.BasePath, path), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(api.email, api.token)

	resp, err := api.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error requesting %s returned %s", path, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("cannot read body", err)
		}
		fmt.Println(string(b))
	}
	return err
}

func (api *Client) post(path string, data interface{}, result interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", api.BasePath, path), bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.SetBasicAuth(api.email, api.token)

	resp, err := api.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error requesting %s returned %s", path, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil
}

func (api *Client) put(path string, data interface{}, result interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s", api.BasePath, path), bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.SetBasicAuth(api.email, api.token)

	resp, err := api.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error requesting %s returned %s", path, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil
}

func (api *Client) delete(path string, result interface{}) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", api.BasePath, path), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(api.email, api.token)

	resp, err := api.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error requesting %s returned %s", path, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil
}
