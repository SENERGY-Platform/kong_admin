package kong_client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Route struct {
	Paths []string `json:"paths"`
}

type RouteKongResult struct {
	Data []Route `json:"data"`
}

type Routes []string

func (client *KongClient) LoadRoutes() (routesList Routes, err error) {
	u, err := url.Parse(client.KONG_ADMIN_URL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, "routes")
	routeUrl := u.String()
	resp, err := http.Get(routeUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var routesFromKong RouteKongResult
	json.Unmarshal(body, &routesFromKong)

	for _, routes := range routesFromKong.Data {
		routesList = append(routesList, routes.Paths...)
	}
	return
}
