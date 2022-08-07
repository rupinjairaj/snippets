package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AuthController interface {
	Login(response http.ResponseWriter, request *http.Request)
	Callback(response http.ResponseWriter, request *http.Request)
}

type AuthCtrl struct {
}

func NewAuthController() AuthController {
	return &AuthCtrl{}
}

func (ac *AuthCtrl) Login(response http.ResponseWriter, request *http.Request) {
	githubClientID := getEnvVariable("GITHUB_CLIENT_ID")
	githubCallbackUrl := getEnvVariable("GITHUB_OAUTH_CALLBACK")
	githubOauthSuccessRedirectUrl := getEnvVariable("GITHUB_OAUTH_SUCCESS_REDIRECT")
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s?path=%s&scope=user:email",
		githubClientID,
		githubCallbackUrl,
		githubOauthSuccessRedirectUrl,
	)
	http.Redirect(response, request, redirectURL, 301)
}

func (ac *AuthCtrl) Callback(response http.ResponseWriter, request *http.Request) {
	code := request.URL.Query().Get("code")
	redirectURL := request.URL.Query().Get("path")

	githubAccessToken := getGithubAccessToken(code)

	githubData := getGithubData(githubAccessToken)

	response.Header().Set("Content-type", "application/json")
	var prettyJSON bytes.Buffer
	// json.indent is a library utility function to prettify JSON indentation
	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error")
	}
	log.Println(githubData)
	cookie := http.Cookie{
		Name:     "snippet",
		Value:    githubAccessToken,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(response, &cookie)
	http.Redirect(response, request, redirectURL, 301)
}

func getGithubAccessToken(code string) string {

	clientID := getEnvVariable("GITHUB_CLIENT_ID")
	clientSecret := getEnvVariable("GITHUB_CLIENT_SECRET")

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, reqerr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if reqerr != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Response body converted to stringified JSON
	respbody, _ := ioutil.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken
}

func getEnvVariable(key string) string {
	clientSecret, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("'%s' secret not set in env", key)
	}

	return clientSecret
}

func getGithubData(accessToken string) string {
	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	return string(respbody)
}
