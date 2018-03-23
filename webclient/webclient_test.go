package webclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type httpClientMock struct {
	mock.Mock
}

func (hc *httpClientMock) Get(endpoint *url.URL) (*http.Response, error) {
	args := hc.Called(endpoint)

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestApiClientGetTeam(t *testing.T) {
	endpoint, _ := url.Parse("https://vintagemonster.onefootball.com")
	httpClientMock := new(httpClientMock)

	expectedEndpoint, _ := url.Parse("https://vintagemonster.onefootball.com/api/teams/en/10.json")

	httpClientMock.On("Get", expectedEndpoint).Return(
		&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(respString)),
		},
		nil,
	)

	client := apiClient{
		api:     endpoint,
		client:  httpClientMock,
		builder: BuildTeamResponse,
	}

	team, err := client.GetTeam(10)

	assert.Equal(t, &Team{
		ID:   10,
		Name: "CSKA Moscow",
		Players: []Player{
			{
				Country:   "Sweden",
				ID:        "124",
				FirstName: "Pontus",
				LastName:  "Wernbloom",
				Name:      "Pontus Wernbloom",
				Position:  "Midfielder",
				Number:    3,
				Age:       NewStringedInt(30),
			},
			{
				Country:   "Russian Federation",
				ID:        "282",
				FirstName: "Sergei",
				LastName:  "Ignashevich",
				Name:      "Sergei Ignashevich",
				Position:  "Defender",
				Number:    4,
				Age:       NewStringedInt(37),
			},
			{
				Country:   "Russian Federation",
				ID:        "284",
				FirstName: "Igor",
				LastName:  "Akinfeev",
				Name:      "Igor Akinfeev",
				Position:  "Goalkeeper",
				Number:    35,
				Age:       NewStringedInt(30),
			},
		},
	}, team)
	assert.Nil(t, err)

}
func TestApiClientGetTeamError(t *testing.T) {
	endpoint, _ := url.Parse("https://vintagemonster.onefootball.com")
	httpClientMock := new(httpClientMock)

	expectedEndpoint, _ := url.Parse("https://vintagemonster.onefootball.com/api/teams/en/10.json")

	httpClientMock.On("Get", expectedEndpoint).Return(&http.Response{}, errors.New("Server error"))

	client := apiClient{
		api:     endpoint,
		client:  httpClientMock,
		builder: BuildTeamResponse,
	}

	team, err := client.GetTeam(10)

	assert.Nil(t, team)
	assert.Error(t, err)
}

var respString = `{
  "status": "ok",
  "code": 0,
  "data": {
    "team": {
      "id": 10,
      "optaId": 1340,
      "name": "CSKA Moscow",
      "players": [
        {
          "country": "Sweden",
          "id": "124",
          "firstName": "Pontus",
          "lastName": "Wernbloom",
          "name": "Pontus Wernbloom",
          "position": "Midfielder",
          "number": 3,
          "birthDate": "1986-06-25",
          "age": "30",
          "height": 187,
          "weight": 85,
          "thumbnailSrc": "https:\/\/images.onefootball.com\/players\/124.jpg"
        },
        {
          "country": "Russian Federation",
          "id": "282",
          "firstName": "Sergei",
          "lastName": "Ignashevich",
          "name": "Sergei Ignashevich",
          "position": "Defender",
          "number": 4,
          "birthDate": "1979-07-14",
          "age": "37",
          "height": 187,
          "weight": 84,
          "thumbnailSrc": "https:\/\/images.onefootball.com\/PlayerPictures\/ru\/282.jpg"
        },
        {
          "country": "Russian Federation",
          "id": "284",
          "firstName": "Igor",
          "lastName": "Akinfeev",
          "name": "Igor Akinfeev",
          "position": "Goalkeeper",
          "number": 35,
          "birthDate": "1986-04-08",
          "age": "30",
          "height": 186,
          "weight": 82,
          "thumbnailSrc": "https:\/\/images.onefootball.com\/players\/284.jpg"
        }
      ]
    }
  },
  "message": "Team feed successfully generated. Api Version: 1"
}}`
