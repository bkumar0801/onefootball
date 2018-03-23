package webclient

import (
	"net/http"
	"net/url"
)

/*
HTTPClientInterface ...
*/
type HTTPClientInterface interface {
	Get(*url.URL) (*http.Response, error)
}

type httpClient struct {
}

func (hc *httpClient) Get(api *url.URL) (*http.Response, error) {
	request, err := http.NewRequest("GET", api.String(), nil)
	if nil != err {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response, err
}
