package webclient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type fakeHTTPClient struct {
}

type fakeRequest map[string]*fakeResult

type fakeResult struct {
	Response *http.Response
	Error    error
}

func (fc *fakeHTTPClient) Get(api *url.URL) (*http.Response, error) {
	var request = fakeRequest{
		"https://fake.com/success": &fakeResult{
			&http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString(jsonString)),
			},
			nil,
		},
		"https://fake.com/error": &fakeResult{
			&http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString("")),
			},
			errors.New("something failed"),
		},
	}

	return request[api.String()].Response, request[api.String()].Error
}

func TestHttpClientGetSuccess(t *testing.T) {
	fakeClient := new(fakeHTTPClient)
	api, _ := url.Parse("https://fake.com/success")
	resp, err := fakeClient.Get(api)
	actualResponse, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	if nil != err {
		t.Errorf("\nGet returned error::\n\texpected nil\n\n\t actual %s", err.Error())
	}
	expectedResponse := []byte(jsonString)
	if !bytes.Equal(expectedResponse, actualResponse) {
		t.Errorf("\n Get returned unexpected response:\n\texpected: %q\n\n\tactual: %q", string(expectedResponse), actualResponse)
	}
}

func TestHttpClientGetError(t *testing.T) {
	fakeClient := new(fakeHTTPClient)
	api, _ := url.Parse("https://fake.com/error")
	resp, err := fakeClient.Get(api)
	actualResponse, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	expectedError := errors.New("something failed")

	if expectedError.Error() != err.Error() {
		t.Errorf("\nGet returned error::\n\texpected: %s\n\n\t actual: %s", expectedError.Error(), err.Error())
	}
	expectedResponse := []byte("")
	if !bytes.Equal(expectedResponse, actualResponse) {
		t.Errorf("\n Get returned unexpected response:\n\texpected: %q\n\n\tactual: %q", string(expectedResponse), actualResponse)
	}
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
