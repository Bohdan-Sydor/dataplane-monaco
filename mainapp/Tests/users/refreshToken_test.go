package usertests

import (
	"dataplane/Tests/testutils"
	"dataplane/auth"
	"log"
	"net/http"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

/*
go test -p 1 -v -count=1 -run TestRefreshToken dataplane/Tests/users
* Login
* Refresh token
*/
func TestRefreshToken(t *testing.T) {

	graphQLUrl := "http://localhost:9000/public/graphql"

	testUser := testutils.AdminUser
	testPassword := testutils.AdminPassword

	//--------- Login ------------

	loginUser := `{
		loginUser(
		  username: "` + testUser + `",
		  password: "` + testPassword + `",
		) {
		  access_token
		  refresh_token
		}
	  }`

	loginUserResponse, httpLoginResponse := testutils.GraphQLRequestPublic(loginUser, "{}", graphQLUrl, t)

	log.Println(string(loginUserResponse))

	if strings.Contains(string(loginUserResponse), `"errors":`) {
		t.Errorf("Error in graphql response")
	}

	assert.Equalf(t, http.StatusOK, httpLoginResponse.StatusCode, "Login user 200 status code")

	// -------------- Me ---------------------
	reqQuery := `{}`

	refreshToken := jsoniter.Get(loginUserResponse, "data", "loginUser", "refresh_token").ToString()

	url := "http://localhost:9000/refreshtoken"
	exchangeUserResponse, httpExchangeResponse := testutils.RestRequestPrivate(reqQuery, refreshToken, "POST", url, t)

	// log.Println(string(exchangeUserResponse))
	// log.Println(httpExchangeResponse)
	assert.Equalf(t, http.StatusOK, httpExchangeResponse.StatusCode, "Exchange token 200 status code")
	accessTokenExchange := jsoniter.Get(exchangeUserResponse, "access_token").ToString()

	// log.Println("Exchanged token: ", accessTokenExchange)
	validatetokenExchange, _ := auth.ValidateAccessToken(accessTokenExchange)
	assert.Equalf(t, true, validatetokenExchange, "Exchange access token validation")

}
