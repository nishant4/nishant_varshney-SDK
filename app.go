package lotr

import (
	"fmt"
	"lotr/consts"
	"lotr/network"
	"strings"
)

type App struct {
	conf        *Config
	accessToken string
}

func (this *App) SetConfig(conf *Config) {
	this.conf = conf
}

func (this *App) buildEndPointWithIdentifiers(url string, params map[string]string) string {
	for key, val := range params {
		url = strings.Replace(url, "%{"+key+"}s", fmt.Sprintf("%s", val), -1)
	}
	return url
}

func (this *App) ApiOptions(endPoint string, options *GetOptions, identifiers map[string]string) ([]byte, error) {
	var params map[string]string

	if options != nil {
		if err := options.Validate(); err != nil {
			return nil, err
		}
		params = options.GetParams()
	}

	endPoint = this.conf.GetEndPoint(endPoint)
	// replace id etc. with values from the map
	if identifiers != nil && len(identifiers) > 0 {
		endPoint = this.buildEndPointWithIdentifiers(endPoint, identifiers)
	}

	return network.HttpBearerAuthEndPointWithParams(this.accessToken, this.conf.Route, endPoint, params)
}

func NewWithConfig(accessToken string, conf *Config) *App {
	ans := new(App)
	ans.accessToken = accessToken
	ans.SetConfig(conf)
	return ans
}

func New(accessToken string) *App {
	conf := NewConfig(consts.DEFAULT_API_ROUTE)
	return NewWithConfig(accessToken, conf)
}
