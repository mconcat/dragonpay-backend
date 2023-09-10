package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConf *oauth2.Config

var oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthInit() {
	var clientID = os.Getenv("GOOGLE_CLIENT_ID")

	var clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	fmt.Println("clientID: ", clientID)
	fmt.Println("clientSecret: ", clientSecret)

	oauthConf = &oauth2.Config{
		ClientID:     clientID, // + ".apps.googleusercontent.com",
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/login/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func init() {
	oauthInit()
}

func getToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawStdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return oauthConf.AuthCodeURL(state)
}

func login(c *gin.Context) {
	token := getToken()
	url := getLoginURL(token)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Socal Login",
		"url":   url,
	})
}

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

var tokens = make(map[string]int64) // token -> id

func loginCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	fmt.Println(string(contents))
	user := User{}
	json.Unmarshal(contents, &user)

	tokens[token.AccessToken] = user.Id

	c.JSON(200, gin.H{"Message": "OK"})
}
