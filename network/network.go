package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func paramsPrintf(url string, params map[string]string) string {
	for key, val := range params {
		url = strings.Replace(url, "%{"+key+"}s", fmt.Sprintf("%s", val), -1)
	}
	return url
}

func buildUrl(address string, endpoint string) string {
	if len(endpoint) > 0 && endpoint[0] != '/' && address[len(address)-1:] != "/" {
		return address + "/" + endpoint
	}
	return address + endpoint
}

func buildUrlWithIdentifiers(address string, endpoint string, params map[string]string) string {
	return buildUrl(address, paramsPrintf(endpoint, params))
}

// func HttpCall(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return ioutil.ReadAll(resp.Body)
// }

// func HttpCallEndPoint(address string, endpoint string) ([]byte, error) {
// 	url := buildUrl(address, endpoint)
// 	return HttpCall(url)
// }

// func HttpCallWithIdentifiers(address string, endpoint string, params map[string]string) ([]byte, error) {
// 	url := buildUrlWithIdentifiers(address, endpoint, params)
// 	log.Println(url)
// 	return HttpCall(url)
// }

// func HttpBearerAuth(accessToken string, url string) ([]byte, error) {
// 	var bearer = "Bearer " + accessToken

// 	// Create a new request using http
// 	req, err := http.NewRequest("GET", url, nil)

// 	// add authorization header to the req
// 	req.Header.Add("Authorization", bearer)

// 	// Send req using http Client
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	return ioutil.ReadAll(resp.Body)
// }

// func HttpBearerAuthEndPoint(accessToken string, address string, endpoint string) ([]byte, error) {
// 	url := buildUrl(address, endpoint)
// 	log.Println(url)
// 	return HttpBearerAuth(accessToken, url)
// }

func HttpBearerAuthWithParams(accessToken string, url string, params map[string]string) ([]byte, error) {
	var bearer = "Bearer " + accessToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	if params != nil {
		q := req.URL.Query()
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	log.Println("url: ", req.URL.String())

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(bytes, &respMap)
	if err != nil {
		return nil, err
	}

	// check if it failed because unauthorized
	if val, ok := respMap["message"]; ok {
		if val == "Unauthorized." {
			return nil, fmt.Errorf("Unauthorized")
		}
	}

	return bytes, nil
}

func HttpBearerAuthEndPointWithParams(accessToken string, address string, endpoint string, params map[string]string) ([]byte, error) {
	url := buildUrl(address, endpoint)
	return HttpBearerAuthWithParams(accessToken, url, params)
}
