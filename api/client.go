package api

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
)

var cache = &modelsCache{}

// Client is the nine-to-fiver client, responsible for
// fetching and update items
type Client struct {
	Endpoint string
	Username string
	Password string
	CacheDir string
	NoCache  bool
}

type modelList interface {
	apiURL() string
	slug() string
	augment() error
	isEmpty() bool
}

type modelsCache struct {
	models []modelList
}

func (c *modelsCache) register(m modelList) {
	c.models = append(c.models, m)
}

func (c *modelsCache) fetchRemotely(client *Client) error {
	for _, m := range c.models {
		if !m.isEmpty() {
			// We aready have it
			continue
		}
		err := client.FetchList(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) getCacheDir() string {
	baseDir, _ := homedir.Expand(c.CacheDir)
	h := sha1.New()
	uniqueSha1 := fmt.Sprintf("%s?user=%s", c.Endpoint, c.Username)
	h.Write([]byte(uniqueSha1))
	sha1 := h.Sum(nil)
	clientDir := fmt.Sprintf("%x", sha1)
	os.Mkdir(baseDir, 0700)
	cacheDir := path.Join(baseDir, clientDir)
	os.Mkdir(cacheDir, 0700)
	return cacheDir
}

func (c *modelsCache) fetchLocally(client *Client) error {
	if client.NoCache {
		return nil
	}
	dir := client.getCacheDir()
	for _, m := range c.models {
		modelPath := path.Join(dir, fmt.Sprintf("%s.json", m.slug()))
		_, err := os.Stat(modelPath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		modelsData, _ := ioutil.ReadFile(modelPath)
		err = json.Unmarshal(modelsData, &m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *modelsCache) save(client *Client) error {
	if client.NoCache {
		return nil
	}
	dir := client.getCacheDir()
	for _, m := range c.models {
		modelPath := path.Join(dir, fmt.Sprintf("%s.json", m.slug()))
		modelsData, _ := json.Marshal(m)
		err := ioutil.WriteFile(modelPath, modelsData, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *modelsCache) fetch(client *Client) error {
	c.fetchLocally(client)
	c.fetchRemotely(client)
	for _, m := range c.models {
		err := m.augment()
		if err != nil {
			return err
		}
	}
	return nil
}

// FetchCache will fill the cache with content from
// localdisk or ninetofiver instance
func (c *Client) FetchCache() error {
	err := cache.fetch(c)
	if err != nil {
		return err
	}
	return cache.save(c)
}

// FetchList fetches a list of objects from the backend
func (c *Client) FetchList(m modelList) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/%s/?page_size=9999", c.Endpoint, m.apiURL()))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(m)
	if err != nil {
		return err
	}
	return nil
}

// PatchRequest will make a PATCH request for a given model
// and expect error code to be 200
func (c *Client) PatchRequest(url string, i interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)
	return c.Request("PATCH", url, 200, b)
}

// PostRequest will make a POST request for a given model
// and expect error code to be 201
func (c *Client) PostRequest(url string, i interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)
	return c.Request("POST", url, 201, b)
}

// DeleteRequest will make a DELETE request for a given model
// and expect error code to be 204
func (c *Client) DeleteRequest(url string) (*http.Response, error) {
	return c.Request("DELETE", url, 204, nil)
}

// GetRequest will make a GET request for a given model
// and expect error code to be 200
func (c *Client) GetRequest(url string) (*http.Response, error) {
	return c.Request("GET", url, 200, nil)
}

// Request makes a requests, sets the correct headers,
// and checks the return code of the response.
func (c *Client) Request(verb, url string, code int, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("12to8/0.0.1 (%s)", runtime.Version()))
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != code {
		var content []byte
		out := make([]byte, 100)
		_, err = resp.Body.Read(out)
		if err == nil {
			content = out
		}
		return resp, fmt.Errorf("Received %d, expecting %d status code while fetching %s\n%s", resp.StatusCode, code, url, string(content))
	}
	return resp, err
}
