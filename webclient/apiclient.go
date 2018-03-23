package webclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

const (
	teamPathTemplate = "/api/teams/en/%d.json"
)

type teamResponseBuilder func(body io.ReadCloser) (*TeamResponse, error)

type apiClient struct {
	api     *url.URL
	builder teamResponseBuilder
	client  HTTPClientInterface
}

/*
APIClientInterface ...
*/
type APIClientInterface interface {
	GetTeam(id int) (*Team, error)
}

/*
NewAPIClient ...
*/
func NewAPIClient() APIClientInterface {
	api, _ := url.Parse("https://vintagemonster.onefootball.com")

	return &apiClient{
		api:     api,
		builder: BuildTeamResponse,
		client:  &httpClient{},
	}
}

func (ac *apiClient) GetTeam(id int) (*Team, error) {
	api := *ac.api
	api.Path = fmt.Sprintf(teamPathTemplate, id)

	response, err := ac.client.Get(&api)
	if nil != err {
		return nil, err
	}
	defer response.Body.Close()
	teamResponse, err := ac.builder(response.Body)
	if nil != err {
		return nil, err
	}

	return teamResponse.Data.Team, nil
}

/*
BuildTeamResponse ...
*/
func BuildTeamResponse(body io.ReadCloser) (*TeamResponse, error) {
	decoder := json.NewDecoder(body)
	result := &TeamResponse{}
	result.Data.Team = &Team{}

	if err := decoder.Decode(result); nil != err {
		return nil, err
	}

	return result, nil
}
