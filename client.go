package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OauthToken struct {
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	ExpiresIn        int64  `json:"expires_in"`
	AccessToken      string `json:"access_token"`
	createdTimestamp int64
}

type ApiClient struct {
	clientId     string
	clientSecret string
	baseApiUrl   string
	client       *http.Client
	oauthToken   OauthToken
}

func (api *ApiClient) getUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", api.baseApiUrl, endpoint)
}

func (api *ApiClient) setOauthToken() {
	if api.oauthToken.AccessToken == "" {
		requestBody := new(bytes.Buffer)
		requestBodyData := struct {
			GrantType    string `json:"grant_type"`
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}{
			"client_credentials",
			api.clientId,
			api.clientSecret,
		}
		json.NewEncoder(requestBody).Encode(requestBodyData)
		res, err := api.client.Post(api.getUrl("/oauth/token"), "application/json; charset=utf-8", requestBody)
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(res.Body).Decode(&api.oauthToken)
		api.oauthToken.createdTimestamp = time.Now().UTC().Unix()
	} else if time.Now().UTC().Unix()-api.oauthToken.createdTimestamp >= api.oauthToken.ExpiresIn {
		// Refresh oauth token
		api.oauthToken = OauthToken{}
		api.setOauthToken()
	}
}

func (api *ApiClient) getAuthorizationHeader() string {
	api.setOauthToken()
	return fmt.Sprintf("%s %s", strings.Title(api.oauthToken.TokenType), api.oauthToken.AccessToken)
}

func (api *ApiClient) getReactions(tagName string, gfyCount int) GfycatReactions {
	endpoint := "/reactions/populated"
	queryStringValue := url.Values{}
	queryStringValue.Set("tagName", tagName)
	queryStringValue.Set("gfyCount", strconv.Itoa(gfyCount))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", api.getUrl(endpoint), queryStringValue.Encode()), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", api.getAuthorizationHeader())
	res, err := api.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	gfycatReactions := GfycatReactions{}
	json.NewDecoder(res.Body).Decode(&gfycatReactions)
	return gfycatReactions

}

// For debugging purpose (just to view raw response body)
func printResponseBody(res *http.Response) {
	defer res.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseBody := string(responseBodyBytes)
	fmt.Println(responseBody)
}
