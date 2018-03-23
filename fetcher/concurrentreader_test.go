package fetcher

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/onefootball/webclient"
)

type fakeAPIClient struct {
	Error error
}

func (f fakeAPIClient) GetTeam(id int) (*webclient.Team, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(jsonString)),
	}
	teamResponse, err := webclient.BuildTeamResponse(response.Body)
	err = f.Error
	if nil != err {
		return nil, err
	}
	return teamResponse.Data.Team, nil
}

func TestShouldStartConcurrentReading(t *testing.T) {
	multiThreadAPIReader := NewMultiThreadAPIReader(fakeAPIClient{Error: nil}, 1, 1)

	multiThreadAPIReader.Start()
	team := multiThreadAPIReader.Read()

	expectedTeam := webclient.Team{
		ID:   10,
		Name: "CSKA Moscow",
	}

	if team.ID != expectedTeam.ID {
		t.Errorf("Unexpected ID received \n\t\t expected: %q \n\n\t\t actual: %q", team.ID, expectedTeam.ID)
	}
	if team.Name != expectedTeam.Name {
		t.Errorf("Unexpected Name received \n\t\t expected: %s \n\n\t\t actual: %s", team.Name, expectedTeam.Name)
	}
	multiThreadAPIReader.Stop()
}

var jsonString = `{
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
			  }
			]
		  }
		},
		"message": "Team feed successfully generated. Api Version: 1"
	  }}`