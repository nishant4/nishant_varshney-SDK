package lotr

import "lotr/endPoints"

type Config struct {
	Route     string // Main http address for the call
	EndPoints map[string]string
}

// override an endpoint
func (this *Config) SetEndpointUrl(endp string, url string) {
	this.EndPoints[endp] = url
}

func (this *Config) GetEndPoint(endp string) string {
	return this.EndPoints[endp]
}

// populates the endpoint map with defaults, can be overridden
func (this *Config) BuildDefaultEndPoints() {
	this.EndPoints = map[string]string{
		endPoints.LIST_BOOKS:        "/book",
		endPoints.GET_BOOK:          "/book/%{id}s",
		endPoints.GET_BOOK_CHAPTERS: "/book/%{id}s/chapter",
		endPoints.LIST_CHARACTERS:   "/character",
	}
}

func NewConfig(apiRoute string) *Config {
	conf := new(Config)
	conf.Route = apiRoute
	conf.BuildDefaultEndPoints()
	return conf
}
