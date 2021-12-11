// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package publicgraphql

import (
	"dataplane/database/models"
)

type AddAdminsInput struct {
	PlatformInput *PlatformInput `json:"PlatformInput"`
	AddUsersInput *AddUsersInput `json:"AddUsersInput"`
}

type AddUsersInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Password  string `json:"password"`
	Timezone  string `json:"timezone"`
}

type Admin struct {
	Platform *models.Platform `json:"Platform"`
	User     *models.Users    `json:"User"`
	Auth     *Authtoken       `json:"Auth"`
}

type Authtoken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PlatformInput struct {
	BusinessName string `json:"business_name"`
	Timezone     string `json:"timezone"`
	Complete     bool   `json:"complete"`
}
