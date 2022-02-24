// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package privategraphql

import (
	"dataplane/mainapp/database/models"
	"time"
)

type AccessGroupsInput struct {
	AccessGroupID string `json:"AccessGroupID"`
	Name          string `json:"Name"`
	Description   string `json:"Description"`
	Active        bool   `json:"Active"`
	EnvironmentID string `json:"EnvironmentID"`
}

type AddEnvironmentInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type AddPreferencesInput struct {
	Preference string `json:"preference"`
	Value      string `json:"value"`
}

type AddSecretsInput struct {
	Secret        string  `json:"Secret"`
	Description   *string `json:"Description"`
	Value         string  `json:"Value"`
	EnvironmentID string  `json:"EnvironmentId"`
	Active        bool    `json:"Active"`
}

type AddUpdateMeInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Timezone  string `json:"timezone"`
}

type AddUsersInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Password  string `json:"password"`
	Timezone  string `json:"timezone"`
}

type ChangePasswordInput struct {
	Password string `json:"password"`
	UserID   string `json:"user_id"`
}

type DataInput struct {
	Language string `json:"language"`
}

type PipelineEdgesInput struct {
	EdgeID string                  `json:"edgeID"`
	From   string                  `json:"from"`
	To     string                  `json:"to"`
	Meta   *PipelineEdgesMetaInput `json:"meta"`
	Active bool                    `json:"active"`
}

type PipelineEdgesMetaInput struct {
	SourceHandle  string `json:"sourceHandle"`
	TargetHandle  string `json:"targetHandle"`
	EdgeType      string `json:"edgeType"`
	ArrowHeadType string `json:"arrowHeadType"`
}

type PipelineFlow struct {
	Edges []*models.PipelineEdges `json:"edges"`
	Nodes []*models.PipelineNodes `json:"nodes"`
}

type PipelineFlowInput struct {
	NodesInput []*PipelineNodesInput `json:"nodesInput"`
	EdgesInput []*PipelineEdgesInput `json:"edgesInput"`
}

type PipelineNodesInput struct {
	NodeID       string                  `json:"nodeID"`
	Name         string                  `json:"name"`
	NodeType     string                  `json:"nodeType"`
	NodeTypeDesc string                  `json:"nodeTypeDesc"`
	Description  string                  `json:"description"`
	Meta         *PipelineNodesMetaInput `json:"meta"`
	Active       bool                    `json:"active"`
}

type PipelineNodesMetaInput struct {
	Position *PositionInput `json:"position"`
	Data     *DataInput     `json:"data"`
}

type PipelinePermissionsOutput struct {
	Access        string `json:"Access"`
	Subject       string `json:"Subject"`
	SubjectID     string `json:"SubjectID"`
	PipelineName  string `json:"PipelineName"`
	ResourceID    string `json:"ResourceID"`
	EnvironmentID string `json:"EnvironmentID"`
	Active        bool   `json:"Active"`
	Level         string `json:"Level"`
	Label         string `json:"Label"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	Email         string `json:"Email"`
	JobTitle      string `json:"JobTitle"`
}

type Platform struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Timezone     string `json:"timezone"`
	Complete     bool   `json:"complete"`
}

type PositionInput struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Preferences struct {
	Preference string `json:"preference"`
	Value      string `json:"value"`
}

type UpdateEnvironment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateSecretsInput struct {
	Secret        string  `json:"Secret"`
	Description   *string `json:"Description"`
	EnvironmentID string  `json:"EnvironmentId"`
	Active        bool    `json:"Active"`
}

type UpdateUsersInput struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	JobTitle  string `json:"job_title"`
	Timezone  string `json:"timezone"`
}

type WorkerGroup struct {
	WorkerGroup string    `json:"WorkerGroup"`
	Status      string    `json:"Status"`
	T           time.Time `json:"T"`
	Interval    int       `json:"Interval"`
	Env         string    `json:"Env"`
	Lb          string    `json:"LB"`
	WorkerType  string    `json:"WorkerType"`
}

type Workers struct {
	WorkerGroup string    `json:"WorkerGroup"`
	WorkerID    string    `json:"WorkerID"`
	Status      string    `json:"Status"`
	T           time.Time `json:"T"`
	Interval    int       `json:"Interval"`
	CPUPerc     float64   `json:"CPUPerc"`
	Load        float64   `json:"Load"`
	MemoryPerc  float64   `json:"MemoryPerc"`
	MemoryUsed  float64   `json:"MemoryUsed"`
	Env         string    `json:"Env"`
	Lb          string    `json:"LB"`
	WorkerType  string    `json:"WorkerType"`
}

type UpdatePlatformInput struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Timezone     string `json:"timezone"`
	Complete     bool   `json:"complete"`
}
