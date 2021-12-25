// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package privategraphql

type AddEnvironmentInput struct {
	Name string `json:"name"`
}

type AddPreferencesInput struct {
	Preference string `json:"preference"`
	Value      string `json:"value"`
}

type AddUpdateMeInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Timezone  string `json:"timezone"`
}

type AddUsersInput struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	JobTitle      string `json:"job_title"`
	Password      string `json:"password"`
	Timezone      string `json:"timezone"`
	EnvironmentID string `json:"environmentID"`
}

type ChangePasswordInput struct {
	Password      string `json:"password"`
	EnvironmentID string `json:"environmentID"`
	UserID        string `json:"user_id"`
}

type Pipelines struct {
	Name string `json:"name"`
}

type Platform struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Timezone     string `json:"timezone"`
	Complete     bool   `json:"complete"`
}

type Preferences struct {
	Preference string `json:"preference"`
	Value      string `json:"value"`
}

type RenameEnvironment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Workers struct {
	Name string `json:"name"`
}

type UpdatePlatformInput struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Timezone     string `json:"timezone"`
	Complete     bool   `json:"complete"`
}
