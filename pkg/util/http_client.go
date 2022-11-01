package util

import (
	"fmt"
	"io"
	"net/http"

	"dev.azure.com/banpudev/it.sd.dev/_git/go-azure-sdk.git/token"
)

type HttpClient struct {
	client *http.Client
	token  string
	header http.Header
}

func NewHttpClient() HttpClient {
	token, err := token.GetAccessToken()
	if err != nil {
		panic(err)
	}
	header := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{fmt.Sprintf("Bearer %v", token)},
	}
	return HttpClient{
		&http.Client{},
		token,
		header,
	}
}

func (h HttpClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = h.header
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
