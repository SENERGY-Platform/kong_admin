package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClientType interface {
	LoadRoutes() (*[]string, error, int)
}

type Client struct {
	BaseUrl string
}

func (c *Client) LoadRoutes() (*[]string, error, int) {
	req, err := http.NewRequest(http.MethodGet, c.BaseUrl+"/routes", nil)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return doWithResponse[*[]string](req)
}

func NewClient(url string) ClientType {
	return &Client{BaseUrl: url}
}

func doWithResponse[T any](req *http.Request) (result T, err error, code int) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		temp, _ := io.ReadAll(resp.Body) //read error response end ensure that resp.Body is read to EOF
		return result, fmt.Errorf("unexpected statuscode %v: %v", resp.StatusCode, string(temp)), resp.StatusCode
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		_, _ = io.ReadAll(resp.Body) //ensure resp.Body is read to EOF
		return result, err, http.StatusInternalServerError
	}
	return
}
