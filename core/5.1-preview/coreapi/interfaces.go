package coreapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CreateTeam(ctx context.Context, organization string, body core.WebAPITeam, projectID string) (result core.WebAPITeam, err error)
	DeleteTeam(ctx context.Context, organization string, projectID string, teamID string) (result autorest.Response, err error)
	GetAllTeams(ctx context.Context, organization string, mine *bool, top *int32, skip *int32) (result core.ListWebAPITeam, err error)
	GetProcessByID(ctx context.Context, organization string, processID uuid.UUID) (result core.Process, err error)
	GetProcesses(ctx context.Context, organization string) (result core.ListProcess, err error)
	GetProject(ctx context.Context, organization string, projectID string, includeCapabilities *bool, includeHistory *bool) (result core.TeamProject, err error)
	GetProjectProperties(ctx context.Context, organization string, projectID uuid.UUID, keys string) (result core.ListProjectProperty, err error)
	GetProjects(ctx context.Context, organization string, stateFilter core.ProjectState, top *int32, skip *int32, continuationToken string, getDefaultTeamImageURL *bool) (result core.ListTeamProjectReference, err error)
	GetTeam(ctx context.Context, organization string, projectID string, teamID string) (result core.WebAPITeam, err error)
	GetTeamMembersWithExtendedProperties(ctx context.Context, organization string, projectID string, teamID string, top *int32, skip *int32) (result core.ListTeamMember, err error)
	GetTeams(ctx context.Context, organization string, projectID string, mine *bool, top *int32, skip *int32) (result core.ListWebAPITeam, err error)
	QueueCreateProject(ctx context.Context, organization string, body core.TeamProject) (result core.OperationReference, err error)
	QueueDeleteProject(ctx context.Context, organization string, projectID uuid.UUID) (result core.OperationReference, err error)
	RemoveProjectAvatar(ctx context.Context, organization string, projectID string) (result autorest.Response, err error)
	SetProjectAvatar(ctx context.Context, organization string, body core.ProjectAvatar, projectID string) (result autorest.Response, err error)
	SetProjectProperties(ctx context.Context, organization string, projectID uuid.UUID, body interface{}) (result autorest.Response, err error)
	UpdateProject(ctx context.Context, organization string, body core.TeamProject, projectID uuid.UUID) (result core.OperationReference, err error)
	UpdateTeam(ctx context.Context, organization string, body core.WebAPITeam, projectID string, teamID string) (result core.WebAPITeam, err error)
}

var _ BaseClientAPI = (*core.BaseClient)(nil)
